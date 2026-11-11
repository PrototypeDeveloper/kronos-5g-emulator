// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalGNBIDExtIEsEntryPresent uint32

const (
	GlobalGNBIDExtIEsEntryPresentNothing GlobalGNBIDExtIEsEntryPresent = iota
	GlobalGNBIDExtIEsEntryPresentUnknown
)

type GlobalGNBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalGNBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalGNBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalGNB-ID-ExtIEsEntry", v)
}

// DecodeGlobalGNBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalGNBIDExtIEsEntry(data []byte) (GlobalGNBIDExtIEsEntry, int, error) {
	var v GlobalGNBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalGNB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalGNBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
