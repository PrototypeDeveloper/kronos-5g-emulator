// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceResumeItemRESReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceResumeItemRESReqExtIEsEntryPresentNothing PDUSessionResourceResumeItemRESReqExtIEsEntryPresent = iota
	PDUSessionResourceResumeItemRESReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceResumeItemRESReqExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceResumeItemRESReqExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeItemRESReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeItemRESReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceResumeItemRESReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeItemRESReqExtIEsEntry(data []byte) (PDUSessionResourceResumeItemRESReqExtIEsEntry, int, error) {
	var v PDUSessionResourceResumeItemRESReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeItemRESReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeItemRESReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
