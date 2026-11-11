// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIBasedQMCExtIEsEntryPresent uint32

const (
	TAIBasedQMCExtIEsEntryPresentNothing TAIBasedQMCExtIEsEntryPresent = iota
	TAIBasedQMCExtIEsEntryPresentUnknown
)

type TAIBasedQMCExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIBasedQMCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBasedQMCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBasedQMC-ExtIEsEntry", v)
}

// DecodeTAIBasedQMCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBasedQMCExtIEsEntry(data []byte) (TAIBasedQMCExtIEsEntry, int, error) {
	var v TAIBasedQMCExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIBasedQMC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBasedQMCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
