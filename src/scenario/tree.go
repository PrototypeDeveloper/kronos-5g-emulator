package scenario

import (
	"encoding/hex"
	"fmt"
	"net/netip"
	"reflect"
	"strconv"

	libnas "libnas"
	"libnas/nasType"
	libngap "libngap"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/common/plmn"
	"kronos/models"
)

// Tree renders an NGAP PDU as a tree for the web console: the
// IEs present with their values, the chosen alternative of each CHOICE, the
// NAS-PDUs and the PDU session Transfers decoded in place.
//
// plain is the plaintext of the NAS-PDUs in the order of NASPDUs (nil where
// the NAS-PDU is in clear). A ciphered NAS-PDU without its plaintext is shown
// as ciphered, unless asTemplate is set: templates are captured with NEA0,
// so their bodies are read as clear text.
func Tree(pdu []byte, plain [][]byte, asTemplate bool) (*models.TreeNode, error) {
	msg, err := libngap.Decode(pdu)
	if msg == nil {
		return nil, fmt.Errorf("cannot decode NGAP: %v", err)
	}
	t := &treeBuilder{plain: map[string][]byte{}, template: asTemplate}
	for i, b := range NASPDUs(msg) {
		if i < len(plain) && plain[i] != nil {
			t.plain[string(b)] = plain[i]
		}
	}
	root := &models.TreeNode{Name: "NGAP"}
	if info, ok := NGAPOf(msg); ok {
		root.Name = info.Name
	}
	if h, herr := ngapcodec.DecodeHeader(pdu); herr == nil {
		root.Value = fmt.Sprintf("%s (procedure code %d)", pduKinds[min(h.Kind, pduKindUnknown)], h.ProcedureCode)
	}
	root.Kids = t.fields(reflect.ValueOf(msg).Elem())
	if err != nil {
		root.Kids = append(root.Kids, &models.TreeNode{Name: "decoding notice", Value: err.Error()})
	}
	return root, nil
}

// pduKinds names the NGAP PDU kinds by header kind ("?" for any other).
var pduKinds = [...]string{"initiatingMessage", "successfulOutcome", "unsuccessfulOutcome", "?"}

var (
	bitStringType    = reflect.TypeOf(ngapType.BitString{})
	plmnIdentityType = reflect.TypeOf(ngapType.PLMNIdentity{})
	transportType    = reflect.TypeOf(ngapType.TransportLayerAddress{})
	nasContainerType = reflect.TypeOf(nasType.NASMessageContainer{})
)

// transferTypes are the OCTET STRING containers of NGAP that hold another
// APER value, by the field name that carries them.
var transferTypes = map[string]reflect.Type{}

// init fills transferTypes: each type name is also the field name that carries it.
func init() {
	for _, v := range []any{
		ngapType.PDUSessionResourceSetupRequestTransfer{}, ngapType.PDUSessionResourceSetupResponseTransfer{},
		ngapType.PDUSessionResourceSetupUnsuccessfulTransfer{}, ngapType.PDUSessionResourceReleaseCommandTransfer{},
		ngapType.PDUSessionResourceReleaseResponseTransfer{}, ngapType.PDUSessionResourceModifyRequestTransfer{},
		ngapType.PDUSessionResourceModifyResponseTransfer{}, ngapType.PDUSessionResourceModifyUnsuccessfulTransfer{},
		ngapType.PDUSessionResourceModifyIndicationTransfer{}, ngapType.PDUSessionResourceModifyConfirmTransfer{},
		ngapType.PDUSessionResourceModifyIndicationUnsuccessfulTransfer{}, ngapType.PDUSessionResourceNotifyTransfer{},
		ngapType.PDUSessionResourceNotifyReleasedTransfer{}, ngapType.PathSwitchRequestTransfer{},
		ngapType.PathSwitchRequestSetupFailedTransfer{}, ngapType.PathSwitchRequestAcknowledgeTransfer{},
		ngapType.PathSwitchRequestUnsuccessfulTransfer{}, ngapType.HandoverRequiredTransfer{},
		ngapType.HandoverCommandTransfer{}, ngapType.HandoverRequestAcknowledgeTransfer{},
		ngapType.HandoverPreparationUnsuccessfulTransfer{}, ngapType.HandoverResourceAllocationUnsuccessfulTransfer{},
		ngapType.SecondaryRATDataUsageReportTransfer{}, ngapType.UEContextResumeRequestTransfer{},
		ngapType.UEContextResumeResponseTransfer{}, ngapType.UEContextSuspendRequestTransfer{},
	} {
		t := reflect.TypeOf(v)
		transferTypes[t.Name()] = t
	}
}

// fields renders the fields of a struct: optional ones only when present
// (Has<Name>), and of a CHOICE or an IE entry only the chosen alternative.
func (t *treeBuilder) fields(v reflect.Value) []*models.TreeNode {
	typ := v.Type()
	if f, ok := typ.FieldByName("Present"); ok && len(f.Index) == 1 {
		return t.choice(v, f.Index[0])
	}
	var out []*models.TreeNode
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.IsExported() || (f.Type.Kind() == reflect.Bool && len(f.Name) > 3 && f.Name[:3] == "Has") {
			continue
		}
		if has := v.FieldByName("Has" + f.Name); has.IsValid() && has.Kind() == reflect.Bool && !has.Bool() {
			continue
		}
		fv := v.Field(i)
		switch f.Name {
		case "IEExtensions", "UnknownExtensions", "UnknownValue", "ChoiceExtensions", "AdditionalIEs", "UnknownIEs":
			if fv.IsZero() || (fv.Kind() == reflect.Slice && fv.Len() == 0) {
				continue
			}
		}
		if n := t.node(f.Name, fv); n != nil {
			out = append(out, n)
		}
	}
	return out
}

// choice renders the alternative of a CHOICE or an IE entry: libngap puts
// the alternatives right after Present, in the order of its values (1 is the first).
func (t *treeBuilder) choice(v reflect.Value, presentAt int) []*models.TreeNode {
	var k int
	switch p := v.Field(presentAt); p.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		k = int(p.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		k = int(p.Int())
	}
	if k <= 0 || presentAt+k >= v.NumField() {
		if u := v.FieldByName("UnknownValue"); u.IsValid() && u.Kind() == reflect.Slice && u.Len() > 0 {
			return []*models.TreeNode{{Name: "UnknownValue", Value: hexValue(u.Bytes())}}
		}
		return nil
	}
	f := v.Type().Field(presentAt + k)
	if n := t.node(f.Name, v.Field(presentAt+k)); n != nil {
		return []*models.TreeNode{n}
	}
	return nil
}

// isIEEntry reports a protocol IE entry (ID, Criticality, Present, values).
func isIEEntry(typ reflect.Type) bool {
	if typ.Kind() != reflect.Struct {
		return false
	}
	_, id := typ.FieldByName("ID")
	_, present := typ.FieldByName("Present")
	return id && present
}

// node renders one named value by its type; nil for a nil pointer.
func (t *treeBuilder) node(name string, v reflect.Value) *models.TreeNode {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return nil
		}
		return t.node(name, v.Elem())
	}
	// Types with a rendering of their own.
	typ := v.Type()
	switch typ {
	case bitStringType:
		return &models.TreeNode{Name: name, Value: bitStringText(v.Interface().(ngapType.BitString))}
	case nasPDUType:
		return t.nas(name, v.Field(0).Bytes())
	case plmnIdentityType:
		b := v.Field(0).Bytes()
		return &models.TreeNode{Name: name, Value: fmt.Sprintf("%x (%s)", b, plmn.String(b))}
	case transportType:
		bs := v.Field(0).Interface().(ngapType.BitString)
		return &models.TreeNode{Name: name, Value: transportText(bs)}
	case nasContainerType:
		return t.nested(name, v.Field(0).Bytes())
	}
	switch v.Kind() {
	case reflect.Struct:
		kids := t.fields(v)
		// A wrapper of one value (RANUENGAPID{Value}) shows the value itself.
		if typ.NumField() == 1 && len(kids) == 1 && kids[0].Name == typ.Field(0).Name && len(kids[0].Kids) == 0 {
			return &models.TreeNode{Name: name, Value: kids[0].Value}
		}
		return &models.TreeNode{Name: name, Kids: kids}
	case reflect.Slice, reflect.Array:
		// Octets: a Transfer or a NAS container is decoded in place, else hex.
		if typ.Elem().Kind() == reflect.Uint8 {
			b := byteSlice(v)
			if dt, ok := transferTypes[name]; ok && len(b) > 0 {
				return t.transfer(name, b, dt)
			}
			if name == "PayloadContainer" || name == "NASMessageContainer" {
				return t.nested(name, b)
			}
			return &models.TreeNode{Name: name, Value: hexValue(b)}
		}
		// A list: IE entries are flattened into their chosen values.
		count := fmt.Sprintf("%d items", v.Len())
		if v.Len() == 1 {
			count = "1 item"
		}
		n := &models.TreeNode{Name: name, Value: count}
		for i := 0; i < v.Len(); i++ {
			if isIEEntry(typ.Elem()) {
				n.Kids = append(n.Kids, t.fields(v.Index(i))...)
				continue
			}
			if c := t.node(fmt.Sprintf("[%d]", i), v.Index(i)); c != nil {
				n.Kids = append(n.Kids, c)
			}
		}
		return n
	case reflect.Bool:
		return &models.TreeNode{Name: name, Value: strconv.FormatBool(v.Bool())}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &models.TreeNode{Name: name, Value: strconv.FormatInt(v.Int(), 10)}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
		if u > decimalOnlyMax {
			return &models.TreeNode{Name: name, Value: fmt.Sprintf("%d (0x%x)", u, u)}
		}
		return &models.TreeNode{Name: name, Value: strconv.FormatUint(u, 10)}
	case reflect.String:
		return &models.TreeNode{Name: name, Value: strconv.Quote(v.String())}
	}
	return &models.TreeNode{Name: name, Value: fmt.Sprint(v.Interface())}
}

// byteSlice returns the octets of a byte slice or array.
func byteSlice(v reflect.Value) []byte {
	if v.Kind() == reflect.Slice {
		return v.Bytes()
	}
	b := make([]byte, v.Len())
	for i := range b {
		b[i] = byte(v.Index(i).Uint())
	}
	return b
}

// transfer decodes an OCTET STRING that holds an APER value.
func (t *treeBuilder) transfer(name string, b []byte, typ reflect.Type) *models.TreeNode {
	p := reflect.New(typ)
	err := ngapcodec.DecodeParameter(typ.Name(), b, p.Interface())
	if ngapcodec.IsFatal(err) {
		return &models.TreeNode{Name: name, Value: hexValue(b) + " (cannot decode: " + err.Error() + ")"}
	}
	return &models.TreeNode{Name: name, Value: fmt.Sprintf("%d octets", len(b)), Kids: t.fields(p.Elem())}
}

// nested shows a NAS message carried inside another (payload container,
// NAS message container) when it decodes, else its octets.
func (t *treeBuilder) nested(name string, b []byte) *models.TreeNode {
	if len(b) >= gmmHeaderLen && (b[0] == models.EPD5GMM || b[0] == models.EPD5GSM) {
		if m := t.nasMessage(b); m.Kids != nil {
			return &models.TreeNode{Name: name, Value: fmt.Sprintf("%d octets", len(b)), Kids: []*models.TreeNode{m}}
		}
	}
	return &models.TreeNode{Name: name, Value: hexValue(b)}
}

// securityHeaders names the 5GMM security header types by value.
var securityHeaders = [...]string{
	"plain", "integrity protected", "integrity protected and ciphered", "integrity protected (new context)", "integrity protected and ciphered (new context)",
}

// nas renders a NAS-PDU: the security header and the message, using the
// plaintext when the body is ciphered.
func (t *treeBuilder) nas(name string, b []byte) *models.TreeNode {
	n := &models.TreeNode{Name: name, Value: fmt.Sprintf("%d octets", len(b))}
	if len(b) < gmmHeaderLen {
		n.Value = hexValue(b)
		return n
	}
	if b[0] != models.EPD5GMM || b[1]&secHeaderTypeMask == 0 || len(b) < secHeaderLen+1 {
		n.Kids = []*models.TreeNode{t.nasMessage(b)}
		return n
	}
	sht := b[1] & secHeaderTypeMask
	label := "?"
	if int(sht) < len(securityHeaders) {
		label = securityHeaders[sht]
	}
	n.Kids = []*models.TreeNode{
		{Name: "SecurityHeaderType", Value: fmt.Sprintf("%d (%s)", sht, label)},
		{Name: "MAC", Value: hex.EncodeToString(b[2:6])},
		{Name: "SequenceNumber", Value: strconv.Itoa(int(b[6]))},
	}
	body := b[secHeaderLen:]
	switch p := t.plain[string(b)]; {
	case p != nil:
		n.Kids = append(n.Kids, &models.TreeNode{Name: "plaintext (deciphered body)", Kids: []*models.TreeNode{t.nasMessage(p)}})
	case sht == shtIntegrity || sht == shtIntegrityNew:
		n.Kids = append(n.Kids, t.nasMessage(body)) // integrity protected only: the body is clear
	case t.template:
		n.Kids = append(n.Kids, &models.TreeNode{Name: "body (a template is read as NEA0)", Kids: []*models.TreeNode{t.nasMessage(body)}})
	default:
		n.Kids = append(n.Kids, &models.TreeNode{Name: "ciphered body", Value: hexValue(body)})
	}
	return n
}

// nasMessage decodes a plain 5GMM or 5GSM message.
func (t *treeBuilder) nasMessage(b []byte) *models.TreeNode {
	name := "NAS"
	switch {
	case len(b) >= gmmHeaderLen && b[0] == models.EPD5GMM && b[1] == 0:
		if info, ok := NASByType(models.EPD5GMM, b[2]); ok {
			name = info.Name
		}
	case len(b) >= gsmHeaderLen && b[0] == models.EPD5GSM:
		if info, ok := NASByType(models.EPD5GSM, b[3]); ok {
			name = info.Name
		}
	}
	m, err := libnas.Decode(b, nil)
	if err != nil {
		return &models.TreeNode{Name: name, Value: hexValue(b) + " (cannot decode: " + err.Error() + ")"}
	}
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	kids := t.fields(v)
	if kids == nil {
		kids = []*models.TreeNode{} // decoded, with no IEs
	}
	return &models.TreeNode{Name: name, Kids: kids}
}

// bitStringText shows a short BIT STRING as its value (APER puts the bits
// first in the octets: AMF Set ID 1 is 0040), a long one as octets.
func bitStringText(bs ngapType.BitString) string {
	if bs.BitLength == 0 || bs.BitLength > maxBitsAsNumber || uint64(len(bs.Bytes)*8) < bs.BitLength {
		return fmt.Sprintf("%x (%d bits)", bs.Bytes, bs.BitLength)
	}
	var v uint64
	for i := uint64(0); i < bs.BitLength; i++ {
		v = v<<1 | uint64(bs.Bytes[i/8]>>(7-i%8)&1)
	}
	return fmt.Sprintf("0x%x (%d bits)", v, bs.BitLength)
}

// transportText shows a TransportLayerAddress as IPv4, IPv6 or both.
func transportText(bs ngapType.BitString) string {
	switch {
	case bs.BitLength == ipv4AddrLen*8 && len(bs.Bytes) == ipv4AddrLen:
		return netip.AddrFrom4([4]byte(bs.Bytes)).String()
	case bs.BitLength == ipv6AddrLen*8 && len(bs.Bytes) == ipv6AddrLen:
		return netip.AddrFrom16([16]byte(bs.Bytes)).String()
	case bs.BitLength == dualAddrLen*8 && len(bs.Bytes) == dualAddrLen:
		return netip.AddrFrom4([4]byte(bs.Bytes[:ipv4AddrLen])).String() + " / " + netip.AddrFrom16([16]byte(bs.Bytes[ipv4AddrLen:])).String()
	}
	return bitStringText(bs)
}

// hexValue shows octets in hex, or "(empty)" when there are none.
func hexValue(b []byte) string {
	if len(b) == 0 {
		return "(empty)"
	}
	return hex.EncodeToString(b)
}
