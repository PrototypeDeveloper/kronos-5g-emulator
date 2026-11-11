// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ProcedureCode represents ASN.1 ProcedureCode.
type ProcedureCode struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ProcedureCode) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ProcedureCode", v) }

// DecodeProcedureCode decodes a standalone APER value and reports the consumed octets.
func DecodeProcedureCode(data []byte) (ProcedureCode, int, error) {
	var v ProcedureCode
	err := ngapcodec.DecodeParameter("ProcedureCode", data, &v)
	if ngapcodec.IsFatal(err) {
		return ProcedureCode{}, 0, err
	}
	return v, len(data), err
}
