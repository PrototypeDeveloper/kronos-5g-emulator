// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PLMNAreaBasedQMCExtIEsEntryPresent uint32

const (
	PLMNAreaBasedQMCExtIEsEntryPresentNothing PLMNAreaBasedQMCExtIEsEntryPresent = iota
	PLMNAreaBasedQMCExtIEsEntryPresentUnknown
)

type PLMNAreaBasedQMCExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PLMNAreaBasedQMCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNAreaBasedQMCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNAreaBasedQMC-ExtIEsEntry", v)
}

// DecodePLMNAreaBasedQMCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePLMNAreaBasedQMCExtIEsEntry(data []byte) (PLMNAreaBasedQMCExtIEsEntry, int, error) {
	var v PLMNAreaBasedQMCExtIEsEntry
	err := ngapcodec.DecodeParameter("PLMNAreaBasedQMC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNAreaBasedQMCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
