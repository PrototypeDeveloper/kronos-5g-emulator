package aper

import (
	"fmt"
	"math/big"
	"reflect"
)

func criticality(s string) (uint64, error) {
	switch s {
	case "reject":
		return 0, nil
	case "ignore":
		return 1, nil
	case "notify":
		return 2, nil
	}
	return 0, fmt.Errorf("invalid criticality %s", s)
}
func (s *state) objectSet(name string) (ObjectSet, error) {
	set, ok := s.codec.schema.Sets[name]
	if !ok {
		return set, fmt.Errorf("unknown IE set %s", name)
	}
	return set, nil
}
func (s *state) unknownIE(id int64, crit uint64) error {
	switch crit {
	case 0:
		return &CriticalityError{id, fmt.Sprintf("unknown IE %d", id)}
	case 1:
		return nil
	case 2:
		s.diagnostics = append(s.diagnostics, Diagnostic{id, "notify", "unknown IE"})
		return nil
	}
	return fmt.Errorf("invalid criticality %d", crit)
}
func (s *state) encodeEntry(w *Writer, t *Type, v reflect.Value, depth int) error {
	v = primitive(v) // Named aliases may wrap a single-container entry in Value.
	set, e := s.objectSet(t.Ref)
	if e != nil {
		return e
	}
	pv, e := field(v, "Present")
	if e != nil {
		return e
	}
	i := int(pv.Uint()) - 1
	if i < 0 || i > len(set.IEs) {
		return fmt.Errorf("invalid IE selection")
	}
	var id, crit uint64
	var b []byte
	if i < len(set.IEs) {
		ie := set.IEs[i]
		id = uint64(ie.ID)
		crit, e = criticality(ie.Criticality)
		if e != nil {
			return e
		}
		fv, e := field(v, ie.GoField)
		if e != nil {
			return e
		}
		b, e = s.encodeOpen(ie.Type, fv, depth)
		if e != nil {
			return fmt.Errorf("IE %d: %w", id, e)
		}
	} else {
		cv, e := field(v, "Criticality")
		if e != nil {
			return e
		}
		n, e := integerValue(cv)
		if e != nil || !n.IsUint64() {
			return fmt.Errorf("invalid IE criticality")
		}
		crit = n.Uint64()
		if crit > 2 {
			return fmt.Errorf("invalid IE criticality")
		}
		raw, e := field(v, "UnknownValue")
		if e != nil {
			return e
		}
		b = raw.Bytes()
		if set.Class != "NGAP-PRIVATE-IES" {
			iv, e := field(v, "ID")
			if e != nil {
				return e
			}
			n, e := integerValue(iv)
			if e != nil || !n.IsUint64() {
				return fmt.Errorf("invalid IE ID")
			}
			id = n.Uint64()
			for _, ie := range set.IEs {
				if id == uint64(ie.ID) {
					return fmt.Errorf("raw IE collides with known ID %d", id)
				}
			}
		}
	}
	if set.Class == "NGAP-PRIVATE-IES" {
		iv, e := field(v, "ID")
		if e != nil {
			return e
		}
		d, e := s.codec.descriptor("PrivateIE-ID")
		if e != nil {
			return e
		}
		if e = s.encode(w, d, iv, depth); e != nil {
			return e
		}
	} else {
		w.Constrained(id, 0, 65535)
	}
	w.Constrained(crit, 0, 2)
	w.OpenType(b)
	return w.Err()
}
func (s *state) decodeEntry(r *Reader, t *Type, v reflect.Value, depth int) error {
	v = primitive(v)
	set, e := s.objectSet(t.Ref)
	if e != nil {
		return e
	}
	iv, e := field(v, "ID")
	if e != nil {
		return e
	}
	var id uint64
	if set.Class == "NGAP-PRIVATE-IES" {
		d, e := s.codec.descriptor("PrivateIE-ID")
		if e != nil {
			return e
		}
		if e = s.decode(r, d, iv, depth); e != nil {
			return e
		}
	} else {
		id = r.Constrained(0, 65535)
		if e = setInteger(iv, new(big.Int).SetUint64(id)); e != nil {
			return e
		}
	}
	crit := r.Constrained(0, 2)
	cv, e := field(v, "Criticality")
	if e != nil {
		return e
	}
	if e = setInteger(cv, new(big.Int).SetUint64(crit)); e != nil {
		return e
	}
	b := r.OpenType(s.codec.limits.MaxBytes)
	if r.Err() != nil {
		return r.Err()
	}
	if e = s.allocate(1, len(b)); e != nil {
		return e
	}
	p, e := field(v, "Present")
	if e != nil {
		return e
	}
	for i, ie := range set.IEs {
		if uint64(ie.ID) != id {
			continue
		}
		expected, e := criticality(ie.Criticality)
		if e != nil {
			return e
		}
		if crit != expected {
			return fmt.Errorf("IE %d criticality %d does not match %s", id, crit, ie.Criticality)
		}
		fv, e := field(v, ie.GoField)
		if e != nil {
			return e
		}
		if e = s.decodeOpen(b, ie.Type, fv, depth); e != nil {
			return fmt.Errorf("IE %d: %w", id, e)
		}
		p.SetUint(uint64(i + 1))
		return nil
	}
	if e = s.unknownIE(int64(id), crit); e != nil {
		return e
	}
	p.SetUint(uint64(len(set.IEs) + 1))
	raw, e := field(v, "UnknownValue")
	if e != nil {
		return e
	}
	raw.SetBytes(b)
	return nil
}

type wireIE struct {
	id, crit uint64
	value    []byte
}

func (s *state) encodeMessage(w *Writer, t *Type, v reflect.Value, depth int) error {
	set, e := s.objectSet(t.Ref)
	if e != nil {
		return e
	}
	ext, e := s.rawExtensions(v)
	if e != nil {
		return e
	}
	if t.Extensible {
		if len(ext) > 0 {
			w.Bits(1, 1)
		} else {
			w.Bits(0, 1)
		}
	} else if len(ext) > 0 {
		return fmt.Errorf("unexpected message extensions")
	}
	if set.Class == "NGAP-PRIVATE-IES" {
		f, e := field(v, "PrivateIEs")
		if e != nil {
			return e
		}
		ct := &Type{Kind: "CONTAINER", Ref: t.Ref, Constraints: []Constraint{{Kind: "size", Lower: "1", Upper: "65535"}}}
		if e = s.encodeList(w, ct, f, depth); e != nil {
			return e
		}
	} else {
		var ies []wireIE
		known := map[uint64]bool{}
		for _, ie := range set.IEs {
			id := uint64(ie.ID)
			known[id] = true
			if ie.Presence != "mandatory" {
				p, e := field(v, "Has"+ie.GoField)
				if e != nil {
					return e
				}
				if !p.Bool() {
					continue
				}
			}
			fv, e := field(v, ie.GoField)
			if e != nil {
				return e
			}
			b, e := s.encodeOpen(ie.Type, fv, depth)
			if e != nil {
				return fmt.Errorf("%s: %w", ie.GoField, e)
			}
			crit, e := criticality(ie.Criticality)
			if e != nil {
				return e
			}
			ies = append(ies, wireIE{id, crit, b})
		}
		raw, e := field(v, "AdditionalIEs")
		if e != nil {
			return e
		}
		seen := map[uint64]bool{}
		for i := 0; i < raw.Len(); i++ {
			x := raw.Index(i)
			id, e := integerValue(x.FieldByName("ID"))
			if e != nil {
				return e
			}
			cr, e := integerValue(x.FieldByName("Criticality"))
			if e != nil {
				return e
			}
			if !id.IsUint64() || id.Uint64() > 65535 || !cr.IsUint64() || cr.Uint64() > 2 {
				return fmt.Errorf("invalid raw IE header")
			}
			n := id.Uint64()
			if known[n] || seen[n] {
				return fmt.Errorf("duplicate or known AdditionalIE ID %d", n)
			}
			seen[n] = true
			ies = append(ies, wireIE{n, cr.Uint64(), x.FieldByName("Value").Bytes()})
		}
		if len(ies) > s.codec.limits.MaxElements {
			return fmt.Errorf("too many IEs")
		}
		w.Constrained(uint64(len(ies)), 0, 65535)
		for _, ie := range ies {
			w.Constrained(ie.id, 0, 65535)
			w.Constrained(ie.crit, 0, 2)
			w.OpenType(ie.value)
		}
	}
	if len(ext) > 0 {
		return s.writeExtensions(w, ext)
	}
	return w.Err()
}
func (s *state) decodeMessage(r *Reader, t *Type, v reflect.Value, depth int) error {
	set, e := s.objectSet(t.Ref)
	if e != nil {
		return e
	}
	ext := t.Extensible && r.Bits(1) != 0
	if set.Class == "NGAP-PRIVATE-IES" {
		f, e := field(v, "PrivateIEs")
		if e != nil {
			return e
		}
		ct := &Type{Kind: "CONTAINER", Ref: t.Ref, Constraints: []Constraint{{Kind: "size", Lower: "1", Upper: "65535"}}}
		if e = s.decodeList(r, ct, f, depth); e != nil {
			return e
		}
	} else {
		n := r.Constrained(0, 65535)
		if n > uint64(s.codec.limits.MaxElements) {
			return fmt.Errorf("too many IEs")
		}
		known := map[uint64]IE{}
		for _, ie := range set.IEs {
			known[uint64(ie.ID)] = ie
		}
		seen := map[uint64]bool{}
		for i := uint64(0); i < n; i++ {
			id := r.Constrained(0, 65535)
			crit := r.Constrained(0, 2)
			b := r.OpenType(s.codec.limits.MaxBytes)
			if r.Err() != nil {
				return r.Err()
			}
			if e = s.allocate(1, len(b)); e != nil {
				return e
			}
			if seen[id] {
				return fmt.Errorf("duplicate IE %d", id)
			}
			seen[id] = true
			if ie, ok := known[id]; ok {
				expected, e := criticality(ie.Criticality)
				if e != nil {
					return e
				}
				if crit != expected {
					return fmt.Errorf("IE %d criticality %d does not match %s", id, crit, ie.Criticality)
				}
				fv, e := field(v, ie.GoField)
				if e != nil {
					return e
				}
				if e = s.decodeOpen(b, ie.Type, fv, depth); e != nil {
					return fmt.Errorf("%s (IE %d): %w", ie.GoField, id, e)
				}
				if ie.Presence != "mandatory" {
					if e = setPresence(v, ie.GoField); e != nil {
						return e
					}
				}
			} else {
				if e = s.unknownIE(int64(id), crit); e != nil {
					return e
				}
				raw, e := field(v, "AdditionalIEs")
				if e != nil {
					return e
				}
				if e = s.allocate(raw.Type().Elem().Size(), 1); e != nil {
					return e
				}
				x := reflect.New(raw.Type().Elem()).Elem()
				if e = setInteger(x.FieldByName("ID"), new(big.Int).SetUint64(id)); e != nil {
					return e
				}
				if e = setInteger(x.FieldByName("Criticality"), new(big.Int).SetUint64(crit)); e != nil {
					return e
				}
				x.FieldByName("Value").SetBytes(b)
				raw.Set(reflect.Append(raw, x))
			}
		}
		for _, ie := range set.IEs {
			if ie.Presence == "mandatory" && !seen[uint64(ie.ID)] {
				return fmt.Errorf("mandatory IE %s (%d) missing", ie.GoField, ie.ID)
			}
		}
	}
	if ext {
		bitmap, e := s.extensionBitsRead(r)
		if e != nil {
			return e
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
			if e = s.appendExtension(v, i, b); e != nil {
				return e
			}
		}
	}
	return r.Err()
}
