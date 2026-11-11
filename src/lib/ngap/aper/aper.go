// Package aper implements aligned Packed Encoding Rules (ITU-T X.691).
package aper

import (
	"fmt"
	"math/big"
	"reflect"
)

type Codec struct {
	schema Schema
	limits Limits
}

func NewCodec(schema Schema, limits Limits) (*Codec, error) {
	if limits == (Limits{}) {
		limits = DefaultLimits()
	}
	if limits.MaxBytes <= 0 || limits.MaxElements <= 0 || limits.MaxDepth <= 0 || limits.MaxAllocation <= 0 {
		return nil, fmt.Errorf("invalid APER limits")
	}
	if len(schema.Types) == 0 {
		return nil, fmt.Errorf("empty APER schema")
	}
	return &Codec{schema, limits}, nil
}
func (c *Codec) descriptor(name string) (*Type, error) {
	if n, ok := c.schema.Names[name]; ok {
		name = n
	}
	t := c.schema.Types[name]
	if t == nil {
		return nil, fmt.Errorf("unknown ASN.1 type %s", name)
	}
	return t, nil
}

func (c *Codec) Marshal(name string, value any) ([]byte, error) {
	t, e := c.descriptor(name)
	if e != nil {
		return nil, e
	}
	v := reflect.ValueOf(value)
	for v.IsValid() && v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, fmt.Errorf("nil %s", name)
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return nil, fmt.Errorf("nil %s", name)
	}
	w := NewWriter(c.limits.MaxBytes)
	s := state{codec: c}
	if e = s.encode(w, t, v, 0); e != nil {
		return nil, fmt.Errorf("%s: %w", name, e)
	}
	return w.Bytes()
}

// Unmarshal replaces dst only after successful structural validation. A
// NotifyError accompanies a fully decoded value; all other errors leave dst unchanged.
func (c *Codec) Unmarshal(name string, data []byte, dst any) error {
	if len(data) > c.limits.MaxBytes {
		return fmt.Errorf("APER input exceeds limit")
	}
	t, e := c.descriptor(name)
	if e != nil {
		return e
	}
	v := reflect.ValueOf(dst)
	if !v.IsValid() || v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("APER destination must be a non-nil pointer")
	}
	s := state{codec: c}
	if e = s.allocate(v.Elem().Type().Size(), 1); e != nil {
		return e
	}
	tmp := reflect.New(v.Elem().Type()).Elem()
	r := NewReader(data)
	if e = s.decode(r, t, tmp, 0); e != nil {
		return fmt.Errorf("%s: %w", name, e)
	}
	if e = r.Finish(); e != nil {
		return fmt.Errorf("%s: %w", name, e)
	}
	v.Elem().Set(tmp)
	if len(s.diagnostics) > 0 {
		return &NotifyError{s.diagnostics}
	}
	return nil
}

type state struct {
	codec       *Codec
	allocated   uint64
	steps       int
	diagnostics []Diagnostic
}

func (s *state) guard(depth int) error {
	s.steps++
	if depth > s.codec.limits.MaxDepth || s.steps > s.codec.limits.MaxElements*64 {
		return fmt.Errorf("APER complexity limit exceeded")
	}
	return nil
}
func (s *state) allocate(size uintptr, n int) error {
	if n < 0 || n > s.codec.limits.MaxElements && size != 1 {
		return fmt.Errorf("APER element limit exceeded")
	}
	if size != 0 && uint64(n) > uint64(s.codec.limits.MaxAllocation)/uint64(size) {
		return fmt.Errorf("APER allocation limit exceeded")
	}
	s.allocated += uint64(size) * uint64(n)
	if s.allocated > uint64(s.codec.limits.MaxAllocation) {
		return fmt.Errorf("APER allocation limit exceeded")
	}
	return nil
}
func primitive(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Struct {
		if f := v.FieldByName("Value"); f.IsValid() {
			return f
		}
	}
	return v
}
func field(v reflect.Value, name string) (reflect.Value, error) {
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("expected struct for %s, got %s", name, v.Kind())
	}
	f := v.FieldByName(name)
	if !f.IsValid() {
		return f, fmt.Errorf("missing Go field %s", name)
	}
	return f, nil
}
func listValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Struct {
		if f := v.FieldByName("List"); f.IsValid() {
			return f
		}
	}
	return v
}
func constraint(t *Type, kind string) *Constraint {
	for i := range t.Constraints {
		if t.Constraints[i].Kind == kind {
			return &t.Constraints[i]
		}
	}
	return nil
}
func number(text string) (*big.Int, error) {
	v, ok := new(big.Int).SetString(text, 10)
	if !ok {
		return nil, fmt.Errorf("invalid ASN.1 bound %q", text)
	}
	return v, nil
}
func bounds(c *Constraint) (*big.Int, *big.Int, error) {
	lo, e := number(c.Lower)
	if e != nil {
		return nil, nil, e
	}
	hi, e := number(c.Upper)
	if e != nil {
		return nil, nil, e
	}
	if hi.Cmp(lo) < 0 {
		return nil, nil, fmt.Errorf("inverted ASN.1 range")
	}
	return lo, hi, nil
}
func inRanges(v *big.Int, rs []ValueRange) bool {
	for _, r := range rs {
		lo, e := number(r.Lower)
		if e != nil {
			return false
		}
		hi, e := number(r.Upper)
		if e != nil {
			return false
		}
		if v.Cmp(lo) >= 0 && v.Cmp(hi) <= 0 {
			return true
		}
	}
	return false
}
func integerValue(v reflect.Value) (*big.Int, error) {
	v = primitive(v)
	switch v.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return big.NewInt(v.Int()), nil
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return new(big.Int).SetUint64(v.Uint()), nil
	}
	return nil, fmt.Errorf("expected integer, got %s", v.Kind())
}
func setInteger(v reflect.Value, n *big.Int) error {
	v = primitive(v)
	switch v.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if !n.IsInt64() || v.OverflowInt(n.Int64()) {
			return fmt.Errorf("decoded integer overflows %s", v.Type())
		}
		v.SetInt(n.Int64())
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		if !n.IsUint64() || v.OverflowUint(n.Uint64()) {
			return fmt.Errorf("decoded integer overflows %s", v.Type())
		}
		v.SetUint(n.Uint64())
	default:
		return fmt.Errorf("expected integer, got %s", v.Kind())
	}
	return nil
}

func (s *state) encode(w *Writer, t *Type, v reflect.Value, depth int) error {
	if e := s.guard(depth); e != nil {
		return e
	}
	if w.Err() != nil {
		return w.Err()
	}
	if t == nil {
		return fmt.Errorf("missing ASN.1 type")
	}
	switch t.Kind {
	case "REFERENCE", "ALIAS":
		if t.Kind == "ALIAS" {
			v = primitive(v)
		}
		d, e := s.codec.descriptor(t.Ref)
		if e != nil {
			return e
		}
		return s.encode(w, d, v, depth+1)
	case "INTEGER":
		return encodeInteger(w, t, v)
	case "ENUMERATED":
		return encodeEnum(w, t, v)
	case "BOOLEAN":
		v = primitive(v)
		if v.Kind() != reflect.Bool {
			return fmt.Errorf("expected bool")
		}
		if v.Bool() {
			w.Bits(1, 1)
		} else {
			w.Bits(0, 1)
		}
	case "NULL":
	case "OCTET STRING", "BIT STRING", "PrintableString", "VisibleString", "UTF8String", "OBJECT IDENTIFIER":
		return s.encodeString(w, t, v)
	case "SEQUENCE":
		return s.encodeSequence(w, t, v, depth+1)
	case "CHOICE":
		return s.encodeChoice(w, t, v, depth+1)
	case "SEQUENCE OF", "CONTAINER":
		return s.encodeList(w, t, v, depth+1)
	case "ENTRY":
		return s.encodeEntry(w, t, v, depth+1)
	case "MESSAGE":
		return s.encodeMessage(w, t, v, depth+1)
	default:
		return fmt.Errorf("unsupported ASN.1 kind %s", t.Kind)
	}
	return w.Err()
}
func (s *state) decode(r *Reader, t *Type, v reflect.Value, depth int) error {
	if e := s.guard(depth); e != nil {
		return e
	}
	if r.Err() != nil {
		return r.Err()
	}
	if t == nil {
		return fmt.Errorf("missing ASN.1 type")
	}
	switch t.Kind {
	case "REFERENCE", "ALIAS":
		if t.Kind == "ALIAS" {
			v = primitive(v)
		}
		d, e := s.codec.descriptor(t.Ref)
		if e != nil {
			return e
		}
		return s.decode(r, d, v, depth+1)
	case "INTEGER":
		return decodeInteger(r, t, v)
	case "ENUMERATED":
		return decodeEnum(r, t, v)
	case "BOOLEAN":
		v = primitive(v)
		if v.Kind() != reflect.Bool {
			return fmt.Errorf("expected bool")
		}
		v.SetBool(r.Bits(1) != 0)
	case "NULL":
	case "OCTET STRING", "BIT STRING", "PrintableString", "VisibleString", "UTF8String", "OBJECT IDENTIFIER":
		return s.decodeString(r, t, v)
	case "SEQUENCE":
		return s.decodeSequence(r, t, v, depth+1)
	case "CHOICE":
		return s.decodeChoice(r, t, v, depth+1)
	case "SEQUENCE OF", "CONTAINER":
		return s.decodeList(r, t, v, depth+1)
	case "ENTRY":
		return s.decodeEntry(r, t, v, depth+1)
	case "MESSAGE":
		return s.decodeMessage(r, t, v, depth+1)
	default:
		return fmt.Errorf("unsupported ASN.1 kind %s", t.Kind)
	}
	return r.Err()
}
