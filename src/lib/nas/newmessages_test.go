package libnas_test

import (
	"bytes"
	"fmt"
	nas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// newMessageCase is a message with its mandatory fields set, its own decoder
// and the layout of its optional IEs in encoding order: "XX-" half octet,
// "XXL" TLV, "XXE" TLV-E, "XXTn" TV of n octets.
type newMessageCase struct {
	msg    any
	decode func([]byte) (any, error)
	layout string
}

// own adapts a message decoder to newMessageCase.decode.
func own[T any](f func([]byte) (*T, error)) func([]byte) (any, error) {
	return func(b []byte) (any, error) { return f(b) }
}

var eap = []byte{1, 1, 0, 5, 1}

func newMessageCases() []newMessageCase {
	return []newMessageCase{
		{&nasMessage.DeregistrationRequestNW{DeregistrationType: 1}, own(nasMessage.DecodeDeregistrationRequestNW), "58T2 5FL 6DL 75E"},
		{&nasMessage.DeregistrationAcceptNW{}, own(nasMessage.DecodeDeregistrationAcceptNW), ""},
		{&nasMessage.ServiceReject{Cause: 9}, own(nasMessage.DecodeServiceReject), "50L 5FL 78E 6BL 75E"},
		{&nasMessage.ControlPlaneServiceRequest{ControlPlaneServiceTypeAndNgKSI: 0x31}, own(nasMessage.DecodeControlPlaneServiceRequest), "6FL 8- 7BE 12T2 50L F- 40L 71E 24L"},
		{&nasMessage.NetworkSliceSpecificAuthenticationCommand{SNSSAI: []byte{1}, EAPMessage: eap}, own(nasMessage.DecodeNetworkSliceSpecificAuthenticationCommand), ""},
		{&nasMessage.NetworkSliceSpecificAuthenticationComplete{SNSSAI: []byte{1, 0, 0, 1}, EAPMessage: eap}, own(nasMessage.DecodeNetworkSliceSpecificAuthenticationComplete), ""},
		{&nasMessage.NetworkSliceSpecificAuthenticationResult{SNSSAI: []byte{1}, EAPMessage: []byte{3, 1, 0, 4}}, own(nasMessage.DecodeNetworkSliceSpecificAuthenticationResult), ""},
		{&nasMessage.ConfigurationUpdateCommand{}, own(nasMessage.DecodeConfigurationUpdateCommand), "D- 77E 54L 15L 27L 43L 45L 46T2 47T8 49L 79E B- 9- 31L 11L 76E F- 6CL 75E 67L A- 44L 1BL C-"},
		{&nasMessage.ConfigurationUpdateComplete{}, own(nasMessage.DecodeConfigurationUpdateComplete), ""},
		{&nasMessage.AuthenticationResult{NgKSI: 2, EAPMessage: []byte{3, 0, 0, 4}}, own(nasMessage.DecodeAuthenticationResult), "38L"},
		{&nasMessage.Notification{AccessType: 2}, own(nasMessage.DecodeNotification), ""},
		{&nasMessage.NotificationResponse{}, own(nasMessage.DecodeNotificationResponse), "50L"},
		{&nasMessage.PDUSessionAuthenticationCommand{PDUSessionID: 5, ProcedureTransactionID: 1, EAPMessage: eap}, own(nasMessage.DecodePDUSessionAuthenticationCommand), "7BE"},
		{&nasMessage.PDUSessionAuthenticationComplete{PDUSessionID: 5, ProcedureTransactionID: 1, EAPMessage: eap}, own(nasMessage.DecodePDUSessionAuthenticationComplete), "7BE"},
		{&nasMessage.PDUSessionAuthenticationResult{PDUSessionID: 5, ProcedureTransactionID: 1}, own(nasMessage.DecodePDUSessionAuthenticationResult), "78E 7BE"},
		{&nasMessage.PDUSessionModificationRequest{PDUSessionID: 5, ProcedureTransactionID: 1}, own(nasMessage.DecodePDUSessionModificationRequest), "28L 59T2 55T3 B- 13T3 7AE 79E 75E 7BE 74E 66L 1FL"},
		{&nasMessage.PDUSessionModificationReject{PDUSessionID: 5, ProcedureTransactionID: 1, Cause: 26}, own(nasMessage.DecodePDUSessionModificationReject), "37L 7BE 1DL 61L"},
		{&nasMessage.PDUSessionModificationCommand{PDUSessionID: 5, ProcedureTransactionID: 1}, own(nasMessage.DecodePDUSessionModificationCommand), "59T2 2AL 56T2 8- 7AE 75E 79E 7BE 77E 66L 74E 1EL 1FL"},
		{&nasMessage.PDUSessionModificationComplete{PDUSessionID: 5, ProcedureTransactionID: 1}, own(nasMessage.DecodePDUSessionModificationComplete), "7BE 74E"},
		{&nasMessage.PDUSessionModificationCommandReject{PDUSessionID: 5, ProcedureTransactionID: 1, Cause: 43}, own(nasMessage.DecodePDUSessionModificationCommandReject), "7BE"},
	}
}

// Unknown optional IEs (half octet, TLV, TLV-E) that no new message models.
var (
	testUnknownIEs  = []nasType.IE{{ID: 0xe0, Value: []byte{5}}, {ID: 0x41, Value: []byte{1, 2}}, {ID: 0x7c, Value: []byte{3, 4, 5}}}
	testUnknownWire = []byte{0xe5, 0x41, 2, 1, 2, 0x7c, 0, 3, 3, 4, 5}
)

// valueLength is the octets a test gives an optional IE value: the smallest
// valid length where the default of 2 is not valid.
var valueLength = map[string]int{
	"GUTI5GS": 11, "TAIList": 7, "ServiceAreaList": 4, "TimeZoneAndTime": 7, "EAPMessage": 4,
	"QoSRules": 4, "QoSFlowDescriptions": 3, "MappedEPSBearerContexts": 4,
	"PortManagementInformationContainer": 5, "IPHeaderCompressionConfiguration": 3, "SessionAMBR": 6,
}

// withAllOptions returns a copy of msg with every optional IE present, each
// with its own value, and with testUnknownIEs.
func withAllOptions(msg any) any {
	v := reflect.New(reflect.TypeOf(msg).Elem())
	v.Elem().Set(reflect.ValueOf(msg).Elem())
	s := v.Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Type().Field(i)
		if f.Type.Kind() != reflect.Bool || !strings.HasPrefix(f.Name, "Has") {
			continue
		}
		s.Field(i).SetBool(true)
		value := s.FieldByName(f.Name[3:]).Field(0)
		if value.Kind() == reflect.Uint8 {
			value.SetUint(uint64(i%15 + 1))
			continue
		}
		n := valueLength[s.FieldByName(f.Name[3:]).Type().Name()]
		if n == 0 {
			n = 2
		}
		value.SetBytes(bytes.Repeat([]byte{byte(i)}, n))
	}
	s.FieldByName("UnknownIEs").Set(reflect.ValueOf(append([]nasType.IE(nil), testUnknownIEs...)))
	return v.Interface()
}

// layoutIEI is the IEI of a layout token (a half octet one as 0xX0).
func layoutIEI(tok string) uint8 {
	if tok[1] == '-' {
		v, _ := strconv.ParseUint(tok[:1], 16, 8)
		return uint8(v) << 4
	}
	v, _ := strconv.ParseUint(tok[:2], 16, 8)
	return uint8(v)
}

// checkLayout walks the optional IEs of wire from offset by layout and
// returns the octets that follow them.
func checkLayout(t *testing.T, wire []byte, offset int, layout string) []byte {
	t.Helper()
	b := wire[offset:]
	for _, tok := range strings.Fields(layout) {
		if len(b) == 0 {
			t.Fatalf("%s: wire ends", tok)
		}
		if tok[1] == '-' {
			if b[0]&0xf0 != layoutIEI(tok) {
				t.Fatalf("%s: found %02x", tok, b[0])
			}
			b = b[1:]
			continue
		}
		if b[0] != layoutIEI(tok) {
			t.Fatalf("%s: found %02x", tok, b[0])
		}
		n := 0
		switch tok[2] {
		case 'T':
			n, _ = strconv.Atoi(tok[3:])
		case 'L':
			n = 2 + int(b[1])
		case 'E':
			n = 3 + int(b[1])<<8 + int(b[2])
		}
		if n > len(b) {
			t.Fatalf("%s: truncated", tok)
		}
		b = b[n:]
	}
	return b
}

func TestNewMessagesRoundTrip(t *testing.T) {
	for _, tc := range newMessageCases() {
		name := reflect.TypeOf(tc.msg).Elem().Name()
		t.Run(name, func(t *testing.T) {
			base, err := nas.Encode(tc.msg, nil)
			if err != nil {
				t.Fatal(err)
			}
			full := withAllOptions(tc.msg)
			wire, err := nas.Encode(full, nil)
			if err != nil {
				t.Fatal(err)
			}
			if rest := checkLayout(t, wire, len(base), tc.layout); !bytes.Equal(rest, testUnknownWire) {
				t.Fatalf("after the modelled IEs %x", rest)
			}
			got, err := nas.Decode(wire, nil)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, full) {
				t.Fatalf("decoded\n%#v\nwant\n%#v", got, full)
			}
			again, err := nas.Encode(got, nil)
			if err != nil || !bytes.Equal(again, wire) {
				t.Fatalf("re-encoded %x: %v", again, err)
			}
			if own, err := tc.decode(wire); err != nil || !reflect.DeepEqual(own, full) {
				t.Fatalf("own decoder: %v", err)
			}
			// Unknown IEs before the modelled ones come back too, and are
			// encoded after them.
			moved := append(append(append([]byte(nil), wire[:len(base)]...), testUnknownWire...), wire[len(base):len(wire)-len(testUnknownWire)]...)
			if got, err = nas.Decode(moved, nil); err != nil || !reflect.DeepEqual(got, full) {
				t.Fatalf("unknown IEs first: %v", err)
			}
			// The mandatory part alone decodes with nothing optional.
			if got, err = nas.Decode(base, nil); err != nil || !reflect.DeepEqual(got, tc.msg) {
				t.Fatalf("mandatory only: %#v %v", got, err)
			}
			// A cut never decodes as something else: it fails, or it ends
			// between IEs and encodes back to itself.
			for n := 0; n < len(wire); n++ {
				got, err := nas.Decode(wire[:n], nil)
				if n < len(base) && err == nil {
					t.Fatalf("cut %d inside the mandatory part accepted", n)
				}
				if err == nil {
					if b, _ := nas.Encode(got, nil); !bytes.Equal(b, wire[:n]) {
						t.Fatalf("cut %d decoded as %x", n, b)
					}
				}
			}
			if _, err = nas.Decode(append(append([]byte(nil), wire...), 0x7e, 0, 0), nil); err == nil {
				t.Fatal("comprehension-required IE ignored")
			}
		})
	}
}

func TestNewMessagesWrongHeader(t *testing.T) {
	for _, tc := range newMessageCases() {
		name := reflect.TypeOf(tc.msg).Elem().Name()
		wire, err := nas.Encode(withAllOptions(tc.msg), nil)
		if err != nil {
			t.Fatal(name, err)
		}
		typeAt := 2
		other := nasType.EPD5GSM
		if wire[0] == nasType.EPD5GSM {
			typeAt, other = 3, nasType.EPD5GMM
		}
		for _, change := range []func(b []byte){
			func(b []byte) { b[typeAt]++ },
			func(b []byte) { b[typeAt] = nasType.MsgStatus5GMM },
			func(b []byte) { b[0] = other },
			func(b []byte) { b[0] = 0x0e },
			func(b []byte) { b[1] = 0x10 }, // security header / PDU session ID out of range
		} {
			bad := append([]byte(nil), wire...)
			change(bad)
			if m, err := tc.decode(bad); err == nil || !reflect.ValueOf(m).IsNil() {
				t.Fatalf("%s: wrong header %x accepted", name, bad[:typeAt+1])
			}
		}
		if m, err := tc.decode(wire[:typeAt]); err == nil || !reflect.ValueOf(m).IsNil() {
			t.Fatalf("%s: short header accepted", name)
		}
	}
}

// A modelled IE in UnknownIEs is refused: it belongs in its named field.
func TestNewMessagesUnknownIEIsNotModelled(t *testing.T) {
	for _, tc := range newMessageCases() {
		toks := strings.Fields(tc.layout)
		if len(toks) == 0 {
			continue
		}
		id := layoutIEI(toks[0])
		full := withAllOptions(tc.msg)
		reflect.ValueOf(full).Elem().FieldByName("UnknownIEs").Set(reflect.ValueOf([]nasType.IE{{ID: id, Value: []byte{1}}}))
		if _, err := nas.Encode(full, nil); err == nil {
			t.Fatalf("%T: modelled IE %02x accepted as unknown", tc.msg, id)
		}
	}
}

func TestNewMessagesWireFixtures(t *testing.T) {
	for _, tc := range []struct{ name, hex string }{
		{"DeregistrationRequestNW", "7e004701580b5f01216d021001750000"},
		{"DeregistrationAcceptNW", "7e0048"},
		{"ServiceReject", "7e004d0a500204005f0121780004040100046b0101"},
		{"ControlPlaneServiceRequest", "7e004f316f020102817b0002aabb120550022000f140022000" + "7100037e0043240101"},
		{"NetworkSliceSpecificAuthenticationCommand", "7e0050010100050101000501"},
		{"NetworkSliceSpecificAuthenticationComplete", "7e0051010100050201000501"},
		{"NetworkSliceSpecificAuthenticationResult", "7e00520101000403010004"},
		{"ConfigurationUpdateCommand", "7e0054d143028041450280414623472201111230102349010" + "0f16c0121c1"},
		{"ConfigurationUpdateComplete", "7e0055"},
		{"AuthenticationResult", "7e005a00000403000004380200" + "00"},
		{"Notification", "7e006502"},
		{"NotificationResponse", "7e006650020400"},
		{"PDUSessionAuthenticationCommand", "2e0501c5000501010005017b000180"},
		{"PDUSessionAuthenticationComplete", "2e0501c600050201000501"},
		{"PDUSessionAuthenticationResult", "2e0501c778000403000004" + "7b000180"},
		{"PDUSessionModificationRequest", "2e0501c92801015924550010b113ffff7b000180"},
		{"PDUSessionModificationReject", "2e0501ca1a3701217b0001801d0101610101"},
		{"PDUSessionModificationCommand", "2e0501cb59242a06060064060064562181" + "7a000401000131" + "1e020001"},
		{"PDUSessionModificationComplete", "2e0501cc7b000180"},
		{"PDUSessionModificationCommandReject", "2e0501cd2b7b000180"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire := unhex(tc.hex)
			msg, err := nas.Decode(wire, nil)
			if err != nil {
				t.Fatal(err)
			}
			if got := fmt.Sprintf("%T", msg); got != "*nasMessage."+tc.name {
				t.Fatalf("type %s", got)
			}
			if v := reflect.ValueOf(msg).Elem().FieldByName("UnknownIEs"); v.Len() != 0 {
				t.Fatalf("modelled IEs taken as unknown: %v", v)
			}
			b, err := nas.Encode(msg, nil)
			if err != nil || !bytes.Equal(b, wire) {
				t.Fatalf("encoded %x: %v", b, err)
			}
		})
	}
	// Spot checks of the mandatory half octets and of field assignment.
	m, _ := nas.Decode(unhex("7e004f316f020102817b0002aabb"), nil)
	cp := m.(*nasMessage.ControlPlaneServiceRequest)
	if cp.ControlPlaneServiceTypeAndNgKSI != 0x31 || !cp.HasPayloadContainerType || cp.PayloadContainerType.Value != 1 || !bytes.Equal(cp.PayloadContainer.Value, []byte{0xaa, 0xbb}) {
		t.Fatalf("%+v", cp)
	}
	m, _ = nas.Decode(unhex("2e0501cb59242a06060064060064562181"), nil)
	mc := m.(*nasMessage.PDUSessionModificationCommand)
	if mc.PDUSessionID != 5 || mc.ProcedureTransactionID != 1 || mc.Cause5GSM.Value != 0x24 || mc.RQTimerValue.Value != 0x21 || mc.AlwaysOnPDUSessionIndication.Value != 1 || len(mc.SessionAMBR.Value) != 6 {
		t.Fatalf("%+v", mc)
	}
}

func TestNewMessagesMandatoryBounds(t *testing.T) {
	for _, msg := range []any{
		&nasMessage.DeregistrationRequestNW{DeregistrationType: 16},
		&nasMessage.AuthenticationResult{NgKSI: 16, EAPMessage: []byte{3, 0, 0, 4}},
		&nasMessage.AuthenticationResult{EAPMessage: []byte{3, 0, 0}},
		&nasMessage.AuthenticationResult{EAPMessage: make([]byte, 1501)},
		&nasMessage.Notification{AccessType: 4},
		&nasMessage.NetworkSliceSpecificAuthenticationCommand{EAPMessage: eap},
		&nasMessage.NetworkSliceSpecificAuthenticationCommand{SNSSAI: make([]byte, 9), EAPMessage: eap},
		&nasMessage.PDUSessionAuthenticationCommand{EAPMessage: nil},
		&nasMessage.PDUSessionModificationCommand{PDUSessionID: 16},
		&nasMessage.PDUSessionModificationCommand{HasSessionAMBR: true, SessionAMBR: nasType.SessionAMBR{Value: make([]byte, 5)}},
		&nasMessage.ConfigurationUpdateCommand{HasSMSIndication: true, SMSIndication: nasType.SMSIndication{Value: 16}},
		&nasMessage.ConfigurationUpdateCommand{HasUniversalTimeAndLocalTimeZone: true, UniversalTimeAndLocalTimeZone: nasType.TimeZoneAndTime{Value: make([]byte, 6)}},
	} {
		if b, err := nas.Encode(msg, nil); err == nil {
			t.Fatalf("%T encoded out of bounds: %x", msg, b)
		}
	}
}

// 5GSM messages travel in UL/DL NAS Transport; 5GMM ones are protected.
func TestNewMessagesInTransportAndProtected(t *testing.T) {
	for _, tc := range newMessageCases() {
		full := withAllOptions(tc.msg)
		inner, err := nas.Encode(full, nil)
		if err != nil {
			t.Fatal(err)
		}
		if inner[0] == nasType.EPD5GSM {
			for _, outer := range []any{
				&nasMessage.DLNASTransport{PayloadContainerType: 1, PayloadContainer: inner, HasPDUSessionID: true, PDUSessionID: nasType.PDUSessionID{Value: 5}},
				&nasMessage.ULNASTransport{PayloadContainerType: 1, PayloadContainer: inner, HasPDUSessionID: true, PDUSessionID: nasType.PDUSessionID{Value: 5}},
			} {
				tx, rx := securityContext(), securityContext()
				wire, err := nas.Encode(outer, &nas.EncodeContext{Security: &tx, SecurityHeaderType: nas.SecurityHeaderIntegrityProtectedAndCiphered})
				if err != nil {
					t.Fatal(err)
				}
				got, err := nas.Decode(wire, &nas.DecodeContext{Security: &rx})
				if err != nil {
					t.Fatal(err)
				}
				var container []byte
				switch m := got.(type) {
				case *nasMessage.DLNASTransport:
					container = m.PayloadContainer
				case *nasMessage.ULNASTransport:
					container = m.PayloadContainer
				}
				sm, err := nas.Decode(container, nil)
				if err != nil || !reflect.DeepEqual(sm, full) {
					t.Fatalf("%T in %T: %v", full, outer, err)
				}
			}
			continue
		}
		for _, d := range []nas.Direction{nas.Uplink, nas.Downlink} {
			tx, rx := securityContext(), securityContext()
			wire, err := nas.Encode(full, &nas.EncodeContext{Security: &tx, SecurityHeaderType: nas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: d})
			if err != nil {
				t.Fatal(err)
			}
			got, err := nas.Decode(wire, &nas.DecodeContext{Security: &rx, Direction: d})
			if err != nil || !reflect.DeepEqual(got, full) {
				t.Fatalf("protected %T: %v", full, err)
			}
		}
	}
}
