// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellBasedMDTNR represents ASN.1 CellBasedMDT-NR.
type CellBasedMDTNR struct {
	CellIdListforMDT  CellIdListforMDTNR
	HasIEExtensions   bool
	IEExtensions      []CellBasedMDTNRExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedMDTNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellBasedMDT-NR", v)
}

// DecodeCellBasedMDTNR decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedMDTNR(data []byte) (CellBasedMDTNR, int, error) {
	var v CellBasedMDTNR
	err := ngapcodec.DecodeParameter("CellBasedMDT-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedMDTNR{}, 0, err
	}
	return v, len(data), err
}
