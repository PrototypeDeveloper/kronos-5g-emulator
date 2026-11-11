// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIOTFNameExtIEsEntryPresent uint32

const (
	AIOTFNameExtIEsEntryPresentNothing AIOTFNameExtIEsEntryPresent = iota
	AIOTFNameExtIEsEntryPresentUnknown
)

type AIOTFNameExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIOTFNameExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTFNameExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTFName-ExtIEsEntry", v)
}

// DecodeAIOTFNameExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTFNameExtIEsEntry(data []byte) (AIOTFNameExtIEsEntry, int, error) {
	var v AIOTFNameExtIEsEntry
	err := ngapcodec.DecodeParameter("AIOTFName-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTFNameExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
