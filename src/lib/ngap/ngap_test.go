package libngap_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	ngap "libngap"
	"libngap/aper"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"
	"reflect"
	"testing"
)

const initialUEHex = "000f40480000050055000200020026001a197e004179000d0199f9590000000001000000602e04f0f0f0f0007900135099f959000000010099f959000001ee489aa5005a4001180070400100"

func sample(t testing.TB) []byte {
	t.Helper()
	b, e := hex.DecodeString(initialUEHex)
	if e != nil {
		t.Fatal(e)
	}
	return b
}

func TestInitialUERegistration(t *testing.T) {
	wire := sample(t)
	v, e := ngap.Decode(wire)
	if e != nil {
		t.Fatal(e)
	}
	m, ok := v.(*ngapMessage.InitialUEMessage)
	if !ok {
		t.Fatalf("decoded %T", v)
	}
	if m.RANUENGAPID.Value != 2 {
		t.Fatalf("RAN UE ID %d", m.RANUENGAPID.Value)
	}
	nas, _ := hex.DecodeString("7e004179000d0199f9590000000001000000602e04f0f0f0f0")
	if !bytes.Equal(m.NASPDU.Value, nas) {
		t.Fatalf("NAS %x", m.NASPDU.Value)
	}
	if m.RRCEstablishmentCause.Value != ngapType.RRCEstablishmentCausePresentMoSignalling {
		t.Fatalf("RRC cause %d", m.RRCEstablishmentCause.Value)
	}
	if !m.HasUEContextRequest || m.UEContextRequest.Value != ngapType.UEContextRequestPresentRequested {
		t.Fatal("missing UE Context Request")
	}
	if m.UserLocationInformation.Present != ngapType.UserLocationInformationPresentUserLocationInformationNR {
		t.Fatal("not NR location")
	}
	nr := m.UserLocationInformation.UserLocationInformationNR
	if !bytes.Equal(nr.NRCGI.PLMNIdentity.Value, []byte{0x99, 0xf9, 0x59}) || !bytes.Equal(nr.TAI.PLMNIdentity.Value, []byte{0x99, 0xf9, 0x59}) {
		t.Fatal("PLMN bytes")
	}
	if nr.NRCGI.NRCellIdentity.Value.BitLength != 36 || !bytes.Equal(nr.NRCGI.NRCellIdentity.Value.Bytes, []byte{0, 0, 0, 1, 0}) {
		t.Fatalf("NR cell identity %+v", nr.NRCGI.NRCellIdentity.Value)
	}
	if !bytes.Equal(nr.TAI.TAC.Value, []byte{0, 0, 1}) {
		t.Fatalf("TAC %x", nr.TAI.TAC.Value)
	}
	if !nr.HasTimeStamp || !bytes.Equal(nr.TimeStamp.Value, []byte{0xee, 0x48, 0x9a, 0xa5}) {
		t.Fatalf("timestamp %x", nr.TimeStamp.Value)
	}
	encoded, e := ngap.Encode(m)
	if e != nil {
		t.Fatal(e)
	}
	if !bytes.Equal(encoded, wire) {
		t.Fatalf("round trip\nwant %x\n got %x", wire, encoded)
	}
	// Received values own their byte slices; reusing the network buffer is safe.
	for i := range wire {
		wire[i] = 0
	}
	if !bytes.Equal(m.NASPDU.Value, nas) {
		t.Fatal("decoded NAS aliases input")
	}
}

func TestMalformedPDUs(t *testing.T) {
	b := sample(t)
	for n := 0; n < len(b); n++ {
		if v, e := ngap.Decode(b[:n]); e == nil || v != nil {
			t.Fatalf("accepted truncation at %d: %T %v", n, v, e)
		}
	}
	for _, bad := range [][]byte{append(append([]byte{}, b...), 0), {0x80}, {0x60, 0, 0, 0}, {0, 0xff, 0, 0}} {
		if v, e := ngap.Decode(bad); e == nil || v != nil {
			t.Fatalf("accepted bad PDU %x", bad)
		}
	}
	var nilSignal *ngapMessage.InitialUEMessage
	if _, e := ngap.Encode(nilSignal); e == nil {
		t.Fatal("accepted nil signal")
	}
	if _, e := ngap.Encode(42); e == nil {
		t.Fatal("accepted arbitrary value")
	}
}

func TestUnknownIEsAndExtensions(t *testing.T) {
	v, e := ngap.Decode(sample(t))
	if e != nil {
		t.Fatal(e)
	}
	m := v.(*ngapMessage.InitialUEMessage)
	m.AdditionalIEs = []ngapType.UnknownIE{{ID: ngapType.ProtocolIEID{Value: 65500}, Criticality: ngapType.Criticality{Value: 1}, Value: []byte{1, 2, 3}}}
	m.UnknownExtensions = []ngapType.UnknownExtension{{Index: 3, Value: []byte{0xff}}}
	b, e := ngap.Encode(m)
	if e != nil {
		t.Fatal(e)
	}
	got, e := ngap.Decode(b)
	if e != nil {
		t.Fatal(e)
	}
	m = got.(*ngapMessage.InitialUEMessage)
	if len(m.AdditionalIEs) != 1 || len(m.UnknownExtensions) != 1 || m.UnknownExtensions[0].Index != 3 {
		t.Fatal("unknown data lost")
	}
	again, e := ngap.Encode(m)
	if e != nil || !bytes.Equal(b, again) {
		t.Fatalf("unknown roundtrip %v", e)
	}
	m.AdditionalIEs[0].Criticality.Value = 2
	b, e = ngap.Encode(m)
	if e != nil {
		t.Fatal(e)
	}
	got, e = ngap.Decode(b)
	var notify *ngap.NotifyError
	if got == nil || !errors.As(e, &notify) || len(notify.Diagnostics) != 1 {
		t.Fatalf("notify: %T %v", got, e)
	}
	m.AdditionalIEs[0].Criticality.Value = 0
	b, e = ngap.Encode(m)
	if e != nil {
		t.Fatal(e)
	}
	got, e = ngap.Decode(b)
	var reject *ngap.CriticalityError
	if got != nil || !errors.As(e, &reject) {
		t.Fatalf("reject: %T %v", got, e)
	}
}

func FuzzDecode(f *testing.F) {
	b, _ := hex.DecodeString(initialUEHex)
	f.Add(b)
	f.Add([]byte{0})
	schema, _ := testSchema(f)
	for _, example := range messageExamples {
		v := reflect.New(reflect.TypeOf(example).Elem()).Elem()
		populate(f, schema, schema.Types[schema.Names[v.Type().Name()]], v, 0)
		seed, e := ngap.Encode(v.Addr().Interface())
		if e != nil {
			f.Fatal(e)
		}
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		v, e := ngap.Decode(b)
		if e == nil && v != nil {
			encoded, err := ngap.Encode(v)
			if err != nil {
				t.Fatalf("valid decoded value not encodable: %v", err)
			}
			if _, err = ngap.Decode(encoded); err != nil {
				t.Fatal(err)
			}
		}
	})
}

func TestIEValidation(t *testing.T) {
	type ie struct {
		id, criticality uint64
		value           []byte
	}
	h, e := ngapcodec.DecodeHeader(sample(t))
	if e != nil {
		t.Fatal(e)
	}
	r := aper.NewReader(h.Value)
	r.Bits(1)
	count := r.Constrained(0, 65535)
	var original []ie
	for i := uint64(0); i < count; i++ {
		original = append(original, ie{r.Constrained(0, 65535), r.Constrained(0, 2), r.OpenType(1000)})
	}
	if e = r.Finish(); e != nil {
		t.Fatal(e)
	}
	pack := func(ies []ie) []byte {
		body := aper.NewWriter(1000)
		body.Bits(0, 1)
		body.Constrained(uint64(len(ies)), 0, 65535)
		for _, x := range ies {
			body.Constrained(x.id, 0, 65535)
			body.Constrained(x.criticality, 0, 2)
			body.OpenType(x.value)
		}
		b, e := body.Bytes()
		if e != nil {
			t.Fatal(e)
		}
		w := aper.NewWriter(1000)
		w.Bits(0, 1)
		w.Constrained(h.Kind, 0, 2)
		w.Constrained(h.ProcedureCode, 0, 255)
		w.Constrained(h.Criticality, 0, 2)
		w.OpenType(b)
		b, e = w.Bytes()
		if e != nil {
			t.Fatal(e)
		}
		return b
	}
	cases := map[string][]ie{"missing mandatory": original[1:], "duplicate": append(append([]ie{}, original...), original[0])}
	badCriticality := append([]ie{}, original...)
	badCriticality[0].criticality = 1
	cases["criticality mismatch"] = badCriticality
	badPayload := append([]ie{}, original...)
	badPayload[3].value = []byte{0x18, 0}
	cases["trailing open-type data"] = badPayload
	badEnum := append([]ie{}, original...)
	badEnum[3].value = []byte{0x50}
	cases["invalid root enumerator"] = badEnum
	for name, ies := range cases {
		t.Run(name, func(t *testing.T) {
			if v, e := ngap.Decode(pack(ies)); e == nil || v != nil {
				t.Fatalf("accepted invalid message: %T %v", v, e)
			}
		})
	}
}

func TestParameterAPI(t *testing.T) {
	v := ngapType.RANUENGAPID{Value: 2}
	b, e := v.Encode()
	if e != nil {
		t.Fatal(e)
	}
	if !bytes.Equal(b, []byte{0, 2}) {
		t.Fatalf("RAN ID %x", b)
	}
	got, n, e := ngapType.DecodeRANUENGAPID(b)
	if e != nil || n != 2 || got.Value != 2 {
		t.Fatalf("parameter API: %+v %d %v", got, n, e)
	}
	if _, _, e = ngapType.DecodeRANUENGAPID(append(b, 0)); e == nil {
		t.Fatal("accepted parameter trailing data")
	}
}
