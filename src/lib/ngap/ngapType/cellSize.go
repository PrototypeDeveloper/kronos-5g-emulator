// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellSize represents ASN.1 CellSize.
const (
	CellSizePresentVerysmall Enumerated = 0
	CellSizePresentSmall     Enumerated = 1
	CellSizePresentMedium    Enumerated = 2
	CellSizePresentLarge     Enumerated = 3
)

type CellSize struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellSize) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CellSize", v) }

// DecodeCellSize decodes a standalone APER value and reports the consumed octets.
func DecodeCellSize(data []byte) (CellSize, int, error) {
	var v CellSize
	err := ngapcodec.DecodeParameter("CellSize", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellSize{}, 0, err
	}
	return v, len(data), err
}
