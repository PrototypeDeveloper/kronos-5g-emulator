// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBroadcastEUTRAItem represents ASN.1 TAIBroadcastEUTRA-Item.
type TAIBroadcastEUTRAItem struct {
	TAI                      TAI
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA
	HasIEExtensions          bool
	IEExtensions             []TAIBroadcastEUTRAItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastEUTRAItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastEUTRA-Item", v)
}

// DecodeTAIBroadcastEUTRAItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastEUTRAItem(data []byte) (TAIBroadcastEUTRAItem, int, error) {
	var v TAIBroadcastEUTRAItem
	err := ngapcodec.DecodeParameter("TAIBroadcastEUTRA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastEUTRAItem{}, 0, err
	}
	return v, len(data), err
}
