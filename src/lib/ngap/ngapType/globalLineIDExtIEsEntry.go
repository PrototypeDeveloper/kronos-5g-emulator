// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalLineIDExtIEsEntryPresent uint32

const (
	GlobalLineIDExtIEsEntryPresentNothing GlobalLineIDExtIEsEntryPresent = iota
	GlobalLineIDExtIEsEntryPresentTAI
	GlobalLineIDExtIEsEntryPresentUnknown
)

type GlobalLineIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalLineIDExtIEsEntryPresent
	TAI          TAI `ngap:"id:213,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalLineIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalLine-ID-ExtIEsEntry", v)
}

// DecodeGlobalLineIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalLineIDExtIEsEntry(data []byte) (GlobalLineIDExtIEsEntry, int, error) {
	var v GlobalLineIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalLine-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalLineIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
