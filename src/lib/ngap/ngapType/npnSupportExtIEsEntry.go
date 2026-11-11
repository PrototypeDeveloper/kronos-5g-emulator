// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NPNSupportExtIEsEntryPresent uint32

const (
	NPNSupportExtIEsEntryPresentNothing NPNSupportExtIEsEntryPresent = iota
	NPNSupportExtIEsEntryPresentUnknown
)

type NPNSupportExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NPNSupportExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNSupportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-Support-ExtIEsEntry", v)
}

// DecodeNPNSupportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNPNSupportExtIEsEntry(data []byte) (NPNSupportExtIEsEntry, int, error) {
	var v NPNSupportExtIEsEntry
	err := ngapcodec.DecodeParameter("NPN-Support-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNSupportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
