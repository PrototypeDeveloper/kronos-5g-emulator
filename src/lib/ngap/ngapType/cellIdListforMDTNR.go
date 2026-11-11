// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIdListforMDTNR represents ASN.1 CellIdListforMDT-NR.
type CellIdListforMDTNR struct {
	List []NRCGI `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIdListforMDTNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIdListforMDT-NR", v)
}

// DecodeCellIdListforMDTNR decodes a standalone APER value and reports the consumed octets.
func DecodeCellIdListforMDTNR(data []byte) (CellIdListforMDTNR, int, error) {
	var v CellIdListforMDTNR
	err := ngapcodec.DecodeParameter("CellIdListforMDT-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIdListforMDTNR{}, 0, err
	}
	return v, len(data), err
}
