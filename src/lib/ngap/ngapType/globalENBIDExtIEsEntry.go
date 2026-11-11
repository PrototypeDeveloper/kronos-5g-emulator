// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalENBIDExtIEsEntryPresent uint32

const (
	GlobalENBIDExtIEsEntryPresentNothing GlobalENBIDExtIEsEntryPresent = iota
	GlobalENBIDExtIEsEntryPresentUnknown
)

type GlobalENBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalENBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalENBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalENB-ID-ExtIEsEntry", v)
}

// DecodeGlobalENBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalENBIDExtIEsEntry(data []byte) (GlobalENBIDExtIEsEntry, int, error) {
	var v GlobalENBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalENB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalENBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
