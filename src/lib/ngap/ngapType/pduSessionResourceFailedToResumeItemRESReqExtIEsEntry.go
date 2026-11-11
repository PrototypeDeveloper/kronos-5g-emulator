// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToResumeItemRESReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToResumeItemRESReqExtIEsEntryPresentNothing PDUSessionResourceFailedToResumeItemRESReqExtIEsEntryPresent = iota
	PDUSessionResourceFailedToResumeItemRESReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToResumeItemRESReqExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeItemRESReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToResumeItemRESReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeItemRESReqExtIEsEntry(data []byte) (PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeItemRESReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
