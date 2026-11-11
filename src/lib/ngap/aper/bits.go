package aper

import (
	"fmt"
	"io"
	"math/bits"
)

// Reader reads aligned PER bit fields. Errors are sticky; inspect Err/Finish.
type Reader struct {
	data []byte
	pos  uint64
	err  error
}

func NewReader(data []byte) *Reader { return &Reader{data: data} }
func (r *Reader) Err() error        { return r.err }
func (r *Reader) Fail(err error) {
	if r.err == nil {
		r.err = err
	}
}
func (r *Reader) Bits(n uint) uint64 {
	if r.err != nil {
		return 0
	}
	if n > 64 || uint64(n) > uint64(len(r.data))*8-r.pos {
		r.Fail(io.ErrUnexpectedEOF)
		return 0
	}
	var v uint64
	for i := uint(0); i < n; i++ {
		v = v<<1 | uint64((r.data[r.pos/8]>>(7-r.pos%8))&1)
		r.pos++
	}
	return v
}
func (r *Reader) Align() {
	if n := (8 - r.pos%8) % 8; n != 0 && r.Bits(uint(n)) != 0 {
		r.Fail(fmt.Errorf("nonzero APER alignment padding"))
	}
}
func (r *Reader) Octets(n int) []byte {
	r.Align()
	if r.err != nil {
		return nil
	}
	if n < 0 || uint64(n) > uint64(len(r.data))-r.pos/8 {
		r.Fail(io.ErrUnexpectedEOF)
		return nil
	}
	v := append([]byte(nil), r.data[r.pos/8:r.pos/8+uint64(n)]...)
	r.pos += uint64(n) * 8
	return v
}
func (r *Reader) Finish() error {
	if len(r.data) == 0 {
		return io.ErrUnexpectedEOF
	}
	if r.err != nil {
		return r.err
	}
	remaining := uint64(len(r.data))*8 - r.pos
	// A zero-bit open type is represented by one zero octet.
	if r.pos == 0 && remaining == 8 && r.data[0] == 0 {
		r.pos = 8
		return nil
	}
	if remaining > 7 {
		return fmt.Errorf("trailing APER data (%d bits)", remaining)
	}
	r.Align()
	return r.err
}
func (r *Reader) Constrained(lo, hi uint64) uint64 {
	if hi < lo {
		r.Fail(fmt.Errorf("invalid integer bounds"))
		return 0
	}
	span := hi - lo
	if span == 0 {
		return lo
	}
	var v uint64
	switch {
	case span < 255:
		v = r.Bits(uint(bits.Len64(span)))
	case span == 255:
		r.Align()
		v = r.Bits(8)
	case span <= 65535:
		r.Align()
		v = r.Bits(16)
	default:
		n := r.Constrained(1, uint64((bits.Len64(span)+7)/8))
		r.Align()
		v = r.Bits(uint(n) * 8)
	}
	if v > span {
		r.Fail(fmt.Errorf("integer offset %d exceeds range %d", v, span))
	}
	return lo + v
}

// Length returns one unconstrained length determinant and whether a fragment follows.
func (r *Reader) Length() (int, bool) {
	r.Align()
	b := r.Bits(8)
	if b < 128 {
		return int(b), false
	}
	if b < 192 {
		return int((b&63)<<8 | r.Bits(8)), false
	}
	m := b & 63
	if m < 1 || m > 4 {
		r.Fail(fmt.Errorf("invalid APER fragment multiplier %d", m))
		return 0, false
	}
	return int(m) * 16384, true
}
func (r *Reader) OpenType(max int) []byte {
	var out []byte
	for {
		n, more := r.Length()
		if n > max-len(out) {
			r.Fail(fmt.Errorf("open type exceeds %d bytes", max))
			return nil
		}
		out = append(out, r.Octets(n)...)
		if r.err != nil || !more {
			if len(out) == 0 {
				r.Fail(fmt.Errorf("empty APER open type"))
			}
			return out
		}
	}
}
func (r *Reader) NormallySmall() uint64 {
	if r.Bits(1) == 0 {
		return r.Bits(6)
	}
	b := r.OpenType(8)
	if len(b) == 0 {
		r.Fail(fmt.Errorf("empty normally-small integer"))
	}
	var v uint64
	for _, x := range b {
		v = v<<8 | uint64(x)
	}
	return v
}

// Writer constructs aligned PER values with a bounded output buffer.
type Writer struct {
	data []byte
	pos  uint64
	err  error
	max  int
}

func NewWriter(max int) *Writer {
	if max <= 0 {
		max = 16 << 20
	}
	return &Writer{max: max}
}
func (w *Writer) Fail(err error) {
	if w.err == nil {
		w.err = err
	}
}
func (w *Writer) Err() error { return w.err }
func (w *Writer) Bits(v uint64, n uint) {
	if w.err != nil {
		return
	}
	if n > 64 || (n < 64 && v>>n != 0) {
		w.Fail(fmt.Errorf("value %d does not fit %d bits", v, n))
		return
	}
	if w.pos+uint64(n) > uint64(w.max)*8 {
		w.Fail(fmt.Errorf("APER output exceeds %d bytes", w.max))
		return
	}
	for i := int(n) - 1; i >= 0; i-- {
		if w.pos/8 == uint64(len(w.data)) {
			w.data = append(w.data, 0)
		}
		w.data[w.pos/8] |= byte((v>>i)&1) << (7 - w.pos%8)
		w.pos++
	}
}
func (w *Writer) Align() { w.Bits(0, uint((8-w.pos%8)%8)) }
func (w *Writer) Octets(b []byte) {
	w.Align()
	if w.err != nil {
		return
	}
	if len(b) > w.max-len(w.data) {
		w.Fail(fmt.Errorf("APER output exceeds %d bytes", w.max))
		return
	}
	w.data = append(w.data, b...)
	w.pos += uint64(len(b)) * 8
}
func (w *Writer) Bytes() ([]byte, error) {
	w.Align()
	if w.err != nil {
		return nil, w.err
	}
	if len(w.data) == 0 {
		return []byte{0}, nil
	}
	return w.data, nil
}
func (w *Writer) Constrained(v, lo, hi uint64) {
	if hi < lo || v < lo || v > hi {
		w.Fail(fmt.Errorf("integer %d outside %d..%d", v, lo, hi))
		return
	}
	span := hi - lo
	v -= lo
	switch {
	case span == 0:
		return
	case span < 255:
		w.Bits(v, uint(bits.Len64(span)))
	case span == 255:
		w.Align()
		w.Bits(v, 8)
	case span <= 65535:
		w.Align()
		w.Bits(v, 16)
	default:
		n := (bits.Len64(v) + 7) / 8
		if n == 0 {
			n = 1
		}
		w.Constrained(uint64(n), 1, uint64((bits.Len64(span)+7)/8))
		w.Align()
		w.Bits(v, uint(n)*8)
	}
}

// Length writes one determinant. A fragment consumes 16K..64K units.
func (w *Writer) Length(n int) (int, bool) {
	w.Align()
	if n < 0 {
		w.Fail(fmt.Errorf("negative length"))
		return 0, false
	}
	if n < 128 {
		w.Bits(uint64(n), 8)
		return n, false
	}
	if n < 16384 {
		w.Bits(uint64(n)|0x8000, 16)
		return n, false
	}
	m := n / 16384
	if m > 4 {
		m = 4
	}
	w.Bits(uint64(0xc0|m), 8)
	return m * 16384, true
}
func (w *Writer) OpenType(b []byte) {
	if len(b) == 0 {
		w.Fail(fmt.Errorf("empty APER open type"))
		return
	}
	for {
		n, more := w.Length(len(b))
		w.Octets(b[:n])
		b = b[n:]
		if w.err != nil || !more {
			return
		}
	}
}
func (w *Writer) NormallySmall(v uint64) {
	if v < 64 {
		w.Bits(v, 7)
		return
	}
	w.Bits(1, 1)
	n := (bits.Len64(v) + 7) / 8
	b := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		b[i] = byte(v)
		v >>= 8
	}
	w.OpenType(b)
}
