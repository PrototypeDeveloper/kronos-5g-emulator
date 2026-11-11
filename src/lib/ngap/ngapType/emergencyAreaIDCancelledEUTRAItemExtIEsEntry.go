// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EmergencyAreaIDCancelledEUTRAItemExtIEsEntryPresent uint32

const (
	EmergencyAreaIDCancelledEUTRAItemExtIEsEntryPresentNothing EmergencyAreaIDCancelledEUTRAItemExtIEsEntryPresent = iota
	EmergencyAreaIDCancelledEUTRAItemExtIEsEntryPresentUnknown
)

type EmergencyAreaIDCancelledEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EmergencyAreaIDCancelledEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledEUTRA-Item-ExtIEsEntry", v)
}

// DecodeEmergencyAreaIDCancelledEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledEUTRAItemExtIEsEntry(data []byte) (EmergencyAreaIDCancelledEUTRAItemExtIEsEntry, int, error) {
	var v EmergencyAreaIDCancelledEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
