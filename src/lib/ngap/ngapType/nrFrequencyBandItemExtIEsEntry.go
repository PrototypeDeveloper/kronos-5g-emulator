// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRFrequencyBandItemExtIEsEntryPresent uint32

const (
	NRFrequencyBandItemExtIEsEntryPresentNothing NRFrequencyBandItemExtIEsEntryPresent = iota
	NRFrequencyBandItemExtIEsEntryPresentUnknown
)

type NRFrequencyBandItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRFrequencyBandItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyBandItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyBandItem-ExtIEsEntry", v)
}

// DecodeNRFrequencyBandItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyBandItemExtIEsEntry(data []byte) (NRFrequencyBandItemExtIEsEntry, int, error) {
	var v NRFrequencyBandItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NRFrequencyBandItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyBandItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
