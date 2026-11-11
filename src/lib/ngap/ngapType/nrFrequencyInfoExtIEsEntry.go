// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRFrequencyInfoExtIEsEntryPresent uint32

const (
	NRFrequencyInfoExtIEsEntryPresentNothing NRFrequencyInfoExtIEsEntryPresent = iota
	NRFrequencyInfoExtIEsEntryPresentUnknown
)

type NRFrequencyInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRFrequencyInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyInfo-ExtIEsEntry", v)
}

// DecodeNRFrequencyInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyInfoExtIEsEntry(data []byte) (NRFrequencyInfoExtIEsEntry, int, error) {
	var v NRFrequencyInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("NRFrequencyInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
