// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalTNGFIDExtIEsEntryPresent uint32

const (
	GlobalTNGFIDExtIEsEntryPresentNothing GlobalTNGFIDExtIEsEntryPresent = iota
	GlobalTNGFIDExtIEsEntryPresentUnknown
)

type GlobalTNGFIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalTNGFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalTNGFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalTNGF-ID-ExtIEsEntry", v)
}

// DecodeGlobalTNGFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalTNGFIDExtIEsEntry(data []byte) (GlobalTNGFIDExtIEsEntry, int, error) {
	var v GlobalTNGFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalTNGF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalTNGFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
