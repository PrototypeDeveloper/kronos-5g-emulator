// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ULCPSecurityInformationExtIEsEntryPresent uint32

const (
	ULCPSecurityInformationExtIEsEntryPresentNothing ULCPSecurityInformationExtIEsEntryPresent = iota
	ULCPSecurityInformationExtIEsEntryPresentUnknown
)

type ULCPSecurityInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ULCPSecurityInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULCPSecurityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UL-CP-SecurityInformation-ExtIEsEntry", v)
}

// DecodeULCPSecurityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeULCPSecurityInformationExtIEsEntry(data []byte) (ULCPSecurityInformationExtIEsEntry, int, error) {
	var v ULCPSecurityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("UL-CP-SecurityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULCPSecurityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
