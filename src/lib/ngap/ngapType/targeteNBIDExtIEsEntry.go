// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargeteNBIDExtIEsEntryPresent uint32

const (
	TargeteNBIDExtIEsEntryPresentNothing TargeteNBIDExtIEsEntryPresent = iota
	TargeteNBIDExtIEsEntryPresentUnknown
)

type TargeteNBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargeteNBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargeteNBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargeteNB-ID-ExtIEsEntry", v)
}

// DecodeTargeteNBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargeteNBIDExtIEsEntry(data []byte) (TargeteNBIDExtIEsEntry, int, error) {
	var v TargeteNBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TargeteNB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargeteNBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
