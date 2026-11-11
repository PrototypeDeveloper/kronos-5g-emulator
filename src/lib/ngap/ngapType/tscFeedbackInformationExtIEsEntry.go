// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TSCFeedbackInformationExtIEsEntryPresent uint32

const (
	TSCFeedbackInformationExtIEsEntryPresentNothing TSCFeedbackInformationExtIEsEntryPresent = iota
	TSCFeedbackInformationExtIEsEntryPresentUnknown
)

type TSCFeedbackInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TSCFeedbackInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCFeedbackInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCFeedbackInformation-ExtIEsEntry", v)
}

// DecodeTSCFeedbackInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTSCFeedbackInformationExtIEsEntry(data []byte) (TSCFeedbackInformationExtIEsEntry, int, error) {
	var v TSCFeedbackInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("TSCFeedbackInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCFeedbackInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
