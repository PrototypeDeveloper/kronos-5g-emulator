// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalN3IWFIDExtIEsEntryPresent uint32

const (
	GlobalN3IWFIDExtIEsEntryPresentNothing GlobalN3IWFIDExtIEsEntryPresent = iota
	GlobalN3IWFIDExtIEsEntryPresentUnknown
)

type GlobalN3IWFIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalN3IWFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalN3IWFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalN3IWF-ID-ExtIEsEntry", v)
}

// DecodeGlobalN3IWFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalN3IWFIDExtIEsEntry(data []byte) (GlobalN3IWFIDExtIEsEntry, int, error) {
	var v GlobalN3IWFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalN3IWF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalN3IWFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
