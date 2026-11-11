// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellCAGInformation represents ASN.1 Cell-CAGInformation.
type CellCAGInformation struct {
	NGRANCGI          NGRANCGI
	CellCAGList       CellCAGList
	HasIEExtensions   bool
	IEExtensions      []CellCAGInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellCAGInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Cell-CAGInformation", v)
}

// DecodeCellCAGInformation decodes a standalone APER value and reports the consumed octets.
func DecodeCellCAGInformation(data []byte) (CellCAGInformation, int, error) {
	var v CellCAGInformation
	err := ngapcodec.DecodeParameter("Cell-CAGInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellCAGInformation{}, 0, err
	}
	return v, len(data), err
}
