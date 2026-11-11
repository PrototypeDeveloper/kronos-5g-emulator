// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBroadcastNRItem represents ASN.1 TAIBroadcastNR-Item.
type TAIBroadcastNRItem struct {
	TAI                   TAI
	CompletedCellsInTAINR CompletedCellsInTAINR
	HasIEExtensions       bool
	IEExtensions          []TAIBroadcastNRItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastNRItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastNR-Item", v)
}

// DecodeTAIBroadcastNRItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastNRItem(data []byte) (TAIBroadcastNRItem, int, error) {
	var v TAIBroadcastNRItem
	err := ngapcodec.DecodeParameter("TAIBroadcastNR-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastNRItem{}, 0, err
	}
	return v, len(data), err
}
