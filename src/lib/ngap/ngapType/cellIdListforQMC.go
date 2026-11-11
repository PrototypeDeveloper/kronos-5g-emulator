// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIdListforQMC represents ASN.1 CellIdListforQMC.
type CellIdListforQMC struct {
	List []NGRANCGI `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIdListforQMC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIdListforQMC", v)
}

// DecodeCellIdListforQMC decodes a standalone APER value and reports the consumed octets.
func DecodeCellIdListforQMC(data []byte) (CellIdListforQMC, int, error) {
	var v CellIdListforQMC
	err := ngapcodec.DecodeParameter("CellIdListforQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIdListforQMC{}, 0, err
	}
	return v, len(data), err
}
