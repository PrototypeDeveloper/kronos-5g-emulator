// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaOfInterestTAIItemExtIEsEntryPresent uint32

const (
	AreaOfInterestTAIItemExtIEsEntryPresentNothing AreaOfInterestTAIItemExtIEsEntryPresent = iota
	AreaOfInterestTAIItemExtIEsEntryPresentUnknown
)

type AreaOfInterestTAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AreaOfInterestTAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestTAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestTAIItem-ExtIEsEntry", v)
}

// DecodeAreaOfInterestTAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestTAIItemExtIEsEntry(data []byte) (AreaOfInterestTAIItemExtIEsEntry, int, error) {
	var v AreaOfInterestTAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaOfInterestTAIItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestTAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
