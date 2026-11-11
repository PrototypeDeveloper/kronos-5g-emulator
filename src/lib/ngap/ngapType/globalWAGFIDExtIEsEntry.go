// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalWAGFIDExtIEsEntryPresent uint32

const (
	GlobalWAGFIDExtIEsEntryPresentNothing GlobalWAGFIDExtIEsEntryPresent = iota
	GlobalWAGFIDExtIEsEntryPresentUnknown
)

type GlobalWAGFIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalWAGFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalWAGFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalW-AGF-ID-ExtIEsEntry", v)
}

// DecodeGlobalWAGFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalWAGFIDExtIEsEntry(data []byte) (GlobalWAGFIDExtIEsEntry, int, error) {
	var v GlobalWAGFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalW-AGF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalWAGFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
