// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PagingAttemptInformationExtIEsEntryPresent uint32

const (
	PagingAttemptInformationExtIEsEntryPresentNothing PagingAttemptInformationExtIEsEntryPresent = iota
	PagingAttemptInformationExtIEsEntryPresentUnknown
)

type PagingAttemptInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PagingAttemptInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingAttemptInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingAttemptInformation-ExtIEsEntry", v)
}

// DecodePagingAttemptInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePagingAttemptInformationExtIEsEntry(data []byte) (PagingAttemptInformationExtIEsEntry, int, error) {
	var v PagingAttemptInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("PagingAttemptInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingAttemptInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
