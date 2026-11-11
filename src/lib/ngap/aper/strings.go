package aper

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"unicode/utf8"
)

func sizeBounds(c *Constraint) (uint64, uint64, error) {
	lo, hi, e := bounds(c)
	if e != nil {
		return 0, 0, e
	}
	if !lo.IsUint64() || !hi.IsUint64() {
		return 0, 0, fmt.Errorf("invalid SIZE bounds")
	}
	return lo.Uint64(), hi.Uint64(), nil
}
func validSize(c *Constraint, n int) bool {
	if c == nil {
		return true
	}
	lo, hi, e := sizeBounds(c)
	if e != nil || uint64(n) < lo || uint64(n) > hi {
		return false
	}
	return len(c.Ranges) == 0 || inRanges(big.NewInt(int64(n)), c.Ranges)
}

func (s *state) writeSized(w *Writer, c *Constraint, n int, emit func(int, int, *Constraint) error) error {
	if n < 0 {
		return fmt.Errorf("negative size")
	}
	effective := c
	if c != nil {
		root := validSize(c, n)
		if c.Extensible {
			if root {
				w.Bits(0, 1)
			} else {
				w.Bits(1, 1)
				effective = nil
			}
		} else if !root {
			return fmt.Errorf("size %d outside %s..%s", n, c.Lower, c.Upper)
		}
	}
	if effective != nil {
		lo, hi, e := sizeBounds(effective)
		if e != nil {
			return e
		}
		if hi < 65536 {
			w.Constrained(uint64(n), lo, hi)
			if e = emit(0, n, effective); e != nil {
				return e
			}
			return w.Err()
		}
	}
	off := 0
	for {
		count, more := w.Length(n - off)
		if e := emit(off, count, effective); e != nil {
			return e
		}
		off += count
		if w.Err() != nil || !more {
			return w.Err()
		}
	}
}
func (s *state) readSized(r *Reader, c *Constraint, emit func(int, *Constraint) error) error {
	effective := c
	if c != nil && c.Extensible && r.Bits(1) != 0 {
		effective = nil
	}
	if effective != nil {
		lo, hi, e := sizeBounds(effective)
		if e != nil {
			return e
		}
		if hi < 65536 {
			n := int(r.Constrained(lo, hi))
			if r.Err() != nil {
				return r.Err()
			}
			if !validSize(effective, n) {
				return fmt.Errorf("invalid SIZE value %d", n)
			}
			return emit(n, effective)
		}
	}
	total := 0
	for {
		n, more := r.Length()
		if r.Err() != nil {
			return r.Err()
		}
		if n > s.codec.limits.MaxBytes*8-total {
			return fmt.Errorf("fragmented value exceeds limit")
		}
		total += n
		if e := emit(n, effective); e != nil {
			return e
		}
		if !more {
			if effective != nil && !validSize(effective, total) {
				return fmt.Errorf("invalid SIZE value %d", total)
			}
			return r.Err()
		}
	}
}
func stringAligned(kind string, c *Constraint) bool {
	if c == nil {
		return true
	}
	lo, hi, e := sizeBounds(c)
	if e != nil {
		return true
	}
	if kind == "BIT STRING" {
		return !(lo == hi && hi <= 16)
	}
	if kind == "OCTET STRING" {
		return !(lo == hi && hi <= 2)
	}
	if lo == hi {
		return hi > 2
	}
	return hi >= 2
}
func validateAlphabet(kind string, b []byte) error {
	switch kind {
	case "UTF8String":
		if !utf8.Valid(b) {
			return fmt.Errorf("invalid UTF-8")
		}
	case "VisibleString":
		for _, c := range b {
			if c < 32 || c > 126 {
				return fmt.Errorf("invalid VisibleString character")
			}
		}
	case "PrintableString":
		for _, c := range b {
			if !(c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || strings.ContainsRune(" '()+,-./:=?", rune(c))) {
				return fmt.Errorf("invalid PrintableString character")
			}
		}
	}
	return nil
}
func (s *state) encodeString(w *Writer, t *Type, v reflect.Value) error {
	v = primitive(v)
	var b []byte
	var length int
	c := constraint(t, "size")
	switch t.Kind {
	case "BIT STRING":
		f, e := field(v, "Bytes")
		if e != nil {
			return e
		}
		l, e := field(v, "BitLength")
		if e != nil {
			return e
		}
		if l.Uint() > uint64(s.codec.limits.MaxBytes)*8 {
			return fmt.Errorf("bit string exceeds limit")
		}
		b = f.Bytes()
		length = int(l.Uint())
		if len(b) != (length+7)/8 {
			return fmt.Errorf("BIT STRING byte count does not match BitLength")
		}
		if length%8 != 0 && b[len(b)-1]&byte((1<<(8-length%8))-1) != 0 {
			return fmt.Errorf("nonzero unused BIT STRING bits")
		}
	case "OCTET STRING":
		if v.Kind() != reflect.Slice || v.Type().Elem().Kind() != reflect.Uint8 {
			return fmt.Errorf("expected octet slice")
		}
		b = v.Bytes()
		length = len(b)
	case "OBJECT IDENTIFIER":
		var e error
		b, e = encodeOID(v)
		if e != nil {
			return e
		}
		length = len(b)
		c = nil
	default:
		if v.Kind() != reflect.String {
			return fmt.Errorf("expected string")
		}
		b = []byte(v.String())
		length = len(b)
		if e := validateAlphabet(t.Kind, b); e != nil {
			return e
		}
		if t.Kind == "UTF8String" {
			if c != nil && !c.Extensible && !validSize(c, utf8.RuneCount(b)) {
				return fmt.Errorf("UTF8String violates SIZE")
			}
			c = nil
		}
	}
	if len(b) > s.codec.limits.MaxBytes {
		return fmt.Errorf("string exceeds limit")
	}
	return s.writeSized(w, c, length, func(off, n int, c *Constraint) error {
		if n > 0 && stringAligned(t.Kind, c) {
			w.Align()
		}
		if t.Kind == "BIT STRING" {
			for i := off; i < off+n; i++ {
				w.Bits(uint64((b[i/8]>>uint(7-i%8))&1), 1)
			}
		} else {
			for _, x := range b[off : off+n] {
				w.Bits(uint64(x), 8)
			}
		}
		return w.Err()
	})
}
func (s *state) decodeString(r *Reader, t *Type, v reflect.Value) error {
	v = primitive(v)
	var b []byte
	total := 0
	c := constraint(t, "size")
	if t.Kind == "UTF8String" || t.Kind == "OBJECT IDENTIFIER" {
		c = nil
	}
	e := s.readSized(r, c, func(n int, c *Constraint) error {
		bytesN := n
		if t.Kind == "BIT STRING" {
			bytesN = (n + 7) / 8
		}
		if bytesN > s.codec.limits.MaxBytes-len(b) {
			return fmt.Errorf("string exceeds limit")
		}
		if e := s.allocate(1, bytesN); e != nil {
			return e
		}
		if n > 0 && stringAligned(t.Kind, c) {
			r.Align()
		}
		old := len(b)
		b = append(b, make([]byte, bytesN)...)
		if t.Kind == "BIT STRING" {
			for i := 0; i < n; i++ {
				b[old+i/8] |= byte(r.Bits(1)) << uint(7-i%8)
			}
		} else {
			for i := old; i < len(b); i++ {
				b[i] = byte(r.Bits(8))
			}
		}
		total += n
		return r.Err()
	})
	if e != nil {
		return e
	}
	switch t.Kind {
	case "BIT STRING":
		f, e := field(v, "Bytes")
		if e != nil {
			return e
		}
		l, e := field(v, "BitLength")
		if e != nil {
			return e
		}
		f.SetBytes(b)
		l.SetUint(uint64(total))
	case "OCTET STRING":
		if v.Kind() != reflect.Slice || v.Type().Elem().Kind() != reflect.Uint8 {
			return fmt.Errorf("expected octet slice")
		}
		v.SetBytes(b)
	case "OBJECT IDENTIFIER":
		return s.decodeOID(b, v)
	default:
		if v.Kind() != reflect.String {
			return fmt.Errorf("expected string")
		}
		if e := validateAlphabet(t.Kind, b); e != nil {
			return e
		}
		if t.Kind == "UTF8String" {
			c := constraint(t, "size")
			if c != nil && !c.Extensible && !validSize(c, utf8.RuneCount(b)) {
				return fmt.Errorf("UTF8String violates SIZE")
			}
		}
		v.SetString(string(b))
	}
	return nil
}
func encodeOID(v reflect.Value) ([]byte, error) {
	if v.Kind() != reflect.Slice || v.Len() < 2 {
		return nil, fmt.Errorf("OID requires at least two arcs")
	}
	a, b := v.Index(0).Uint(), v.Index(1).Uint()
	if a > 2 || (a < 2 && b > 39) || b > ^uint64(0)-a*40 {
		return nil, fmt.Errorf("invalid OID arcs")
	}
	var out []byte
	for i := 1; i < v.Len(); i++ {
		n := v.Index(i).Uint()
		if i == 1 {
			n = a*40 + b
		}
		var tmp [10]byte
		p := len(tmp) - 1
		tmp[p] = byte(n & 127)
		for n >>= 7; n > 0; n >>= 7 {
			p--
			tmp[p] = byte(n&127) | 128
		}
		out = append(out, tmp[p:]...)
	}
	return out, nil
}
func (s *state) decodeOID(b []byte, v reflect.Value) error {
	if v.Kind() != reflect.Slice || v.Type().Elem().Kind() != reflect.Uint64 {
		return fmt.Errorf("expected OID arc slice")
	}
	var arcs []uint64
	var n uint64
	start := true
	for _, x := range b {
		if start && x == 128 {
			return fmt.Errorf("nonminimal OID arc")
		}
		if n > ^uint64(0)>>7 {
			return fmt.Errorf("OID arc overflow")
		}
		n = n<<7 | uint64(x&127)
		start = false
		if x&128 == 0 {
			count := 1
			if len(arcs) == 0 {
				count = 2
			}
			if len(arcs)+count > s.codec.limits.MaxElements {
				return fmt.Errorf("OID arc count exceeds limit")
			}
			if e := s.allocate(8, count); e != nil {
				return e
			}
			if len(arcs) == 0 {
				a := n / 40
				if a > 2 {
					a = 2
				}
				arcs = append(arcs, a, n-a*40)
			} else {
				arcs = append(arcs, n)
			}
			n = 0
			start = true
		}
	}
	if !start || len(arcs) < 2 {
		return fmt.Errorf("invalid OID contents")
	}
	out := reflect.MakeSlice(v.Type(), len(arcs), len(arcs))
	for i, n := range arcs {
		out.Index(i).SetUint(n)
	}
	v.Set(out)
	return nil
}
