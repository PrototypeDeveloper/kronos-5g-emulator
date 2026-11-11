// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToResumeItemRESReq represents ASN.1 PDUSessionResourceFailedToResumeItemRESReq.
type PDUSessionResourceFailedToResumeItemRESReq struct {
	PDUSessionID      PDUSessionID
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceFailedToResumeItemRESReqExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeItemRESReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeItemRESReq", v)
}

// DecodePDUSessionResourceFailedToResumeItemRESReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeItemRESReq(data []byte) (PDUSessionResourceFailedToResumeItemRESReq, int, error) {
	var v PDUSessionResourceFailedToResumeItemRESReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeItemRESReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeItemRESReq{}, 0, err
	}
	return v, len(data), err
}
