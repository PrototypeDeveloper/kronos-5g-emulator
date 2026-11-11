// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GUAMIExtIEsEntryPresent uint32

const (
	GUAMIExtIEsEntryPresentNothing GUAMIExtIEsEntryPresent = iota
	GUAMIExtIEsEntryPresentUnknown
)

type GUAMIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GUAMIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GUAMI-ExtIEsEntry", v)
}

// DecodeGUAMIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMIExtIEsEntry(data []byte) (GUAMIExtIEsEntry, int, error) {
	var v GUAMIExtIEsEntry
	err := ngapcodec.DecodeParameter("GUAMI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
