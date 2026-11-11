// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DLCPSecurityInformationExtIEsEntryPresent uint32

const (
	DLCPSecurityInformationExtIEsEntryPresentNothing DLCPSecurityInformationExtIEsEntryPresent = iota
	DLCPSecurityInformationExtIEsEntryPresentUnknown
)

type DLCPSecurityInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DLCPSecurityInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLCPSecurityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DL-CP-SecurityInformation-ExtIEsEntry", v)
}

// DecodeDLCPSecurityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDLCPSecurityInformationExtIEsEntry(data []byte) (DLCPSecurityInformationExtIEsEntry, int, error) {
	var v DLCPSecurityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("DL-CP-SecurityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLCPSecurityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
