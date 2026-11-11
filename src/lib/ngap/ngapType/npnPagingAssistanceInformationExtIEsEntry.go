// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NPNPagingAssistanceInformationExtIEsEntryPresent uint32

const (
	NPNPagingAssistanceInformationExtIEsEntryPresentNothing NPNPagingAssistanceInformationExtIEsEntryPresent = iota
	NPNPagingAssistanceInformationExtIEsEntryPresentUnknown
)

type NPNPagingAssistanceInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NPNPagingAssistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNPagingAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-PagingAssistanceInformation-ExtIEsEntry", v)
}

// DecodeNPNPagingAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNPNPagingAssistanceInformationExtIEsEntry(data []byte) (NPNPagingAssistanceInformationExtIEsEntry, int, error) {
	var v NPNPagingAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("NPN-PagingAssistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNPagingAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
