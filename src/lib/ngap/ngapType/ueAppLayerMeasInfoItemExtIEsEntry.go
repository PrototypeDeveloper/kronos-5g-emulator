// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEAppLayerMeasInfoItemExtIEsEntryPresent uint32

const (
	UEAppLayerMeasInfoItemExtIEsEntryPresentNothing UEAppLayerMeasInfoItemExtIEsEntryPresent = iota
	UEAppLayerMeasInfoItemExtIEsEntryPresentUnknown
)

type UEAppLayerMeasInfoItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEAppLayerMeasInfoItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAppLayerMeasInfoItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAppLayerMeasInfoItem-ExtIEsEntry", v)
}

// DecodeUEAppLayerMeasInfoItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEAppLayerMeasInfoItemExtIEsEntry(data []byte) (UEAppLayerMeasInfoItemExtIEsEntry, int, error) {
	var v UEAppLayerMeasInfoItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UEAppLayerMeasInfoItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAppLayerMeasInfoItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
