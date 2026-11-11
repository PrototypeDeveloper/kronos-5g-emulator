// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SecurityContextExtIEsEntryPresent uint32

const (
	SecurityContextExtIEsEntryPresentNothing SecurityContextExtIEsEntryPresent = iota
	SecurityContextExtIEsEntryPresentUnknown
)

type SecurityContextExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SecurityContextExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityContextExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityContext-ExtIEsEntry", v)
}

// DecodeSecurityContextExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityContextExtIEsEntry(data []byte) (SecurityContextExtIEsEntry, int, error) {
	var v SecurityContextExtIEsEntry
	err := ngapcodec.DecodeParameter("SecurityContext-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityContextExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
