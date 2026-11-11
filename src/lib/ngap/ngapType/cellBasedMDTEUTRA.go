// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellBasedMDTEUTRA represents ASN.1 CellBasedMDT-EUTRA.
type CellBasedMDTEUTRA struct {
	CellIdListforMDT  CellIdListforMDTEUTRA
	HasIEExtensions   bool
	IEExtensions      []CellBasedMDTEUTRAExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedMDTEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellBasedMDT-EUTRA", v)
}

// DecodeCellBasedMDTEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedMDTEUTRA(data []byte) (CellBasedMDTEUTRA, int, error) {
	var v CellBasedMDTEUTRA
	err := ngapcodec.DecodeParameter("CellBasedMDT-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedMDTEUTRA{}, 0, err
	}
	return v, len(data), err
}
