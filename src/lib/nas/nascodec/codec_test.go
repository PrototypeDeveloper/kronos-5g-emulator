package nascodec

import (
	"bytes"
	"libnas/nasType"
	"reflect"
	"testing"
)

func TestOptionalIEInteroperability(t *testing.T) {
	schema := Schema{0x20: {Format: TLV, Min: 1, Max: 2}, 0x70: {Format: TLVE, Min: 1, Max: 2}, 0x90: {Format: HalfTV, Min: 1, Max: 1}, 0x21: {Format: TV, Min: 2, Max: 2}}
	for _, tc := range []struct {
		name string
		wire []byte
		fail bool
	}{
		{"unknown half TV", []byte{0xf3, 0x20, 1, 5}, false},
		{"unknown TLV", []byte{0x40, 2, 0xff, 0x7e, 0x20, 1, 5}, false},
		{"unknown TLVE", []byte{0x71, 0, 2, 0x7e, 0xff, 0x20, 1, 5}, false},
		{"unknown TLV 6x", []byte{0x6b, 1, 0x21, 0x6e, 2, 0x7e, 0xff, 0x20, 1, 5}, false},
		{"release 17.5 optional 7c 7d", []byte{0x7c, 0, 0, 0x7d, 0, 0, 0x20, 1, 5}, false},
		{"duplicate first wins", []byte{0x20, 1, 5, 0x20, 1, 9}, false},
		{"ignore duplicate content length", []byte{0x20, 1, 5, 0x20, 0}, false},
		{"mandatory comprehension TLV", []byte{0x0f, 0}, true},
		{"mandatory comprehension TLVE 7e", []byte{0x7e, 0, 0}, true},
		{"mandatory comprehension TLVE 7f", []byte{0x7f, 0, 0}, true},
		{"unknown TLV missing length", []byte{0x40}, true},
		{"unknown TLVE short length", []byte{0x71, 0}, true},
		{"unknown TLV truncated", []byte{0x40, 2, 1}, true},
		{"unknown TLVE truncated", []byte{0x71, 0xff, 0xff, 1}, true},
		{"duplicate truncated", []byte{0x20, 1, 5, 0x20, 3, 1}, true},
		{"known invalid length", []byte{0x20, 0}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := Reader{Data: tc.wire}
			got := r.Options(schema)
			if (r.Err != nil) != tc.fail {
				t.Fatalf("options=%v err=%v", got, r.Err)
			}
			if !tc.fail && (len(got) != 1 || got[0].ID != 0x20 || !bytes.Equal(got[0].Value, []byte{5}) || len(r.Data) != 0) {
				t.Fatalf("wrong first IE: %v remaining=%x", got, r.Data)
			}
		})
	}
	// Repeated TV, half-TV and TLV-E also retain the first occurrence.
	r := Reader{Data: []byte{0x21, 1, 2, 0x21, 3, 4, 0x91, 0x92, 0x70, 0, 1, 5, 0x70, 0, 1, 6}}
	got := r.Options(schema)
	if r.Err != nil || len(got) != 3 || !bytes.Equal(got[0].Value, []byte{1, 2}) || got[1].Value[0] != 1 || got[2].Value[0] != 5 {
		t.Fatalf("%v %v", got, r.Err)
	}
}

func FuzzOptions(f *testing.F) {
	f.Add([]byte{0x40, 1, 7, 0x20, 1, 5})
	f.Add([]byte{0x71, 0, 1, 7, 0x20, 1, 5})
	f.Fuzz(func(t *testing.T, data []byte) {
		r := Reader{Data: data}
		got := r.Options(Schema{0x20: {Format: TLV, Min: 1, Max: 4}})
		if r.Err == nil && (len(r.Data) != 0 || len(got) > 1) {
			t.Fatal("unconsumed or duplicate IE")
		}
	})
}

func TestUnknownOptionsRoundTrip(t *testing.T) {
	schema := Schema{0x20: {Format: TLV, Min: 1, Max: 2}, 0x90: {Format: HalfTV, Min: 1, Max: 1}}
	// Unknown half-octet, TLV and TLV-E IEs around known ones are kept in order.
	wire := []byte{0xf3, 0x20, 1, 5, 0x40, 2, 0xaa, 0xbb, 0x91, 0x71, 0, 1, 0xcc, 0x7c, 0, 0}
	r := Reader{Data: wire}
	known, unknown := r.OptionsWithUnknown(schema)
	if r.Err != nil || len(known) != 2 || len(r.Data) != 0 {
		t.Fatalf("known=%v err=%v", known, r.Err)
	}
	want := []nasType.IE{{ID: 0xf0, Value: []byte{3}}, {ID: 0x40, Value: []byte{0xaa, 0xbb}}, {ID: 0x71, Value: []byte{0xcc}}, {ID: 0x7c}}
	if !reflect.DeepEqual(unknown, want) {
		t.Fatalf("unknown %v", unknown)
	}
	// Options still skips the same IEs.
	r = Reader{Data: wire}
	if got := r.Options(schema); r.Err != nil || !reflect.DeepEqual(got, known) {
		t.Fatalf("Options %v %v", got, r.Err)
	}
	w := Writer{}
	w.UnknownOptions(unknown, schema)
	if b, err := w.Result(); err != nil || !bytes.Equal(b, []byte{0xf3, 0x40, 2, 0xaa, 0xbb, 0x71, 0, 1, 0xcc, 0x7c, 0, 0}) {
		t.Fatalf("written %x %v", b, err)
	}
	for _, tc := range []struct {
		wire []byte
		fail bool
	}{
		{[]byte{0x40}, true},
		{[]byte{0x40, 2, 1}, true},
		{[]byte{0x71, 0}, true},
		{[]byte{0x71, 0, 2, 1}, true},
		{[]byte{0x0f, 0}, true},
		{[]byte{0x7e, 0, 0}, true},
		{[]byte{0x20, 1, 5, 0x20, 1, 6}, false}, // a duplicate known IE is not an unknown one
	} {
		r := Reader{Data: tc.wire}
		_, unknown := r.OptionsWithUnknown(schema)
		if (r.Err != nil) != tc.fail || (!tc.fail && len(unknown) != 0) {
			t.Fatalf("%x: unknown=%v err=%v", tc.wire, unknown, r.Err)
		}
	}
	// Unknown IEs that the decoder could not produce are refused.
	for _, ie := range []nasType.IE{
		{ID: 0x20, Value: []byte{1}},  // modelled
		{ID: 0x90, Value: []byte{1}},  // modelled half octet
		{ID: 0x0f, Value: []byte{1}},  // comprehension required
		{ID: 0x7e, Value: []byte{1}},  // comprehension required
		{ID: 0xe1, Value: []byte{1}},  // half octet with a low nibble in the IEI
		{ID: 0xe0, Value: []byte{16}}, // half octet overflow
		{ID: 0xe0},                    // half octet without a value
		{ID: 0x40, Value: make([]byte, 256)},
	} {
		w := Writer{}
		w.UnknownOptions([]nasType.IE{ie}, schema)
		if _, err := w.Result(); err == nil {
			t.Fatalf("unknown IE %02x %x accepted", ie.ID, ie.Value)
		}
	}
}
