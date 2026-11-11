// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToResumeItemRESResExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToResumeItemRESResExtIEsEntryPresentNothing PDUSessionResourceFailedToResumeItemRESResExtIEsEntryPresent = iota
	PDUSessionResourceFailedToResumeItemRESResExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToResumeItemRESResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToResumeItemRESResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeItemRESResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeItemRESRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToResumeItemRESResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeItemRESResExtIEsEntry(data []byte) (PDUSessionResourceFailedToResumeItemRESResExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToResumeItemRESResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeItemRESRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeItemRESResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
