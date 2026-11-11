// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UESliceMaximumBitRateItemExtIEsEntryPresent uint32

const (
	UESliceMaximumBitRateItemExtIEsEntryPresentNothing UESliceMaximumBitRateItemExtIEsEntryPresent = iota
	UESliceMaximumBitRateItemExtIEsEntryPresentUnknown
)

type UESliceMaximumBitRateItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UESliceMaximumBitRateItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UESliceMaximumBitRateItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UESliceMaximumBitRateItem-ExtIEsEntry", v)
}

// DecodeUESliceMaximumBitRateItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUESliceMaximumBitRateItemExtIEsEntry(data []byte) (UESliceMaximumBitRateItemExtIEsEntry, int, error) {
	var v UESliceMaximumBitRateItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UESliceMaximumBitRateItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UESliceMaximumBitRateItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
