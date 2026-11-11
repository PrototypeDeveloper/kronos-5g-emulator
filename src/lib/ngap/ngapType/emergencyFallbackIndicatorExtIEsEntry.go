// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EmergencyFallbackIndicatorExtIEsEntryPresent uint32

const (
	EmergencyFallbackIndicatorExtIEsEntryPresentNothing EmergencyFallbackIndicatorExtIEsEntryPresent = iota
	EmergencyFallbackIndicatorExtIEsEntryPresentUnknown
)

type EmergencyFallbackIndicatorExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EmergencyFallbackIndicatorExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EmergencyFallbackIndicatorExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EmergencyFallbackIndicator-ExtIEsEntry", v)
}

// DecodeEmergencyFallbackIndicatorExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEmergencyFallbackIndicatorExtIEsEntry(data []byte) (EmergencyFallbackIndicatorExtIEsEntry, int, error) {
	var v EmergencyFallbackIndicatorExtIEsEntry
	err := ngapcodec.DecodeParameter("EmergencyFallbackIndicator-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EmergencyFallbackIndicatorExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
