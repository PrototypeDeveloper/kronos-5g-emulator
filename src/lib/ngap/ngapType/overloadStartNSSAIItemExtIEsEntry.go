// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type OverloadStartNSSAIItemExtIEsEntryPresent uint32

const (
	OverloadStartNSSAIItemExtIEsEntryPresentNothing OverloadStartNSSAIItemExtIEsEntryPresent = iota
	OverloadStartNSSAIItemExtIEsEntryPresentUnknown
)

type OverloadStartNSSAIItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      OverloadStartNSSAIItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadStartNSSAIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadStartNSSAIItem-ExtIEsEntry", v)
}

// DecodeOverloadStartNSSAIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadStartNSSAIItemExtIEsEntry(data []byte) (OverloadStartNSSAIItemExtIEsEntry, int, error) {
	var v OverloadStartNSSAIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("OverloadStartNSSAIItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadStartNSSAIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
