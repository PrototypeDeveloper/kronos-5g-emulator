// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalTWIFIDExtIEsEntryPresent uint32

const (
	GlobalTWIFIDExtIEsEntryPresentNothing GlobalTWIFIDExtIEsEntryPresent = iota
	GlobalTWIFIDExtIEsEntryPresentUnknown
)

type GlobalTWIFIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalTWIFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalTWIFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalTWIF-ID-ExtIEsEntry", v)
}

// DecodeGlobalTWIFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalTWIFIDExtIEsEntry(data []byte) (GlobalTWIFIDExtIEsEntry, int, error) {
	var v GlobalTWIFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalTWIF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalTWIFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
