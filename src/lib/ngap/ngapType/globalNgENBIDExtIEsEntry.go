// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalNgENBIDExtIEsEntryPresent uint32

const (
	GlobalNgENBIDExtIEsEntryPresentNothing GlobalNgENBIDExtIEsEntryPresent = iota
	GlobalNgENBIDExtIEsEntryPresentUnknown
)

type GlobalNgENBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalNgENBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalNgENBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalNgENB-ID-ExtIEsEntry", v)
}

// DecodeGlobalNgENBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalNgENBIDExtIEsEntry(data []byte) (GlobalNgENBIDExtIEsEntry, int, error) {
	var v GlobalNgENBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalNgENB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalNgENBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
