// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RequestedNSSAIItemExtIEsEntryPresent uint32

const (
	RequestedNSSAIItemExtIEsEntryPresentNothing RequestedNSSAIItemExtIEsEntryPresent = iota
	RequestedNSSAIItemExtIEsEntryPresentUnknown
)

type RequestedNSSAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RequestedNSSAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RequestedNSSAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RequestedNSSAIItem-ExtIEsEntry", v)
}

// DecodeRequestedNSSAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRequestedNSSAIItemExtIEsEntry(data []byte) (RequestedNSSAIItemExtIEsEntry, int, error) {
	var v RequestedNSSAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RequestedNSSAIItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RequestedNSSAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
