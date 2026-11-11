// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EmergencyAreaIDBroadcastNRItemExtIEsEntryPresent uint32

const (
	EmergencyAreaIDBroadcastNRItemExtIEsEntryPresentNothing EmergencyAreaIDBroadcastNRItemExtIEsEntryPresent = iota
	EmergencyAreaIDBroadcastNRItemExtIEsEntryPresentUnknown
)

type EmergencyAreaIDBroadcastNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EmergencyAreaIDBroadcastNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastNR-Item-ExtIEsEntry", v)
}

// DecodeEmergencyAreaIDBroadcastNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastNRItemExtIEsEntry(data []byte) (EmergencyAreaIDBroadcastNRItemExtIEsEntry, int, error) {
	var v EmergencyAreaIDBroadcastNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
