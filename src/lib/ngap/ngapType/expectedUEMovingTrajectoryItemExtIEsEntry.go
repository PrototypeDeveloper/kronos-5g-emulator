// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExpectedUEMovingTrajectoryItemExtIEsEntryPresent uint32

const (
	ExpectedUEMovingTrajectoryItemExtIEsEntryPresentNothing ExpectedUEMovingTrajectoryItemExtIEsEntryPresent = iota
	ExpectedUEMovingTrajectoryItemExtIEsEntryPresentUnknown
)

type ExpectedUEMovingTrajectoryItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExpectedUEMovingTrajectoryItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEMovingTrajectoryItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEMovingTrajectoryItem-ExtIEsEntry", v)
}

// DecodeExpectedUEMovingTrajectoryItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEMovingTrajectoryItemExtIEsEntry(data []byte) (ExpectedUEMovingTrajectoryItemExtIEsEntry, int, error) {
	var v ExpectedUEMovingTrajectoryItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ExpectedUEMovingTrajectoryItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEMovingTrajectoryItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
