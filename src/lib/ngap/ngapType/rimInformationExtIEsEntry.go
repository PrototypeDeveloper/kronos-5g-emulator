// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RIMInformationExtIEsEntryPresent uint32

const (
	RIMInformationExtIEsEntryPresentNothing RIMInformationExtIEsEntryPresent = iota
	RIMInformationExtIEsEntryPresentUnknown
)

type RIMInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RIMInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RIMInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RIMInformation-ExtIEsEntry", v)
}

// DecodeRIMInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRIMInformationExtIEsEntry(data []byte) (RIMInformationExtIEsEntry, int, error) {
	var v RIMInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("RIMInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RIMInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
