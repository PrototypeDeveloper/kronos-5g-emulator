// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EPSTAIExtIEsEntryPresent uint32

const (
	EPSTAIExtIEsEntryPresentNothing EPSTAIExtIEsEntryPresent = iota
	EPSTAIExtIEsEntryPresentUnknown
)

type EPSTAIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EPSTAIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EPSTAIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EPS-TAI-ExtIEsEntry", v)
}

// DecodeEPSTAIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEPSTAIExtIEsEntry(data []byte) (EPSTAIExtIEsEntry, int, error) {
	var v EPSTAIExtIEsEntry
	err := ngapcodec.DecodeParameter("EPS-TAI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EPSTAIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
