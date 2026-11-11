// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GeographyBasedMDTExtIEsEntryPresent uint32

const (
	GeographyBasedMDTExtIEsEntryPresentNothing GeographyBasedMDTExtIEsEntryPresent = iota
	GeographyBasedMDTExtIEsEntryPresentUnknown
)

type GeographyBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GeographyBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GeographyBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GeographyBasedMDT-ExtIEsEntry", v)
}

// DecodeGeographyBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGeographyBasedMDTExtIEsEntry(data []byte) (GeographyBasedMDTExtIEsEntry, int, error) {
	var v GeographyBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("GeographyBasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GeographyBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
