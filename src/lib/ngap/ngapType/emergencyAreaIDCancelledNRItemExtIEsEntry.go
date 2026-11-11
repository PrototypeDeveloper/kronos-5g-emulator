// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EmergencyAreaIDCancelledNRItemExtIEsEntryPresent uint32

const (
	EmergencyAreaIDCancelledNRItemExtIEsEntryPresentNothing EmergencyAreaIDCancelledNRItemExtIEsEntryPresent = iota
	EmergencyAreaIDCancelledNRItemExtIEsEntryPresentUnknown
)

type EmergencyAreaIDCancelledNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EmergencyAreaIDCancelledNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyAreaIDCancelledNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyAreaIDCancelledNR-Item-ExtIEsEntry", v)
}

// DecodeEmergencyAreaIDCancelledNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyAreaIDCancelledNRItemExtIEsEntry(data []byte) (EmergencyAreaIDCancelledNRItemExtIEsEntry, int, error) {
	var v EmergencyAreaIDCancelledNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EmergencyAreaIDCancelledNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyAreaIDCancelledNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
