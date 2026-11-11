// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANCompositeAvailableCapacityGroupExtIEsEntryPresent uint32

const (
	EUTRANCompositeAvailableCapacityGroupExtIEsEntryPresentNothing EUTRANCompositeAvailableCapacityGroupExtIEsEntryPresent = iota
	EUTRANCompositeAvailableCapacityGroupExtIEsEntryPresentUnknown
)

type EUTRANCompositeAvailableCapacityGroupExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANCompositeAvailableCapacityGroupExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCompositeAvailableCapacityGroupExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CompositeAvailableCapacityGroup-ExtIEsEntry", v)
}

// DecodeEUTRANCompositeAvailableCapacityGroupExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCompositeAvailableCapacityGroupExtIEsEntry(data []byte) (EUTRANCompositeAvailableCapacityGroupExtIEsEntry, int, error) {
	var v EUTRANCompositeAvailableCapacityGroupExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-CompositeAvailableCapacityGroup-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCompositeAvailableCapacityGroupExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
