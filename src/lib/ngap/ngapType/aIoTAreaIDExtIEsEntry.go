// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTAreaIDExtIEsEntryPresent uint32

const (
	AIoTAreaIDExtIEsEntryPresentNothing AIoTAreaIDExtIEsEntryPresent = iota
	AIoTAreaIDExtIEsEntryPresentUnknown
)

type AIoTAreaIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTAreaIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTAreaIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoTAreaID-ExtIEsEntry", v)
}

// DecodeAIoTAreaIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTAreaIDExtIEsEntry(data []byte) (AIoTAreaIDExtIEsEntry, int, error) {
	var v AIoTAreaIDExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoTAreaID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTAreaIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
