// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIdListforMDTEUTRA represents ASN.1 CellIdListforMDT-EUTRA.
type CellIdListforMDTEUTRA struct {
	List []EUTRACGI `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIdListforMDTEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIdListforMDT-EUTRA", v)
}

// DecodeCellIdListforMDTEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCellIdListforMDTEUTRA(data []byte) (CellIdListforMDTEUTRA, int, error) {
	var v CellIdListforMDTEUTRA
	err := ngapcodec.DecodeParameter("CellIdListforMDT-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIdListforMDTEUTRA{}, 0, err
	}
	return v, len(data), err
}
