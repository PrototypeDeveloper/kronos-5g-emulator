// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TABasedQMCExtIEsEntryPresent uint32

const (
	TABasedQMCExtIEsEntryPresentNothing TABasedQMCExtIEsEntryPresent = iota
	TABasedQMCExtIEsEntryPresentUnknown
)

type TABasedQMCExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TABasedQMCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TABasedQMCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TABasedQMC-ExtIEsEntry", v)
}

// DecodeTABasedQMCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTABasedQMCExtIEsEntry(data []byte) (TABasedQMCExtIEsEntry, int, error) {
	var v TABasedQMCExtIEsEntry
	err := ngapcodec.DecodeParameter("TABasedQMC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TABasedQMCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
