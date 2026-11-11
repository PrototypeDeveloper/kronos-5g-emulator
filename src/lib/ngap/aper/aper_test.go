package aper

import (
	"bytes"
	"encoding/hex"
	"math"
	"reflect"
	"testing"
)

// These vectors exercise the X.691 aligned constrained-integer cases, with
// independent expected octets rather than encode/decode symmetry alone.
func TestConstrainedIntegerVectors(t *testing.T) {
	cases := []struct {
		value, upper uint64
		want         string
	}{
		{0, 0, "00"}, {3, 7, "60"}, {254, 254, "fe"}, {255, 255, "ff"}, {256, 256, "0100"},
		{65535, 65535, "ffff"}, {2, math.MaxUint32, "0002"}, {65536, math.MaxUint32, "80010000"},
		{math.MaxUint64, math.MaxUint64, "e0ffffffffffffffff"},
	}
	for _, c := range cases {
		w := NewWriter(100)
		w.Constrained(c.value, 0, c.upper)
		b, e := w.Bytes()
		if e != nil {
			t.Fatal(e)
		}
		if hex.EncodeToString(b) != c.want {
			t.Fatalf("%d/%d: %x != %s", c.value, c.upper, b, c.want)
		}
		r := NewReader(b)
		if got := r.Constrained(0, c.upper); got != c.value {
			t.Fatal(got)
		}
		if e = r.Finish(); e != nil {
			t.Fatal(e)
		}
	}
	w := NewWriter(100)
	w.Bits(1, 1)
	w.Constrained(0xab, 0, 255)
	b, _ := w.Bytes()
	if !bytes.Equal(b, []byte{0x80, 0xab}) {
		t.Fatalf("alignment %x", b)
	}
}
func TestLengthDeterminants(t *testing.T) {
	for _, c := range []struct {
		n     int
		want  string
		count int
		more  bool
	}{{0, "00", 0, false}, {127, "7f", 127, false}, {128, "8080", 128, false}, {16383, "bfff", 16383, false}, {16384, "c1", 16384, true}, {65536, "c4", 65536, true}} {
		w := NewWriter(100)
		n, more := w.Length(c.n)
		b, _ := w.Bytes()
		if hex.EncodeToString(b) != c.want || n != c.count || more != c.more {
			t.Fatalf("length %d: %x %d %v", c.n, b, n, more)
		}
		r := NewReader(b)
		n, more = r.Length()
		if n != c.count || more != c.more || r.Finish() != nil {
			t.Fatal("length decode")
		}
	}
	for _, b := range []byte{0xc0, 0xc5, 0xff} {
		r := NewReader([]byte{b})
		r.Length()
		if r.Err() == nil {
			t.Fatal("accepted invalid fragment")
		}
	}
	for _, n := range []int{16383, 16384, 65536, 65537, 100000} {
		b := bytes.Repeat([]byte{0x5a}, n)
		w := NewWriter(200000)
		w.OpenType(b)
		encoded, e := w.Bytes()
		if e != nil {
			t.Fatal(e)
		}
		r := NewReader(encoded)
		got := r.OpenType(200000)
		if r.Finish() != nil || !bytes.Equal(b, got) {
			t.Fatal("fragmented open type")
		}
	}
}
func TestNormallySmall(t *testing.T) {
	for _, n := range []uint64{0, 63, 64, 128, 65536, math.MaxUint64} {
		w := NewWriter(100)
		w.NormallySmall(n)
		b, e := w.Bytes()
		if e != nil {
			t.Fatal(e)
		}
		r := NewReader(b)
		if got := r.NormallySmall(); got != n {
			t.Fatalf("%d != %d", got, n)
		}
		if e = r.Finish(); e != nil {
			t.Fatal(e)
		}
	}
}

func TestPrimitiveCodecs(t *testing.T) {
	s := Schema{Types: map[string]*Type{
		"Signed":    {Kind: "INTEGER"},
		"Union":     {Kind: "INTEGER", Constraints: []Constraint{{Kind: "value", Lower: "1", Upper: "40", Ranges: []ValueRange{{Lower: "1", Upper: "30"}, {Lower: "40", Upper: "40"}}}}},
		"Octets":    {Kind: "OCTET STRING"},
		"Printable": {Kind: "PrintableString", Constraints: []Constraint{{Kind: "size", Lower: "1", Upper: "150", Extensible: true}}},
		"UTF8":      {Kind: "UTF8String", Constraints: []Constraint{{Kind: "size", Lower: "1", Upper: "150", Extensible: true}}},
		"OID":       {Kind: "OBJECT IDENTIFIER"},
	}}
	c, e := NewCodec(s, DefaultLimits())
	if e != nil {
		t.Fatal(e)
	}
	for _, x := range []struct {
		n    int64
		want string
	}{{0, "0100"}, {127, "017f"}, {128, "020080"}, {-1, "01ff"}, {-128, "0180"}, {-129, "02ff7f"}} {
		b, e := c.Marshal("Signed", x.n)
		if e != nil || hex.EncodeToString(b) != x.want {
			t.Fatalf("signed %d: %x %v", x.n, b, e)
		}
		var got int64
		if e = c.Unmarshal("Signed", b, &got); e != nil || got != x.n {
			t.Fatal(got, e)
		}
	}
	if _, e = c.Marshal("Union", int64(35)); e == nil {
		t.Fatal("accepted hole in integer union")
	}
	if _, e = c.Marshal("Printable", "invalid@"); e == nil {
		t.Fatal("accepted bad alphabet")
	}
	if _, e = c.Marshal("UTF8", string([]byte{0xff})); e == nil {
		t.Fatal("accepted bad UTF-8")
	}
	for _, x := range []struct {
		name  string
		value any
	}{{"Octets", bytes.Repeat([]byte{0xaa}, 70000)}, {"Printable", "AMF-01"}, {"UTF8", "\u20ac"}, {"OID", []uint64{2, 999, 3, 128}}} {
		b, e := c.Marshal(x.name, x.value)
		if e != nil {
			t.Fatal(e)
		}
		dst := reflect.New(reflect.TypeOf(x.value))
		if e = c.Unmarshal(x.name, b, dst.Interface()); e != nil {
			t.Fatal(e)
		}
		if !reflect.DeepEqual(x.value, dst.Elem().Interface()) {
			t.Fatal("primitive mismatch")
		}
	}
}
func TestDecodeTransactionalAndLimits(t *testing.T) {
	c, e := NewCodec(Schema{Types: map[string]*Type{"Octets": {Kind: "OCTET STRING"}}}, Limits{MaxBytes: 20, MaxElements: 10, MaxDepth: 10, MaxAllocation: 40})
	if e != nil {
		t.Fatal(e)
	}
	dst := []byte{1, 2, 3}
	if e = c.Unmarshal("Octets", []byte{10, 0}, &dst); e == nil {
		t.Fatal("accepted truncation")
	}
	if !bytes.Equal(dst, []byte{1, 2, 3}) {
		t.Fatal("modified dst on failure")
	}
	if e = c.Unmarshal("Octets", make([]byte, 21), &dst); e == nil {
		t.Fatal("accepted oversized PDU")
	}
	if _, e = c.Marshal("Octets", make([]byte, 21)); e == nil {
		t.Fatal("encoded oversized value")
	}
}
