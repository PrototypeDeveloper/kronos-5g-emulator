package nascodec

import (
	"encoding/binary"
	"fmt"
	"io"
	"libnas/nasType"
)

type Format uint8

const (
	TV Format = iota
	HalfTV
	TLV
	TLVE
)

type Rule struct {
	Format   Format
	Min, Max int
}
type Schema map[uint8]Rule

type Writer struct {
	Data []byte
	Err  error
}

func (w *Writer) fail(err error) {
	if w.Err == nil {
		w.Err = err
	}
}
func (w *Writer) U8(v uint8, max uint8) {
	if v > max {
		w.fail(fmt.Errorf("value %d exceeds %d", v, max))
	}
	w.Data = append(w.Data, v)
}
func (w *Writer) Fixed(v []byte, n int) {
	if len(v) != n {
		w.fail(fmt.Errorf("expected %d bytes, got %d", n, len(v)))
		return
	}
	w.Data = append(w.Data, v...)
}
func (w *Writer) LV(v []byte, width, min, max int) {
	if len(v) < min || len(v) > max {
		w.fail(fmt.Errorf("IE length %d outside %d..%d", len(v), min, max))
		return
	}
	if width == 1 {
		if len(v) > 255 {
			w.fail(fmt.Errorf("LV length overflow"))
			return
		}
		w.Data = append(w.Data, byte(len(v)))
	} else {
		if len(v) > 65535 {
			w.fail(fmt.Errorf("LV-E length overflow"))
			return
		}
		w.Data = append(w.Data, byte(len(v)>>8), byte(len(v)))
	}
	w.Data = append(w.Data, v...)
}
func (w *Writer) Options(options []nasType.IE, schema Schema) {
	seen := map[uint8]bool{}
	for _, ie := range options {
		rule, ok := schema[ie.ID]
		if !ok {
			w.fail(fmt.Errorf("unsupported optional IE 0x%02x", ie.ID))
			return
		}
		if seen[ie.ID] {
			w.fail(fmt.Errorf("duplicate IE 0x%02x", ie.ID))
			return
		}
		seen[ie.ID] = true
		if len(ie.Value) < rule.Min || len(ie.Value) > rule.Max {
			w.fail(fmt.Errorf("invalid length for IE 0x%02x", ie.ID))
			return
		}
		if rule.Format == HalfTV {
			if ie.Value[0] > 15 {
				w.fail(fmt.Errorf("half-octet IE overflow"))
				return
			}
			w.Data = append(w.Data, ie.ID|ie.Value[0])
			continue
		}
		w.Data = append(w.Data, ie.ID)
		switch rule.Format {
		case TV:
			w.Fixed(ie.Value, rule.Min)
		case TLV:
			w.LV(ie.Value, 1, rule.Min, rule.Max)
		case TLVE:
			w.LV(ie.Value, 2, rule.Min, rule.Max)
		}
	}
}

// UnknownOptions writes back IEs that schema does not model (as decoded by
// Reader.OptionsWithUnknown), in the format unknownFormat gives their IEI.
func (w *Writer) UnknownOptions(options []nasType.IE, schema Schema) {
	for _, ie := range options {
		if _, ok := schema[ie.ID]; ok {
			w.fail(fmt.Errorf("IE 0x%02x is a modelled IE, not an unknown one", ie.ID))
			return
		}
		format, err := unknownFormat(ie.ID)
		if err != nil {
			w.fail(err)
			return
		}
		switch format {
		case HalfTV:
			if ie.ID&15 != 0 || len(ie.Value) != 1 || ie.Value[0] > 15 {
				w.fail(fmt.Errorf("invalid half-octet IE 0x%02x", ie.ID))
				return
			}
			w.Data = append(w.Data, ie.ID|ie.Value[0])
		case TLV:
			w.Data = append(w.Data, ie.ID)
			w.LV(ie.Value, 1, 0, 255)
		case TLVE:
			w.Data = append(w.Data, ie.ID)
			w.LV(ie.Value, 2, 0, 65535)
		}
	}
}
func (w *Writer) Result() ([]byte, error) {
	if w.Err != nil {
		return nil, w.Err
	}
	if len(w.Data) > nasType.MaxMessageSize {
		return nil, fmt.Errorf("NAS message too large")
	}
	return w.Data, nil
}

type Reader struct {
	Data []byte
	Err  error
}

func (r *Reader) Take(n int) []byte {
	if r.Err != nil {
		return nil
	}
	if n < 0 || len(r.Data) < n {
		r.Err = io.ErrUnexpectedEOF
		return nil
	}
	v := append([]byte(nil), r.Data[:n]...)
	r.Data = r.Data[n:]
	return v
}
func (r *Reader) U8(max uint8) uint8 {
	b := r.Take(1)
	if len(b) == 0 {
		return 0
	}
	if b[0] > max {
		r.Err = fmt.Errorf("value %d exceeds %d", b[0], max)
	}
	return b[0]
}
func (r *Reader) LV(width, min, max int) []byte {
	b := r.Take(width)
	if r.Err != nil {
		return nil
	}
	n := int(b[0])
	if width == 2 {
		n = int(binary.BigEndian.Uint16(b))
	}
	if n < min || n > max {
		r.Err = fmt.Errorf("IE length %d outside %d..%d", n, min, max)
		return nil
	}
	return r.Take(n)
}
func (r *Reader) Options(schema Schema) []nasType.IE {
	known, _ := r.options(schema, false)
	return known
}

// OptionsWithUnknown reads optional IEs like Options and also returns the
// IEs the schema does not model, which Options skips.
func (r *Reader) OptionsWithUnknown(schema Schema) (known, unknown []nasType.IE) {
	return r.options(schema, true)
}

func (r *Reader) options(schema Schema, keepUnknown bool) (out, unknown []nasType.IE) {
	seen := map[uint8]bool{}
	for len(r.Data) > 0 && r.Err == nil {
		first := r.U8(255)
		id := first
		if id >= 0x80 {
			id &= 0xf0
		}
		rule, ok := schema[id]
		if !ok {
			format, err := unknownFormat(first)
			if err != nil {
				r.Err = err
				break
			}
			switch {
			case format == HalfTV:
				if keepUnknown {
					unknown = append(unknown, nasType.IE{ID: id, Value: []byte{first & 15}})
				}
			case keepUnknown:
				width := 1
				if format == TLVE {
					width = 2
				}
				value := r.LV(width, 0, 65535)
				if r.Err == nil {
					unknown = append(unknown, nasType.IE{ID: first, Value: value})
				}
			case format == TLVE:
				r.skipLV(2)
			default:
				r.skipLV(1)
			}
			continue
		}
		// TS 24.501 7.6.3: keep the first occurrence. Even an ignored
		// duplicate must have a complete envelope, so it cannot hide a
		// truncated IE or consume bytes belonging to the next message.
		if seen[id] {
			switch rule.Format {
			case TV:
				r.Take(rule.Min)
			case TLV:
				r.skipLV(1)
			case TLVE:
				r.skipLV(2)
			}
			continue
		}
		seen[id] = true
		var value []byte
		switch rule.Format {
		case HalfTV:
			value = []byte{first & 15}
		case TV:
			value = r.Take(rule.Min)
		case TLV:
			value = r.LV(1, rule.Min, rule.Max)
		case TLVE:
			value = r.LV(2, rule.Min, rule.Max)
		}
		out = append(out, nasType.IE{ID: id, Value: value})
	}
	return out, unknown
}

// unknownFormat gives the format of an IE the schema does not know.
// TS 24.007 11.2.4/11.2.5 (5GS): unknown type 1/2 is one octet;
// 0x70..0x7f is TLV-E; other full IEIs (0x6X too) are TLV.
// 0x00..0x0f and 0x7e/0x7f require comprehension.
func unknownFormat(iei uint8) (Format, error) {
	switch {
	case iei < 0x10 || iei == 0x7e || iei == 0x7f:
		return 0, fmt.Errorf("unknown comprehension-required IE 0x%02x", iei)
	case iei >= 0x80:
		return HalfTV, nil
	case iei >= 0x70:
		return TLVE, nil
	}
	return TLV, nil
}

func (r *Reader) skipLV(width int) {
	b := r.Take(width)
	if r.Err != nil {
		return
	}
	n := int(b[0])
	if width == 2 {
		n = int(binary.BigEndian.Uint16(b))
	}
	if n > len(r.Data) {
		r.Err = io.ErrUnexpectedEOF
		return
	}
	r.Data = r.Data[n:]
}
