// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemMobilityFailureforVoiceFallbackExtIEsEntryPresent uint32

const (
	IntersystemMobilityFailureforVoiceFallbackExtIEsEntryPresentNothing IntersystemMobilityFailureforVoiceFallbackExtIEsEntryPresent = iota
	IntersystemMobilityFailureforVoiceFallbackExtIEsEntryPresentUnknown
)

type IntersystemMobilityFailureforVoiceFallbackExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemMobilityFailureforVoiceFallbackExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemMobilityFailureforVoiceFallbackExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemMobilityFailureforVoiceFallback-ExtIEsEntry", v)
}

// DecodeIntersystemMobilityFailureforVoiceFallbackExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemMobilityFailureforVoiceFallbackExtIEsEntry(data []byte) (IntersystemMobilityFailureforVoiceFallbackExtIEsEntry, int, error) {
	var v IntersystemMobilityFailureforVoiceFallbackExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemMobilityFailureforVoiceFallback-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemMobilityFailureforVoiceFallbackExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
