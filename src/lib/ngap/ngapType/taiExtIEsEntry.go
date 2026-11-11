// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIExtIEsEntryPresent uint32

const (
	TAIExtIEsEntryPresentNothing TAIExtIEsEntryPresent = iota
	TAIExtIEsEntryPresentUnknown
)

type TAIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAI-ExtIEsEntry", v)
}

// DecodeTAIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIExtIEsEntry(data []byte) (TAIExtIEsEntry, int, error) {
	var v TAIExtIEsEntry
	err := ngapcodec.DecodeParameter("TAI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
