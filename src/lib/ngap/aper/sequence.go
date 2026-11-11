package aper

import (
	"fmt"
	"math/big"
	"reflect"
)

func present(v reflect.Value, f Field) (bool, error) {
	if !f.Optional && f.Default == "" && !f.Extension {
		return true, nil
	}
	p, e := field(v, "Has"+f.GoField)
	if e != nil {
		return false, e
	}
	if p.Kind() != reflect.Bool {
		return false, fmt.Errorf("presence field must be bool")
	}
	return p.Bool(), nil
}
func setPresence(v reflect.Value, name string) error {
	p, e := field(v, "Has"+name)
	if e != nil {
		return e
	}
	if p.Kind() != reflect.Bool {
		return fmt.Errorf("presence field must be bool")
	}
	p.SetBool(true)
	return nil
}
func (s *state) defaultValue(t *Type, text string, v reflect.Value) error {
	if t.Kind == "REFERENCE" {
		d, e := s.codec.descriptor(t.Ref)
		if e != nil {
			return e
		}
		return s.defaultValue(d, text, v)
	}
	if t.Kind == "BOOLEAN" {
		primitive(v).SetBool(text == "TRUE")
		return nil
	}
	if t.Kind == "ENUMERATED" {
		for _, x := range t.Values {
			if x.Name == text {
				return setInteger(v, big.NewInt(x.Number))
			}
		}
	}
	n, e := number(text)
	if e != nil {
		return e
	}
	return setInteger(v, n)
}
func (s *state) encodeOpen(t *Type, v reflect.Value, depth int) ([]byte, error) {
	w := NewWriter(s.codec.limits.MaxBytes)
	if e := s.encode(w, t, v, depth+1); e != nil {
		return nil, e
	}
	return w.Bytes()
}
func (s *state) decodeOpen(b []byte, t *Type, v reflect.Value, depth int) error {
	r := NewReader(b)
	if e := s.decode(r, t, v, depth+1); e != nil {
		return e
	}
	return r.Finish()
}
func (s *state) rawExtensions(v reflect.Value) (map[int][]byte, error) {
	out := map[int][]byte{}
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("expected sequence struct")
	}
	f := v.FieldByName("UnknownExtensions")
	if !f.IsValid() {
		return out, nil
	}
	if f.Kind() != reflect.Slice {
		return nil, fmt.Errorf("invalid UnknownExtensions")
	}
	for i := 0; i < f.Len(); i++ {
		x := f.Index(i)
		index := x.FieldByName("Index").Uint()
		if index >= uint64(s.codec.limits.MaxElements) {
			return nil, fmt.Errorf("extension index exceeds limit")
		}
		if _, ok := out[int(index)]; ok {
			return nil, fmt.Errorf("duplicate extension index %d", index)
		}
		out[int(index)] = x.FieldByName("Value").Bytes()
	}
	return out, nil
}
func (s *state) appendExtension(v reflect.Value, index int, b []byte) error {
	f, e := field(v, "UnknownExtensions")
	if e != nil {
		return e
	}
	if e = s.allocate(f.Type().Elem().Size(), 1); e != nil {
		return e
	}
	x := reflect.New(f.Type().Elem()).Elem()
	x.FieldByName("Index").SetUint(uint64(index))
	x.FieldByName("Value").SetBytes(b)
	f.Set(reflect.Append(f, x))
	return nil
}
func (s *state) extensionBitsWrite(w *Writer, n int, payloads map[int][]byte) error {
	if n < 1 || n > s.codec.limits.MaxElements {
		return fmt.Errorf("extension bitmap exceeds limit")
	}
	if n <= 64 {
		w.Bits(uint64(n-1), 7)
		for i := 0; i < n; i++ {
			_, ok := payloads[i]
			if ok {
				w.Bits(1, 1)
			} else {
				w.Bits(0, 1)
			}
		}
	} else {
		w.Bits(1, 1)
		if e := s.writeSized(w, nil, n, func(off, count int, _ *Constraint) error {
			for i := off; i < off+count; i++ {
				_, ok := payloads[i]
				if ok {
					w.Bits(1, 1)
				} else {
					w.Bits(0, 1)
				}
			}
			return w.Err()
		}); e != nil {
			return e
		}
	}
	return w.Err()
}
func (s *state) extensionBitsRead(r *Reader) ([]bool, error) {
	var bitmap []bool
	read := func(n int, _ *Constraint) error {
		if n > s.codec.limits.MaxElements-len(bitmap) {
			return fmt.Errorf("extension bitmap exceeds limit")
		}
		if e := s.allocate(1, n); e != nil {
			return e
		}
		for i := 0; i < n; i++ {
			bitmap = append(bitmap, r.Bits(1) != 0)
		}
		return r.Err()
	}
	if r.Bits(1) == 0 {
		if e := read(int(r.Bits(6))+1, nil); e != nil {
			return nil, e
		}
	} else {
		if e := s.readSized(r, nil, read); e != nil {
			return nil, e
		}
	}
	if len(bitmap) == 0 {
		return nil, fmt.Errorf("empty extension bitmap")
	}
	return bitmap, r.Err()
}
func (s *state) writeExtensions(w *Writer, payloads map[int][]byte) error {
	n := 0
	for i := range payloads {
		if i >= n {
			n = i + 1
		}
	}
	if e := s.extensionBitsWrite(w, n, payloads); e != nil {
		return e
	}
	for i := 0; i < n; i++ {
		if b, ok := payloads[i]; ok {
			w.OpenType(b)
		}
	}
	return w.Err()
}
func (s *state) encodeSequence(w *Writer, t *Type, v reflect.Value, depth int) error {
	payloads, e := s.rawExtensions(v)
	if e != nil {
		return e
	}
	extIndex := 0
	for _, f := range t.Fields {
		if !f.Extension {
			continue
		}
		has, e := present(v, f)
		if e != nil {
			return e
		}
		if _, ok := payloads[extIndex]; ok {
			return fmt.Errorf("raw extension conflicts with known field %s", f.GoField)
		}
		if has {
			fv, e := field(v, f.GoField)
			if e != nil {
				return e
			}
			b, e := s.encodeOpen(f.Type, fv, depth)
			if e != nil {
				return e
			}
			payloads[extIndex] = b
		}
		extIndex++
	}
	if t.Extensible {
		if len(payloads) > 0 {
			w.Bits(1, 1)
		} else {
			w.Bits(0, 1)
		}
	} else if len(payloads) > 0 {
		return fmt.Errorf("extensions on non-extensible sequence")
	}
	presence := map[string]bool{}
	for _, f := range t.Fields {
		if f.Extension {
			continue
		}
		has, e := present(v, f)
		if e != nil {
			return e
		}
		if has && f.Default != "" {
			fv, e := field(v, f.GoField)
			if e != nil {
				return e
			}
			def := reflect.New(fv.Type()).Elem()
			if e = s.defaultValue(f.Type, f.Default, def); e != nil {
				return e
			}
			if reflect.DeepEqual(def.Interface(), fv.Interface()) {
				has = false
			}
		}
		presence[f.GoField] = has
		if f.Optional || f.Default != "" {
			if has {
				w.Bits(1, 1)
			} else {
				w.Bits(0, 1)
			}
		}
	}
	for _, f := range t.Fields {
		if f.Extension || !presence[f.GoField] {
			continue
		}
		fv, e := field(v, f.GoField)
		if e != nil {
			return e
		}
		if e = s.encode(w, f.Type, fv, depth); e != nil {
			return fmt.Errorf("%s: %w", f.GoField, e)
		}
	}
	if len(payloads) > 0 {
		return s.writeExtensions(w, payloads)
	}
	return w.Err()
}
func (s *state) decodeSequence(r *Reader, t *Type, v reflect.Value, depth int) error {
	ext := t.Extensible && r.Bits(1) != 0
	presence := map[string]bool{}
	for _, f := range t.Fields {
		if f.Extension {
			continue
		}
		has := true
		if f.Optional || f.Default != "" {
			has = r.Bits(1) != 0
			if has {
				if e := setPresence(v, f.GoField); e != nil {
					return e
				}
			}
		}
		presence[f.GoField] = has
	}
	for _, f := range t.Fields {
		if f.Extension {
			continue
		}
		fv, e := field(v, f.GoField)
		if e != nil {
			return e
		}
		if presence[f.GoField] {
			if e = s.decode(r, f.Type, fv, depth); e != nil {
				return fmt.Errorf("%s: %w", f.GoField, e)
			}
		} else if f.Default != "" {
			if e = s.defaultValue(f.Type, f.Default, fv); e != nil {
				return e
			}
		}
	}
	if ext {
		bitmap, e := s.extensionBitsRead(r)
		if e != nil {
			return e
		}
		var known []Field
		for _, f := range t.Fields {
			if f.Extension {
				known = append(known, f)
			}
		}
		for i, has := range bitmap {
			if !has {
				continue
			}
			b := r.OpenType(s.codec.limits.MaxBytes)
			if r.Err() != nil {
				return r.Err()
			}
			if e = s.allocate(1, len(b)); e != nil {
				return e
			}
			if i < len(known) {
				f := known[i]
				fv, e := field(v, f.GoField)
				if e != nil {
					return e
				}
				if e = setPresence(v, f.GoField); e != nil {
					return e
				}
				if e = s.decodeOpen(b, f.Type, fv, depth); e != nil {
					return e
				}
			} else {
				if e = s.appendExtension(v, i, b); e != nil {
					return e
				}
			}
		}
	}
	return r.Err()
}
func (s *state) encodeChoice(w *Writer, t *Type, v reflect.Value, depth int) error {
	p, e := field(v, "Present")
	if e != nil {
		return e
	}
	selected := int(p.Uint()) - 1
	if selected < 0 {
		return fmt.Errorf("CHOICE is not selected")
	}
	if selected >= len(t.Fields) {
		if !t.Extensible || selected != len(t.Fields) {
			return fmt.Errorf("invalid CHOICE selector")
		}
		u, e := field(v, "UnknownExtension")
		if e != nil {
			return e
		}
		idx := u.FieldByName("Index").Uint()
		if idx < uint64(len(t.Fields)-t.RootCount) {
			return fmt.Errorf("unknown CHOICE extension collides with known alternative")
		}
		w.Bits(1, 1)
		w.NormallySmall(idx)
		w.OpenType(u.FieldByName("Value").Bytes())
		return w.Err()
	}
	f := t.Fields[selected]
	fv, e := field(v, f.GoField)
	if e != nil {
		return e
	}
	if t.Extensible {
		if f.Extension {
			w.Bits(1, 1)
		} else {
			w.Bits(0, 1)
		}
	}
	if f.Extension {
		w.NormallySmall(uint64(selected - t.RootCount))
		b, e := s.encodeOpen(f.Type, fv, depth)
		if e != nil {
			return e
		}
		w.OpenType(b)
		return w.Err()
	}
	w.Constrained(uint64(selected), 0, uint64(t.RootCount-1))
	return s.encode(w, f.Type, fv, depth)
}
func (s *state) decodeChoice(r *Reader, t *Type, v reflect.Value, depth int) error {
	p, e := field(v, "Present")
	if e != nil {
		return e
	}
	ext := t.Extensible && r.Bits(1) != 0
	if ext {
		idx := r.NormallySmall()
		b := r.OpenType(s.codec.limits.MaxBytes)
		if r.Err() != nil {
			return r.Err()
		}
		if e = s.allocate(1, len(b)); e != nil {
			return e
		}
		if idx >= uint64(len(t.Fields)-t.RootCount) {
			u, e := field(v, "UnknownExtension")
			if e != nil {
				return e
			}
			u.FieldByName("Index").SetUint(idx)
			u.FieldByName("Value").SetBytes(b)
			p.SetUint(uint64(len(t.Fields) + 1))
			return nil
		}
		f := t.Fields[t.RootCount+int(idx)]
		p.SetUint(uint64(t.RootCount) + idx + 1)
		fv, e := field(v, f.GoField)
		if e != nil {
			return e
		}
		return s.decodeOpen(b, f.Type, fv, depth)
	}
	if t.RootCount == 0 {
		return fmt.Errorf("empty CHOICE root")
	}
	idx := r.Constrained(0, uint64(t.RootCount-1))
	if r.Err() != nil {
		return r.Err()
	}
	p.SetUint(idx + 1)
	f := t.Fields[idx]
	fv, e := field(v, f.GoField)
	if e != nil {
		return e
	}
	return s.decode(r, f.Type, fv, depth)
}
func (s *state) encodeList(w *Writer, t *Type, v reflect.Value, depth int) error {
	v = listValue(v)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("expected list")
	}
	if v.Len() > s.codec.limits.MaxElements {
		return fmt.Errorf("list exceeds element limit")
	}
	element := t.Element
	if t.Kind == "CONTAINER" {
		element = &Type{Kind: "ENTRY", Ref: t.Ref}
		if e := s.validateContainer(t, v); e != nil {
			return e
		}
	}
	return s.writeSized(w, constraint(t, "size"), v.Len(), func(off, n int, _ *Constraint) error {
		for i := off; i < off+n; i++ {
			if e := s.encode(w, element, v.Index(i), depth); e != nil {
				return fmt.Errorf("[%d]: %w", i, e)
			}
		}
		return w.Err()
	})
}
func (s *state) decodeList(r *Reader, t *Type, v reflect.Value, depth int) error {
	v = listValue(v)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("expected list")
	}
	element := t.Element
	if t.Kind == "CONTAINER" {
		element = &Type{Kind: "ENTRY", Ref: t.Ref}
	}
	err := s.readSized(r, constraint(t, "size"), func(n int, _ *Constraint) error {
		if n > s.codec.limits.MaxElements-v.Len() {
			return fmt.Errorf("list exceeds element limit")
		}
		if e := s.allocate(v.Type().Elem().Size(), n); e != nil {
			return e
		}
		chunk := reflect.MakeSlice(v.Type(), n, n)
		for i := 0; i < n; i++ {
			if e := s.decode(r, element, chunk.Index(i), depth); e != nil {
				return fmt.Errorf("[%d]: %w", i, e)
			}
		}
		v.Set(reflect.AppendSlice(v, chunk))
		return r.Err()
	})
	if err != nil {
		return err
	}
	if t.Kind == "CONTAINER" {
		return s.validateContainer(t, v)
	}
	return nil
}

func (s *state) validateContainer(t *Type, v reflect.Value) error {
	set, e := s.objectSet(t.Ref)
	if e != nil {
		return e
	}
	if set.Class == "NGAP-PRIVATE-IES" {
		return nil
	}
	seen := map[uint64]bool{}
	for i := 0; i < v.Len(); i++ {
		x := v.Index(i)
		p, e := field(x, "Present")
		if e != nil {
			return e
		}
		index := p.Uint()
		var id uint64
		if index >= 1 && index <= uint64(len(set.IEs)) {
			id = uint64(set.IEs[index-1].ID)
		} else if index == uint64(len(set.IEs)+1) {
			n, e := integerValue(x.FieldByName("ID"))
			if e != nil || !n.IsUint64() {
				return fmt.Errorf("invalid IE ID")
			}
			id = n.Uint64()
		} else {
			return fmt.Errorf("invalid IE selection")
		}
		if seen[id] {
			return fmt.Errorf("duplicate IE %d", id)
		}
		seen[id] = true
	}
	for _, ie := range set.IEs {
		if ie.Presence == "mandatory" && !seen[uint64(ie.ID)] {
			return fmt.Errorf("mandatory IE %s (%d) missing", ie.GoField, ie.ID)
		}
	}
	return nil
}
