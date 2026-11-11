// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellBasedQMC represents ASN.1 CellBasedQMC.
type CellBasedQMC struct {
	CellIdListforQMC  CellIdListforQMC
	HasIEExtensions   bool
	IEExtensions      []CellBasedQMCExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellBasedQMC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CellBasedQMC", v) }

// DecodeCellBasedQMC decodes a standalone APER value and reports the consumed octets.
func DecodeCellBasedQMC(data []byte) (CellBasedQMC, int, error) {
	var v CellBasedQMC
	err := ngapcodec.DecodeParameter("CellBasedQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellBasedQMC{}, 0, err
	}
	return v, len(data), err
}
