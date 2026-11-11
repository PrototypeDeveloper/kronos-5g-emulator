// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAICancelledNRItem represents ASN.1 TAICancelledNR-Item.
type TAICancelledNRItem struct {
	TAI                   TAI
	CancelledCellsInTAINR CancelledCellsInTAINR
	HasIEExtensions       bool
	IEExtensions          []TAICancelledNRItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledNR-Item", v)
}

// DecodeTAICancelledNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledNRItem(data []byte) (TAICancelledNRItem, int, error) {
	var v TAICancelledNRItem
	err := ngapcodec.DecodeParameter("TAICancelledNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledNRItem{}, 0, err
	}
	return v, len(data), err
}
