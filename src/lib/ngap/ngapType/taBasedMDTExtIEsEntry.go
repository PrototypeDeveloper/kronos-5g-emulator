// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TABasedMDTExtIEsEntryPresent uint32

const (
	TABasedMDTExtIEsEntryPresentNothing TABasedMDTExtIEsEntryPresent = iota
	TABasedMDTExtIEsEntryPresentUnknown
)

type TABasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TABasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TABasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TABasedMDT-ExtIEsEntry", v)
}

// DecodeTABasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTABasedMDTExtIEsEntry(data []byte) (TABasedMDTExtIEsEntry, int, error) {
	var v TABasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("TABasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TABasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
