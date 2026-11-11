// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellType represents ASN.1 CellType.
type CellType struct {
	CellSize          CellSize
	HasIEExtensions   bool
	IEExtensions      []CellTypeExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CellType", v) }

// DecodeCellType decodes a standalone APER value and reports the consumed octets.
func DecodeCellType(data []byte) (CellType, int, error) {
	var v CellType
	err := ngapcodec.DecodeParameter("CellType", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellType{}, 0, err
	}
	return v, len(data), err
}
