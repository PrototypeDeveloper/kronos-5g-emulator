// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CompositeAvailableCapacityExtIEsEntryPresent uint32

const (
	CompositeAvailableCapacityExtIEsEntryPresentNothing CompositeAvailableCapacityExtIEsEntryPresent = iota
	CompositeAvailableCapacityExtIEsEntryPresentUnknown
)

type CompositeAvailableCapacityExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CompositeAvailableCapacityExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompositeAvailableCapacityExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompositeAvailableCapacity-ExtIEsEntry", v)
}

// DecodeCompositeAvailableCapacityExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCompositeAvailableCapacityExtIEsEntry(data []byte) (CompositeAvailableCapacityExtIEsEntry, int, error) {
	var v CompositeAvailableCapacityExtIEsEntry
	err := ngapcodec.DecodeParameter("CompositeAvailableCapacity-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompositeAvailableCapacityExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
