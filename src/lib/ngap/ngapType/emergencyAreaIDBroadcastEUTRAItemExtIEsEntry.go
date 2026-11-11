// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EmergencyAreaIDBroadcastEUTRAItemExtIEsEntryPresent uint32

const (
	EmergencyAreaIDBroadcastEUTRAItemExtIEsEntryPresentNothing EmergencyAreaIDBroadcastEUTRAItemExtIEsEntryPresent = iota
	EmergencyAreaIDBroadcastEUTRAItemExtIEsEntryPresentUnknown
)

type EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EmergencyAreaIDBroadcastEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDBroadcastEUTRA-Item-ExtIEsEntry", v)
}

// DecodeEmergencyAreaIDBroadcastEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDBroadcastEUTRAItemExtIEsEntry(data []byte) (EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry, int, error) {
	var v EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EmergencyAreaIDBroadcastEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDBroadcastEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
