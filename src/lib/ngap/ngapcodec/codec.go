// Package ngapcodec connects the NGAP ASN.1 schema to the generic APER codec.
package ngapcodec

import (
	"errors"
	"fmt"
	"libngap/aper"
)

var engine, schemaError = aper.NewCodec(NewSchema(), aper.DefaultLimits())

type Procedure struct{ Kind, Code, Criticality uint64 }
type Header struct {
	Kind, ProcedureCode, Criticality uint64
	Value                            []byte
}

func DecodeHeader(data []byte) (Header, error) {
	var h Header
	if len(data) > aper.DefaultLimits().MaxBytes {
		return h, fmt.Errorf("NGAP PDU exceeds limit")
	}
	r := aper.NewReader(data)
	if r.Bits(1) != 0 {
		return h, fmt.Errorf("unsupported NGAP-PDU extension")
	}
	h.Kind = r.Constrained(0, 2)
	h.ProcedureCode = r.Constrained(0, 255)
	h.Criticality = r.Constrained(0, 2)
	h.Value = r.OpenType(aper.DefaultLimits().MaxBytes)
	if e := r.Finish(); e != nil {
		return Header{}, e
	}
	return h, nil
}
func Encode(name string, v any) ([]byte, error) {
	if schemaError != nil {
		return nil, schemaError
	}
	p, ok := procedures[name]
	if !ok {
		return nil, fmt.Errorf("unknown NGAP message %s", name)
	}
	b, e := engine.Marshal(name, v)
	if e != nil {
		return nil, e
	}
	w := aper.NewWriter(aper.DefaultLimits().MaxBytes)
	w.Bits(0, 1)
	w.Constrained(p.Kind, 0, 2)
	w.Constrained(p.Code, 0, 255)
	w.Constrained(p.Criticality, 0, 2)
	w.OpenType(b)
	return w.Bytes()
}
func Decode(name string, data []byte, v any) error {
	if schemaError != nil {
		return schemaError
	}
	p, ok := procedures[name]
	if !ok {
		return fmt.Errorf("unknown NGAP message %s", name)
	}
	h, e := DecodeHeader(data)
	if e != nil {
		return e
	}
	if h.Kind != p.Kind || h.ProcedureCode != p.Code {
		return fmt.Errorf("PDU does not contain %s", name)
	}
	if h.Criticality != p.Criticality {
		return fmt.Errorf("%s has unexpected procedure criticality %d", name, h.Criticality)
	}
	return engine.Unmarshal(name, h.Value, v)
}
func EncodeParameter(name string, v any) ([]byte, error) {
	if schemaError != nil {
		return nil, schemaError
	}
	return engine.Marshal(name, v)
}
func DecodeParameter(name string, data []byte, v any) error {
	if schemaError != nil {
		return schemaError
	}
	return engine.Unmarshal(name, data, v)
}
func IsFatal(err error) bool {
	var notify *aper.NotifyError
	return err != nil && !errors.As(err, &notify)
}
