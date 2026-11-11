// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUsetQoSInformationExtIEsEntryPresent uint32

const (
	PDUsetQoSInformationExtIEsEntryPresentNothing PDUsetQoSInformationExtIEsEntryPresent = iota
	PDUsetQoSInformationExtIEsEntryPresentUnknown
)

type PDUsetQoSInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUsetQoSInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUsetQoSInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUsetQoSInformation-ExtIEsEntry", v)
}

// DecodePDUsetQoSInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUsetQoSInformationExtIEsEntry(data []byte) (PDUsetQoSInformationExtIEsEntry, int, error) {
	var v PDUsetQoSInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUsetQoSInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUsetQoSInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
