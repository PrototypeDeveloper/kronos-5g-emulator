// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAICancelledEUTRAItem represents ASN.1 TAICancelledEUTRA-Item.
type TAICancelledEUTRAItem struct {
	TAI                      TAI
	CancelledCellsInTAIEUTRA CancelledCellsInTAIEUTRA
	HasIEExtensions          bool
	IEExtensions             []TAICancelledEUTRAItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledEUTRA-Item", v)
}

// DecodeTAICancelledEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledEUTRAItem(data []byte) (TAICancelledEUTRAItem, int, error) {
	var v TAICancelledEUTRAItem
	err := ngapcodec.DecodeParameter("TAICancelledEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
