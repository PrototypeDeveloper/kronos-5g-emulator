package aper

import (
	"fmt"
	"math/big"
	"reflect"
	"sort"
)

func encodeInteger(w *Writer, t *Type, v reflect.Value) error {
	n, e := integerValue(v)
	if e != nil {
		return e
	}
	c := constraint(t, "value")
	if c != nil {
		lo, hi, e := bounds(c)
		if e != nil {
			return e
		}
		root := n.Cmp(lo) >= 0 && n.Cmp(hi) <= 0
		if len(c.Ranges) > 0 {
			root = root && inRanges(n, c.Ranges)
		}
		if c.Extensible {
			if root {
				w.Bits(0, 1)
			} else {
				w.Bits(1, 1)
			}
		} else if !root {
			return fmt.Errorf("integer %s violates root constraint", n)
		}
		if root {
			span := new(big.Int).Sub(hi, lo)
			offset := new(big.Int).Sub(n, lo)
			if !span.IsUint64() {
				return fmt.Errorf("integer range exceeds 64 bits")
			}
			w.Constrained(offset.Uint64(), 0, span.Uint64())
			return w.Err()
		}
	}
	// Unconstrained INTEGER uses the minimum signed two's-complement octets.
	var b []byte
	if n.Sign() >= 0 {
		b = n.Bytes()
		if len(b) == 0 {
			b = []byte{0}
		}
		if b[0]&0x80 != 0 {
			b = append([]byte{0}, b...)
		}
	} else {
		length := 1
		for {
			min := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), uint(length*8-1)))
			if n.Cmp(min) >= 0 {
				break
			}
			length++
		}
		b = new(big.Int).Add(new(big.Int).Lsh(big.NewInt(1), uint(length*8)), n).FillBytes(make([]byte, length))
	}
	w.OpenType(b)
	return w.Err()
}
func decodeInteger(r *Reader, t *Type, v reflect.Value) error {
	c := constraint(t, "value")
	ext := false
	if c != nil && c.Extensible {
		ext = r.Bits(1) != 0
	}
	var n *big.Int
	if c != nil && !ext {
		lo, hi, e := bounds(c)
		if e != nil {
			return e
		}
		span := new(big.Int).Sub(hi, lo)
		if !span.IsUint64() {
			return fmt.Errorf("integer range exceeds 64 bits")
		}
		n = new(big.Int).Add(lo, new(big.Int).SetUint64(r.Constrained(0, span.Uint64())))
		if len(c.Ranges) > 0 && !inRanges(n, c.Ranges) {
			return fmt.Errorf("integer %s violates root constraint", n)
		}
	} else {
		b := r.OpenType(9)
		if r.Err() != nil {
			return r.Err()
		}
		if len(b) == 0 {
			return fmt.Errorf("empty integer")
		}
		n = new(big.Int).SetBytes(b)
		if b[0]&0x80 != 0 {
			n.Sub(n, new(big.Int).Lsh(big.NewInt(1), uint(len(b)*8)))
		}
	}
	if r.Err() != nil {
		return r.Err()
	}
	return setInteger(v, n)
}
func enumValues(t *Type) (root, ext []EnumValue) {
	for _, v := range t.Values {
		if v.Extension {
			ext = append(ext, v)
		} else {
			root = append(root, v)
		}
	}
	sort.Slice(root, func(i, j int) bool { return root[i].Number < root[j].Number })
	return
}
func unknownEnumBase(t *Type) uint64 {
	var base uint64
	for _, v := range t.Values {
		if v.Number >= 0 && uint64(v.Number) >= base {
			base = uint64(v.Number) + 1
		}
	}
	return base
}
func encodeEnum(w *Writer, t *Type, v reflect.Value) error {
	n, e := integerValue(v)
	if e != nil {
		return e
	}
	if !n.IsUint64() {
		return fmt.Errorf("negative enumerator")
	}
	num := n.Uint64()
	root, ext := enumValues(t)
	for i, x := range root {
		if uint64(x.Number) == num {
			if t.Extensible {
				w.Bits(0, 1)
			}
			w.Constrained(uint64(i), 0, uint64(len(root)-1))
			return w.Err()
		}
	}
	if !t.Extensible {
		return fmt.Errorf("unknown non-extensible enumerator %d", num)
	}
	w.Bits(1, 1)
	for i, x := range ext {
		if uint64(x.Number) == num {
			w.NormallySmall(uint64(i))
			return w.Err()
		}
	}
	base := unknownEnumBase(t)
	if num < base {
		return fmt.Errorf("undefined enumerator %d", num)
	}
	w.NormallySmall(num - base + uint64(len(ext)))
	return w.Err()
}
func decodeEnum(r *Reader, t *Type, v reflect.Value) error {
	root, ext := enumValues(t)
	var n uint64
	if t.Extensible && r.Bits(1) != 0 {
		i := r.NormallySmall()
		if i < uint64(len(ext)) {
			n = uint64(ext[i].Number)
		} else {
			base := unknownEnumBase(t)
			delta := i - uint64(len(ext))
			if ^uint64(0)-base < delta {
				return fmt.Errorf("enumerator overflow")
			}
			n = base + delta
		}
	} else {
		if len(root) == 0 {
			return fmt.Errorf("empty enum root")
		}
		i := r.Constrained(0, uint64(len(root)-1))
		if r.Err() != nil {
			return r.Err()
		}
		n = uint64(root[i].Number)
	}
	if r.Err() != nil {
		return r.Err()
	}
	return setInteger(v, new(big.Int).SetUint64(n))
}
