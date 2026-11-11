// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToResumeListRESReq represents ASN.1 PDUSessionResourceFailedToResumeListRESReq.
type PDUSessionResourceFailedToResumeListRESReq struct {
	List []PDUSessionResourceFailedToResumeItemRESReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeListRESReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeListRESReq", v)
}

// DecodePDUSessionResourceFailedToResumeListRESReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeListRESReq(data []byte) (PDUSessionResourceFailedToResumeListRESReq, int, error) {
	var v PDUSessionResourceFailedToResumeListRESReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeListRESReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeListRESReq{}, 0, err
	}
	return v, len(data), err
}
