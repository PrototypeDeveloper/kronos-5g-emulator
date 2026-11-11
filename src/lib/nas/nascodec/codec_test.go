package nascodec

import (
	"bytes"
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
		{"unknown TLVE", []byte{0x60, 0, 2, 0x7e, 0xff, 0x20, 1, 5}, false},
		{"release 17.5 optional 7c 7d", []byte{0x7c, 0, 0, 0x7d, 0, 0, 0x20, 1, 5}, false},
		{"duplicate first wins", []byte{0x20, 1, 5, 0x20, 1, 9}, false},
		{"ignore duplicate content length", []byte{0x20, 1, 5, 0x20, 0}, false},
		{"mandatory comprehension TLV", []byte{0x0f, 0}, true},
		{"mandatory comprehension TLVE 7e", []byte{0x7e, 0, 0}, true},
		{"mandatory comprehension TLVE 7f", []byte{0x7f, 0, 0}, true},
		{"unknown TLV missing length", []byte{0x40}, true},
		{"unknown TLVE short length", []byte{0x60, 0}, true},
		{"unknown TLV truncated", []byte{0x40, 2, 1}, true},
		{"unknown TLVE truncated", []byte{0x60, 0xff, 0xff, 1}, true},
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
	f.Add([]byte{0x60, 0, 1, 7, 0x20, 1, 5})
	f.Fuzz(func(t *testing.T, data []byte) {
		r := Reader{Data: data}
		got := r.Options(Schema{0x20: {Format: TLV, Min: 1, Max: 4}})
		if r.Err == nil && (len(r.Data) != 0 || len(got) > 1) {
			t.Fatal("unconsumed or duplicate IE")
		}
	})
}
