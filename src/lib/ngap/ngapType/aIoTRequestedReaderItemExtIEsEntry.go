// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTRequestedReaderItemExtIEsEntryPresent uint32

const (
	AIoTRequestedReaderItemExtIEsEntryPresentNothing AIoTRequestedReaderItemExtIEsEntryPresent = iota
	AIoTRequestedReaderItemExtIEsEntryPresentUnknown
)

type AIoTRequestedReaderItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTRequestedReaderItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedReaderItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedReaderItem-ExtIEsEntry", v)
}

// DecodeAIoTRequestedReaderItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedReaderItemExtIEsEntry(data []byte) (AIoTRequestedReaderItemExtIEsEntry, int, error) {
	var v AIoTRequestedReaderItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-RequestedReaderItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedReaderItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
