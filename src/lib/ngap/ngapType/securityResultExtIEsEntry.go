// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SecurityResultExtIEsEntryPresent uint32

const (
	SecurityResultExtIEsEntryPresentNothing SecurityResultExtIEsEntryPresent = iota
	SecurityResultExtIEsEntryPresentUnknown
)

type SecurityResultExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SecurityResultExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityResultExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityResult-ExtIEsEntry", v)
}

// DecodeSecurityResultExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityResultExtIEsEntry(data []byte) (SecurityResultExtIEsEntry, int, error) {
	var v SecurityResultExtIEsEntry
	err := ngapcodec.DecodeParameter("SecurityResult-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityResultExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
