package scenario

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/netip"
	"reflect"
	"strconv"
	"strings"

	libnas "libnas"
	"libnas/nasType"
	libngap "libngap"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/common/plmn"
	"kronos/models"
)

// The signal editor: a signal's octets are decoded into the libngap and
// libnas values, shown as a tree of fields, changed one operation at a
// time through reflection and encoded again. Octets that hold another
// message (a NAS-PDU, a NAS container, a PDU session Transfer) are decoded
// on the way down and encoded back on the way up, so lengths follow.
//
// A path names a field: Go field names, list indexes and ">" (into the
// message held by octets), joined by "/". NAS security header fields are
// "@SecurityHeaderType", "@MAC" and "@SequenceNumber" under a NAS-PDU.

// AutoRule reports whether a role replaces a field when the signal is
// sent (the editor marks it; it can still be edited, for a send without
// the mapping), with a note for a field it changes only in part or only
// sometimes.
type AutoRule func(FieldRef) (auto bool, note string)

// EditSignal applies op (when not nil) to the octets of a signal and
// returns the new octets with their tree. An op the message cannot take, or
// one that does not encode, is an error and changes nothing.
func EditSignal(b []byte, mode Mode, op *models.EditOp, auto AutoRule) ([]byte, *models.EditNode, error) {
	if op != nil {
		nb, err := applyEdit(b, *op)
		if err != nil {
			return nil, nil, err
		}
		b = nb
	}
	return b, EditTree(b, mode, auto), nil
}

// applyEdit decodes the signal, applies op at its path and encodes it again.
func applyEdit(b []byte, op models.EditOp) ([]byte, error) {
	if op.Path == "" && op.Kind == "hex" { // the whole signal
		return parseHex(op.Value)
	}
	msg, err := libngap.Decode(b)
	if msg == nil {
		return nil, fmt.Errorf("the signal does not decode as NGAP (%v): replace its octets as hex", err)
	}
	root := loc{v: reflect.ValueOf(msg).Elem()}
	if err := walkEdit(root, splitPath(op.Path), func(l loc) error { return applyAt(l, op) }); err != nil {
		return nil, err
	}
	out, err := libngap.Encode(msg)
	if err != nil {
		return nil, fmt.Errorf("the change does not encode: %v", err)
	}
	return out, nil
}

// splitPath splits an editor path into its segments.
func splitPath(p string) []string {
	if p == "" {
		return nil
	}
	return strings.Split(p, "/")
}

// joinPath appends a segment to an editor path.
func joinPath(base, seg string) string {
	if base == "" {
		return seg
	}
	return base + "/" + seg
}

// The fields that carry extensions and unknown values: shown, never edited.
var extensionFields = map[string]bool{
	"IEExtensions": true, "UnknownExtensions": true, "AdditionalIEs": true, "UnknownValue": true, "ChoiceExtensions": true,
}

var payloadContainerType = reflect.TypeOf(nasType.PayloadContainer{})

// walkEdit goes down segs from l and runs fn there. Octets that hold a message
// (">") are decoded on the way and written back after fn.
func walkEdit(l loc, segs []string, fn func(loc) error) error {
	if len(segs) == 0 {
		return fn(l)
	}
	seg, rest := segs[0], segs[1:]
	v := l.v
	// ">": open the message held by these octets, go on inside it and write it back.
	if seg == ">" {
		c, err := openContainer(l)
		if err != nil {
			return err
		}
		next := c.root
		// "@X" is a NAS security header field, kept outside the decoded message.
		if len(rest) > 0 && strings.HasPrefix(rest[0], "@") {
			if c.header == nil {
				return fmt.Errorf("%s: this NAS message has no security header", rest[0])
			}
			hv := reflect.ValueOf(c.header).Elem()
			f, ok := hv.Type().FieldByName(rest[0][1:])
			if !ok {
				return fmt.Errorf("no NAS security header field %s", rest[0])
			}
			if c.header.SecurityHeaderType == 0 && f.Name != "SecurityHeaderType" {
				return fmt.Errorf("the NAS message is plain: set @SecurityHeaderType first")
			}
			next, rest = loc{v: hv.FieldByIndex(f.Index), parent: hv, field: f}, rest[1:]
		}
		if err := walkEdit(next, rest, fn); err != nil {
			return err
		}
		b, err := c.close()
		if err != nil {
			return err
		}
		v.SetBytes(b)
		return nil
	}
	// Otherwise step into a struct field or a list item.
	switch v.Kind() {
	case reflect.Struct:
		// Only the chosen alternative of a CHOICE can be entered.
		if alts, at := choiceOf(v.Type()); alts != nil {
			if cur := chosen(v, at, alts); seg != cur {
				return fmt.Errorf("%s is not the alternative chosen (%s)", seg, cur)
			}
		}
		f, ok := v.Type().FieldByName(seg)
		if !ok || !f.IsExported() {
			return fmt.Errorf("%s has no field %s", v.Type().Name(), seg)
		}
		if has := v.FieldByName("Has" + seg); has.IsValid() && has.Kind() == reflect.Bool && !has.Bool() {
			return fmt.Errorf("%s is not present (add it first)", seg)
		}
		return walkEdit(loc{v: v.FieldByIndex(f.Index), parent: v, field: f}, rest, fn)
	case reflect.Slice:
		i, err := strconv.Atoi(seg)
		if err != nil || i < 0 || i >= v.Len() {
			return fmt.Errorf("the list has no item %s", seg)
		}
		return walkEdit(loc{v: v.Index(i), parent: v, field: l.field, item: true, index: i}, rest, fn)
	}
	return fmt.Errorf("cannot go into %s at %s", v.Type(), seg)
}

// choiceOf returns the alternatives of a CHOICE (the fields after Present)
// and the index of Present; nil if typ is not a CHOICE.
func choiceOf(typ reflect.Type) ([]string, int) {
	if typ.Kind() != reflect.Struct {
		return nil, 0
	}
	f, ok := typ.FieldByName("Present")
	if !ok || len(f.Index) != 1 {
		return nil, 0
	}
	var alts []string
	for i := f.Index[0] + 1; i < typ.NumField(); i++ {
		if n := typ.Field(i).Name; !extensionFields[n] {
			alts = append(alts, n)
		}
	}
	return alts, f.Index[0]
}

// chosen returns the name of the alternative chosen ("" if none).
func chosen(v reflect.Value, at int, alts []string) string {
	k := intOf(v.Field(at))
	if k >= 1 && int(k) <= len(alts) {
		return alts[k-1]
	}
	return ""
}

// intOf returns an integer of any kind as int64 (0 for other kinds).
func intOf(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint())
	}
	return 0
}

// setInt stores n into an integer of any kind.
func setInt(v reflect.Value, n int64) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(n))
	}
}

// containerKind tells what message octets hold, from their field: "" when
// they are plain octets.
func containerKind(l loc) string {
	if l.v.Kind() != reflect.Slice || l.v.Type().Elem().Kind() != reflect.Uint8 || l.item {
		return ""
	}
	name := l.field.Name
	var pt reflect.Type
	if l.parent.IsValid() {
		pt = l.parent.Type()
	}
	switch {
	case pt == nasPDUType && name == "Value":
		return "nas-pdu"
	case name == "PayloadContainer" || name == "NASMessageContainer",
		(pt == payloadContainerType || pt == nasContainerType) && name == "Value":
		return "nas"
	}
	if _, ok := transferTypes[name]; ok {
		return "transfer"
	}
	return ""
}

// openContainer decodes the message held by the octets at l; its close encodes it back.
func openContainer(l loc) (*container, error) {
	b := l.v.Bytes()
	switch kind := containerKind(l); kind {
	case "nas-pdu", "nas":
		c := &container{kind: kind}
		body, first := b, byte(0)
		// Split off the security header of a protected NAS-PDU.
		if kind == "nas-pdu" && len(b) >= secHeaderLen && b[0] == models.EPD5GMM && b[1]&secHeaderTypeMask != 0 {
			c.header = &nasHeader{SecurityHeaderType: b[1] & secHeaderTypeMask, MAC: append([]byte(nil), b[2:6]...), SequenceNumber: b[6]}
			first, body = b[1]&secHeaderSpareMask, b[secHeaderLen:]
		} else if kind == "nas-pdu" && len(b) >= gmmHeaderLen && b[0] == models.EPD5GMM {
			c.header = &nasHeader{MAC: make([]byte, macLen)} // plain: the header type can be set
		}
		m, err := libnas.Decode(body, nil)
		if err != nil {
			return nil, fmt.Errorf("the NAS message does not decode (%v): replace its octets as hex", err)
		}
		rv := reflect.ValueOf(m)
		if rv.Kind() != reflect.Pointer {
			return nil, fmt.Errorf("the NAS message %T cannot be edited", m)
		}
		c.root, c.name = loc{v: rv.Elem()}, rv.Elem().Type().Name()
		// Encode the NAS message and put the security header back in front.
		c.close = func() ([]byte, error) {
			out, err := libnas.Encode(m, nil)
			if err != nil || c.header == nil || c.header.SecurityHeaderType == 0 {
				return out, err
			}
			h := c.header
			mac := append(append([]byte(nil), h.MAC...), 0, 0, 0, 0)[:macLen]
			return append(append([]byte{models.EPD5GMM, first | h.SecurityHeaderType}, mac...), append([]byte{h.SequenceNumber}, out...)...), nil
		}
		return c, nil
	case "transfer":
		typ := transferTypes[l.field.Name]
		p := reflect.New(typ)
		if err := ngapcodec.DecodeParameter(typ.Name(), b, p.Interface()); ngapcodec.IsFatal(err) {
			return nil, fmt.Errorf("the %s does not decode (%v): replace its octets as hex", typ.Name(), err)
		}
		return &container{kind: kind, root: loc{v: p.Elem()}, name: typ.Name(), close: func() ([]byte, error) {
			return ngapcodec.EncodeParameter(typ.Name(), p.Elem().Interface())
		}}, nil
	}
	return nil, fmt.Errorf("%s holds no message", l.field.Name)
}

// applyAt makes the change at its place.
func applyAt(l loc, op models.EditOp) error {
	v := l.v
	switch op.Kind {
	case "set":
		return setLeaf(l, op)
	case "hex":
		if v.Kind() != reflect.Slice || v.Type().Elem().Kind() != reflect.Uint8 {
			return errors.New("only octets are replaced as hex")
		}
		b, err := parseHex(op.Value)
		if err != nil {
			return err
		}
		// Octets that hold a message are not size-checked.
		if err := checkSize(parseAper(l.field.Tag), int64(len(b)), "octets"); err != nil && containerKind(l) == "" {
			return err
		}
		v.SetBytes(b)
		return nil
	case "add":
		// On a list: append a minimal item.
		if v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 {
			tag := parseAper(l.field.Tag)
			if tag.sizeUB != nil && !tag.sizeExt && int64(v.Len()) >= *tag.sizeUB {
				return fmt.Errorf("the list holds at most %d items", *tag.sizeUB)
			}
			item := reflect.New(v.Type().Elem()).Elem()
			if err := fill(item, reflect.StructField{}); err != nil {
				return err
			}
			v.Set(reflect.Append(v, item))
			return nil
		}
		// On a struct: make an optional field present with minimal content.
		if v.Kind() != reflect.Struct {
			return errors.New("only a struct or a list takes an added field")
		}
		has := v.FieldByName("Has" + op.Value)
		f, ok := v.Type().FieldByName(op.Value)
		if !ok || !has.IsValid() || has.Kind() != reflect.Bool || extensionFields[op.Value] {
			return fmt.Errorf("%s is not an optional field here", op.Value)
		}
		if has.Bool() {
			return fmt.Errorf("%s is present already", op.Value)
		}
		fv := v.FieldByIndex(f.Index)
		fv.Set(reflect.Zero(fv.Type()))
		if err := fill(fv, f); err != nil {
			return err
		}
		has.SetBool(true)
		return nil
	case "remove":
		// A list item: drop it while the list keeps its minimum size.
		if l.item {
			tag := parseAper(l.field.Tag)
			if tag.sizeLB != nil && !tag.sizeExt && int64(l.parent.Len()) <= *tag.sizeLB {
				return fmt.Errorf("the list holds at least %d items", *tag.sizeLB)
			}
			p := l.parent
			out := reflect.MakeSlice(p.Type(), 0, p.Len()-1)
			out = reflect.AppendSlice(out, p.Slice(0, l.index))
			out = reflect.AppendSlice(out, p.Slice(l.index+1, p.Len()))
			p.Set(out)
			return nil
		}
		// An optional field: clear its presence flag and its value.
		if !l.parent.IsValid() || l.parent.Kind() != reflect.Struct {
			return errors.New("this cannot be removed")
		}
		has := l.parent.FieldByName("Has" + l.field.Name)
		if !has.IsValid() || has.Kind() != reflect.Bool {
			return fmt.Errorf("%s is mandatory", l.field.Name)
		}
		has.SetBool(false)
		v.Set(reflect.Zero(v.Type()))
		return nil
	case "choose":
		alts, at := choiceOf(v.Type())
		if alts == nil {
			return errors.New("this is not a CHOICE")
		}
		// Switch to the named alternative, filled with minimal content.
		for i, a := range alts {
			if a != op.Value {
				continue
			}
			if chosen(v, at, alts) == a {
				return nil
			}
			f, _ := v.Type().FieldByName(a)
			fv := v.FieldByIndex(f.Index)
			fv.Set(reflect.Zero(fv.Type()))
			if err := fill(fv, f); err != nil {
				return err
			}
			setInt(v.Field(at), int64(i+1))
			return nil
		}
		return fmt.Errorf("%s is not an alternative (%s)", op.Value, strings.Join(alts, ", "))
	}
	return fmt.Errorf("unknown edit %q", op.Kind)
}

// fill gives a new value the smallest content its constraints allow:
// lower bounds, the first alternative, the mandatory fields, an encoded
// minimal Transfer.
func fill(v reflect.Value, f reflect.StructField) error {
	tag := parseAper(f.Tag)
	lb := func(p *int64) int64 {
		if p != nil && *p > 0 {
			return *p
		}
		return 0
	}
	// A BIT STRING gets zero bits of its minimum length.
	if v.Type() == bitStringType {
		n := lb(tag.sizeLB)
		v.Set(reflect.ValueOf(ngapType.BitString{Bytes: make([]byte, (n+7)/8), BitLength: uint64(n)}))
		return nil
	}
	switch v.Kind() {
	case reflect.Struct:
		// A CHOICE takes its first alternative.
		if alts, at := choiceOf(v.Type()); alts != nil {
			setInt(v.Field(at), 1)
			af, _ := v.Type().FieldByName(alts[0])
			return fill(v.FieldByIndex(af.Index), af)
		}
		// Otherwise fill the mandatory fields; optional ones stay absent.
		typ := v.Type()
		for i := 0; i < typ.NumField(); i++ {
			sf := typ.Field(i)
			if !sf.IsExported() || extensionFields[sf.Name] || isHasFlag(typ, sf) {
				continue
			}
			if _, optional := typ.FieldByName("Has" + sf.Name); optional {
				continue
			}
			if err := fill(v.Field(i), sf); err != nil {
				return err
			}
		}
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			// Octets that hold a Transfer get an encoded minimal Transfer.
			if typ, ok := transferTypes[f.Name]; ok {
				p := reflect.New(typ)
				if err := fill(p.Elem(), reflect.StructField{}); err != nil {
					return err
				}
				b, err := ngapcodec.EncodeParameter(typ.Name(), p.Elem().Interface())
				if err != nil {
					return fmt.Errorf("cannot make an empty %s: %v", typ.Name(), err)
				}
				v.SetBytes(b)
				return nil
			}
			v.SetBytes(make([]byte, lb(tag.sizeLB)))
			return nil
		}
		// A list gets its minimum number of filled items.
		n := lb(tag.sizeLB)
		s := reflect.MakeSlice(v.Type(), int(n), int(n))
		for i := 0; i < int(n); i++ {
			if err := fill(s.Index(i), reflect.StructField{}); err != nil {
				return err
			}
		}
		v.Set(s)
	case reflect.String:
		v.SetString(strings.Repeat("a", int(lb(tag.sizeLB))))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if tag.lb != nil {
			setInt(v, *tag.lb)
		}
	}
	return nil
}

// isHasFlag reports a presence flag (HasX for a field X).
func isHasFlag(typ reflect.Type, f reflect.StructField) bool {
	if f.Type.Kind() != reflect.Bool || !strings.HasPrefix(f.Name, "Has") || len(f.Name) == 3 {
		return false
	}
	_, ok := typ.FieldByName(f.Name[3:])
	return ok
}

// setLeaf parses a value into a leaf.
func setLeaf(l loc, op models.EditOp) error {
	v, tag, s := l.v, parseAper(l.field.Tag), strings.TrimSpace(op.Value)
	pt := reflect.Type(nil)
	if l.parent.IsValid() {
		pt = l.parent.Type()
	}
	// A BIT STRING: IP addresses for a TransportLayerAddress, else a number or hex.
	if v.Type() == bitStringType {
		var bs ngapType.BitString
		var err error
		if pt == transportType {
			bs, err = parseAddresses(s)
		} else {
			bits := op.Bits
			if bits == 0 {
				bits = int(v.Interface().(ngapType.BitString).BitLength)
			}
			if err := checkSize(tag, int64(bits), "bits"); err != nil {
				return err
			}
			bs, err = parseBits(s, bits)
		}
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(bs))
		return nil
	}
	switch v.Kind() {
	case reflect.Slice:
		if v.Type().Elem().Kind() != reflect.Uint8 {
			break
		}
		// Octets: MCC/MNC for a PLMN identity, else hex.
		var b []byte
		var err error
		if pt == plmnIdentityType {
			b, err = parsePLMN(s)
		} else {
			b, err = parseHex(s)
		}
		if err != nil {
			return err
		}
		if err := checkSize(tag, int64(len(b)), "octets"); err != nil {
			return err
		}
		v.SetBytes(b)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return fmt.Errorf("%q is not an integer", s)
		}
		if err := checkValue(tag, n); err != nil {
			return err
		}
		if v.OverflowInt(n) {
			return fmt.Errorf("%d does not fit in %s", n, v.Type())
		}
		v.SetInt(n)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return fmt.Errorf("%q is not an integer of 0 or more", s)
		}
		// The bounds are int64: values too large for them skip the check.
		if n <= maxBoundCheckedUint {
			if err := checkValue(tag, int64(n)); err != nil {
				return err
			}
		}
		if v.OverflowUint(n) {
			return fmt.Errorf("%d does not fit in %s", n, v.Type())
		}
		v.SetUint(n)
		return nil
	case reflect.String:
		if err := checkSize(tag, int64(len(op.Value)), "characters"); err != nil {
			return err
		}
		v.SetString(op.Value)
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return fmt.Errorf("%q is not true or false", s)
		}
		v.SetBool(b)
		return nil
	}
	return fmt.Errorf("%s is not a value (add, remove or choose its fields)", v.Type())
}

// parseAper reads the value and size constraints of an aper struct tag.
func parseAper(tag reflect.StructTag) aperTag {
	var t aperTag
	for _, part := range strings.Split(tag.Get("aper"), ",") {
		k, val, _ := strings.Cut(part, ":")
		n, err := strconv.ParseInt(val, 10, 64)
		num := func() *int64 {
			if err != nil {
				return nil
			}
			return &n
		}
		switch k {
		case "valueExt":
			t.ext = true
		case "sizeExt":
			t.sizeExt = true
		case "valueLB":
			t.lb = num()
		case "valueUB":
			t.ub = num()
		case "sizeLB":
			t.sizeLB = num()
		case "sizeUB":
			t.sizeUB = num()
		}
	}
	return t
}

// checkValue checks n against the value bounds (no upper bound when extensible).
func checkValue(t aperTag, n int64) error {
	if t.lb != nil && n < *t.lb {
		return fmt.Errorf("%d is below %d", n, *t.lb)
	}
	if t.ub != nil && n > *t.ub && !t.ext {
		return fmt.Errorf("%d is above %d", n, *t.ub)
	}
	return nil
}

// checkSize checks a length of n units against the size bounds, unless extensible.
func checkSize(t aperTag, n int64, unit string) error {
	if t.sizeExt {
		return nil
	}
	if t.sizeLB != nil && t.sizeUB != nil && *t.sizeLB == *t.sizeUB && n != *t.sizeLB {
		return fmt.Errorf("it must have %d %s (%d given)", *t.sizeLB, unit, n)
	}
	if t.sizeLB != nil && n < *t.sizeLB {
		return fmt.Errorf("it must have at least %d %s (%d given)", *t.sizeLB, unit, n)
	}
	if t.sizeUB != nil && n > *t.sizeUB {
		return fmt.Errorf("it must have at most %d %s (%d given)", *t.sizeUB, unit, n)
	}
	return nil
}

// parseHex reads octets written in hex (spaces, ':' and a 0x prefix allowed).
func parseHex(s string) ([]byte, error) {
	s = strings.TrimPrefix(strings.TrimPrefix(strings.TrimSpace(s), "0x"), "0X")
	s = strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "", ":", "").Replace(s)
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("%q is not hex octets", s)
	}
	return b, nil
}

// parseBits reads a BIT STRING of n bits: a number when it fits in 64 bits
// (its bits go first in the octets, as APER puts them), else hex octets.
func parseBits(s string, n int) (ngapType.BitString, error) {
	size := (n + 7) / 8
	if n > maxBitsAsNumber {
		b, err := parseHex(s)
		if err != nil {
			return ngapType.BitString{}, err
		}
		if len(b) != size {
			return ngapType.BitString{}, fmt.Errorf("%d bits take %d octets (%d given)", n, size, len(b))
		}
		return ngapType.BitString{Bytes: b, BitLength: uint64(n)}, nil
	}
	val, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return ngapType.BitString{}, fmt.Errorf("%q is not a number (0x… for hex)", s)
	}
	if n < maxBitsAsNumber && val>>uint(n) != 0 {
		return ngapType.BitString{}, fmt.Errorf("%s does not fit in %d bits", s, n)
	}
	b := make([]byte, size)
	for i := 0; i < n; i++ {
		if val>>uint(n-1-i)&1 != 0 {
			b[i/8] |= 0x80 >> uint(i%8)
		}
	}
	return ngapType.BitString{Bytes: b, BitLength: uint64(n)}, nil
}

// parsePLMN reads "208/93", "208-93" or the 3 octets in hex.
func parsePLMN(s string) ([]byte, error) {
	if mcc, mnc, ok := strings.Cut(strings.ReplaceAll(s, "-", "/"), "/"); ok {
		b, ok := plmn.Encode(strings.TrimSpace(mcc), strings.TrimSpace(mnc))
		if !ok {
			return nil, fmt.Errorf("%q is not an MCC/MNC (3 and 2 or 3 digits)", s)
		}
		return b, nil
	}
	return parseHex(s)
}

// parseAddresses reads a TransportLayerAddress: an IPv4 or IPv6 address, or
// both as "v4,v6".
func parseAddresses(s string) (ngapType.BitString, error) {
	var b []byte
	for _, part := range strings.FieldsFunc(s, func(r rune) bool { return r == ',' || r == ' ' || r == '/' }) {
		a, err := netip.ParseAddr(part)
		if err != nil {
			return ngapType.BitString{}, fmt.Errorf("%q is not an IP address", part)
		}
		b = append(b, a.Unmap().AsSlice()...)
	}
	switch len(b) {
	case ipv4AddrLen, ipv6AddrLen, dualAddrLen:
		return ngapType.BitString{Bytes: b, BitLength: uint64(len(b) * 8)}, nil
	}
	return ngapType.BitString{}, fmt.Errorf("%q: write an IPv4 address, an IPv6 address or both (v4,v6)", s)
}

// EditTree renders the signal for the editor.
func EditTree(b []byte, mode Mode, auto AutoRule) *models.EditNode {
	msg, err := libngap.Decode(b)
	if msg == nil {
		return &models.EditNode{Name: "octets", Kind: "octets", Value: hex.EncodeToString(b),
			Note: fmt.Sprintf("does not decode as NGAP (%v)", err)}
	}
	name := "NGAP"
	if info, ok := NGAPOf(msg); ok {
		name = info.Name
	}
	vb := &viewer{mode: mode, auto: auto, ngap: name}
	root := vb.node(name, "", loc{v: reflect.ValueOf(msg).Elem()}, ctx{}, false)
	if h, herr := ngapcodec.DecodeHeader(b); herr == nil {
		root.Hint = fmt.Sprintf("%s (procedure code %d)", pduKinds[min(h.Kind, pduKindUnknown)], h.ProcedureCode)
	}
	if err != nil {
		root.Note = "decoded with a notice: " + err.Error()
	}
	return root
}

// rel returns path relative to the innermost NAS message.
func (c ctx) rel(path string) string {
	if c.base == "" {
		return path
	}
	return strings.TrimPrefix(strings.TrimPrefix(path, c.base), "/")
}

// node renders the value at l as an editor node, recursing into its fields and items.
func (vb *viewer) node(name, path string, l loc, c ctx, auto bool) *models.EditNode {
	n := &models.EditNode{Name: name, Path: path, Auto: auto}
	// Under an auto-fixed field everything is auto-fixed; otherwise ask the rule.
	if !auto && vb.auto != nil {
		n.Auto, n.Note = vb.auto(FieldRef{Mode: vb.mode, NGAP: vb.ngap, NAS: c.nas, Outer: c.outer, Path: c.rel(path)})
	}
	v, tag := l.v, parseAper(l.field.Tag)
	var pt reflect.Type
	if l.parent.IsValid() {
		pt = l.parent.Type()
	}
	bounds := func(lo, hi *int64, ext bool) {
		n.Min, n.Max, n.Ext = lo, hi, ext
	}
	// A BIT STRING: IP addresses for a TransportLayerAddress, else a number or hex.
	if v.Type() == bitStringType {
		bs := v.Interface().(ngapType.BitString)
		if pt == transportType {
			n.Kind, n.Value = "ip", transportText(bs)
			return n
		}
		n.Kind, n.Bits = "bits", int(bs.BitLength)
		bounds(tag.sizeLB, tag.sizeUB, tag.sizeExt)
		if bs.BitLength <= maxBitsAsNumber && uint64(len(bs.Bytes)*8) >= bs.BitLength {
			var x uint64
			for i := uint64(0); i < bs.BitLength; i++ {
				x = x<<1 | uint64(bs.Bytes[i/8]>>(7-i%8)&1)
			}
			n.Value = fmt.Sprintf("0x%x", x)
		} else {
			n.Value = hex.EncodeToString(bs.Bytes)
		}
		return n
	}
	switch v.Kind() {
	case reflect.Struct:
		typ := v.Type()
		if typ.NumField() == 1 && typ.Field(0).IsExported() { // a wrapper: show its value
			f := typ.Field(0)
			inner := vb.node(name, joinPath(path, f.Name), loc{v: v.Field(0), parent: v, field: f}, c, n.Auto)
			if inner.Note == "" {
				inner.Note = n.Note
			}
			return inner
		}
		if alts, at := choiceOf(typ); alts != nil {
			n.Kind, n.Alts, n.Alt = "choice", alts, chosen(v, at, alts)
			if n.Alt != "" {
				f, _ := typ.FieldByName(n.Alt)
				n.Kids = []*models.EditNode{vb.node(n.Alt, joinPath(path, n.Alt), loc{v: v.FieldByIndex(f.Index), parent: v, field: f}, c, n.Auto)}
			}
			return n
		}
		// A struct: present fields become kids, absent optional ones can be added.
		n.Kind = "struct"
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			if !f.IsExported() || isHasFlag(typ, f) {
				continue
			}
			fv := v.Field(i)
			has := v.FieldByName("Has" + f.Name)
			optional := has.IsValid() && has.Kind() == reflect.Bool
			// Extensions are shown read-only, and only when present.
			if extensionFields[f.Name] {
				if (optional && has.Bool()) || (!optional && !fv.IsZero() && !(fv.Kind() == reflect.Slice && fv.Len() == 0)) {
					n.Kids = append(n.Kids, &models.EditNode{Name: f.Name, Path: joinPath(path, f.Name), Kind: "ext",
						Value: extensionText(fv), Note: "kept as it is (not edited)"})
				}
				continue
			}
			if optional && !has.Bool() {
				n.Add = append(n.Add, f.Name)
				continue
			}
			kid := vb.node(f.Name, joinPath(path, f.Name), loc{v: fv, parent: v, field: f}, c, n.Auto)
			if optional {
				kid.Del = joinPath(path, f.Name)
			}
			n.Kids = append(n.Kids, kid)
		}
		return n
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return vb.octets(n, path, l, c, tag, pt)
		}
		// A list: one kid per item; IE container entries are read-only.
		n.Kind = "list"
		bounds(tag.sizeLB, tag.sizeUB, tag.sizeExt)
		n.Hint = fmt.Sprintf("%d items", v.Len())
		if v.Len() == 1 {
			n.Hint = "1 item"
		}
		readOnly := isIEEntry(v.Type().Elem())
		for i := 0; i < v.Len(); i++ {
			ip := joinPath(path, strconv.Itoa(i))
			if readOnly {
				n.Kids = append(n.Kids, &models.EditNode{Name: fmt.Sprintf("[%d]", i), Path: ip, Kind: "ext", Value: extensionText(v.Index(i)), Note: "kept as it is (not edited)"})
				continue
			}
			kid := vb.node(fmt.Sprintf("[%d]", i), ip, loc{v: v.Index(i), parent: v, field: l.field, item: true, index: i}, c, n.Auto)
			kid.Del = ip
			n.Kids = append(n.Kids, kid)
		}
		return n
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n.Kind, n.Value = "int", strconv.FormatInt(v.Int(), 10)
		bounds(tag.lb, tag.ub, tag.ext)
		// Without an aper tag (NAS), use the range of the Go type.
		if tag.lb == nil && tag.ub == nil {
			lo, hi := intRange(v.Type())
			bounds(&lo, &hi, false)
		}
		return n
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
		n.Kind, n.Value = "int", strconv.FormatUint(u, 10)
		if u > decimalOnlyMax {
			n.Hint = fmt.Sprintf("0x%x", u)
		}
		bounds(tag.lb, tag.ub, tag.ext)
		if tag.lb == nil && tag.ub == nil {
			lo, hi := intRange(v.Type())
			bounds(&lo, &hi, false)
		}
		return n
	case reflect.String:
		n.Kind, n.Value = "string", v.String()
		bounds(tag.sizeLB, tag.sizeUB, tag.sizeExt)
		return n
	case reflect.Bool:
		n.Kind, n.Value = "bool", strconv.FormatBool(v.Bool())
		return n
	}
	n.Kind, n.Value = "ext", fmt.Sprint(v.Interface())
	return n
}

// intRange is the range of an integer type without an aper tag (NAS).
func intRange(t reflect.Type) (int64, int64) {
	bits := t.Bits()
	switch t.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if bits >= 63 {
			return 0, 1<<63 - 1
		}
		return 0, 1<<uint(bits) - 1
	}
	return -1 << uint(bits-1), 1<<uint(bits-1) - 1
}

// octets renders octets: a PLMN, a nested message, or plain hex.
func (vb *viewer) octets(n *models.EditNode, path string, l loc, c ctx, tag aperTag, pt reflect.Type) *models.EditNode {
	b := l.v.Bytes()
	n.Value = hex.EncodeToString(b)
	if pt == plmnIdentityType {
		n.Kind, n.Hint = "plmn", plmn.String(b)
		return n
	}
	kind := containerKind(l)
	if kind == "" {
		n.Kind = "octets"
		n.Min, n.Max, n.Ext = tag.sizeLB, tag.sizeUB, tag.sizeExt
		n.Hint = octetCount(len(b))
		return n
	}
	n.Kind, n.Hint = "container", octetCount(len(b))
	if len(b) == 0 {
		return n
	}
	con, err := openContainer(l)
	if err != nil {
		n.Note = err.Error()
		return n
	}
	in := joinPath(path, ">")
	inner := c
	if con.kind != "transfer" {
		inner = ctx{nas: con.name, outer: c.nas, base: in}
	}
	root := vb.node(con.name, in, con.root, inner, n.Auto)
	if con.header != nil {
		hv := reflect.ValueOf(con.header).Elem()
		var hdr []*models.EditNode
		for i := 0; i < hv.NumField(); i++ {
			f := hv.Type().Field(i)
			if con.header.SecurityHeaderType == 0 && f.Name != "SecurityHeaderType" {
				continue // a plain NAS-PDU has only the type, to protect it
			}
			hn := vb.node("@"+f.Name, joinPath(in, "@"+f.Name), loc{v: hv.Field(i), parent: hv, field: f}, inner, n.Auto)
			if f.Name == "SecurityHeaderType" && int(con.header.SecurityHeaderType) < len(securityHeaders) {
				hn.Hint = securityHeaders[con.header.SecurityHeaderType]
			}
			hdr = append(hdr, hn)
		}
		root.Kids = append(hdr, root.Kids...)
	}
	n.Kids = []*models.EditNode{root}
	return n
}

// octetCount formats an octet count as a hint.
func octetCount(n int) string {
	if n == 1 {
		return "1 octet"
	}
	return fmt.Sprintf("%d octets", n)
}

// extensionText summarizes a read-only extension value.
func extensionText(v reflect.Value) string {
	if v.Kind() == reflect.Slice {
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return hex.EncodeToString(v.Bytes())
		}
		if v.Len() == 1 {
			return "1 item"
		}
		return fmt.Sprintf("%d items", v.Len())
	}
	return "present"
}
