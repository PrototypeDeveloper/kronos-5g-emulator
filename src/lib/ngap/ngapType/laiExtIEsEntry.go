// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LAIExtIEsEntryPresent uint32

const (
	LAIExtIEsEntryPresentNothing LAIExtIEsEntryPresent = iota
	LAIExtIEsEntryPresentUnknown
)

type LAIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LAIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LAIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LAI-ExtIEsEntry", v)
}

// DecodeLAIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLAIExtIEsEntry(data []byte) (LAIExtIEsEntry, int, error) {
	var v LAIExtIEsEntry
	err := ngapcodec.DecodeParameter("LAI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LAIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
