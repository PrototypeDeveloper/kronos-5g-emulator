// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SecurityIndicationExtIEsEntryPresent uint32

const (
	SecurityIndicationExtIEsEntryPresentNothing SecurityIndicationExtIEsEntryPresent = iota
	SecurityIndicationExtIEsEntryPresentMaximumIntegrityProtectedDataRateDL
	SecurityIndicationExtIEsEntryPresentUnknown
)

type SecurityIndicationExtIEsEntry struct {
	ID                                  ProtocolExtensionID
	Criticality                         Criticality
	Present                             SecurityIndicationExtIEsEntryPresent
	MaximumIntegrityProtectedDataRateDL MaximumIntegrityProtectedDataRate `ngap:"id:151,criticality:ignore,presence:optional"`
	UnknownValue                        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityIndication-ExtIEsEntry", v)
}

// DecodeSecurityIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityIndicationExtIEsEntry(data []byte) (SecurityIndicationExtIEsEntry, int, error) {
	var v SecurityIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("SecurityIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
