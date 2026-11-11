// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIBasedMDTExtIEsEntryPresent uint32

const (
	TAIBasedMDTExtIEsEntryPresentNothing TAIBasedMDTExtIEsEntryPresent = iota
	TAIBasedMDTExtIEsEntryPresentUnknown
)

type TAIBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBasedMDT-ExtIEsEntry", v)
}

// DecodeTAIBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBasedMDTExtIEsEntry(data []byte) (TAIBasedMDTExtIEsEntry, int, error) {
	var v TAIBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIBasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
