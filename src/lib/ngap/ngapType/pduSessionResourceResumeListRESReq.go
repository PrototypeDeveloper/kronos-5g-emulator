// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceResumeListRESReq represents ASN.1 PDUSessionResourceResumeListRESReq.
type PDUSessionResourceResumeListRESReq struct {
	List []PDUSessionResourceResumeItemRESReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeListRESReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeListRESReq", v)
}

// DecodePDUSessionResourceResumeListRESReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeListRESReq(data []byte) (PDUSessionResourceResumeListRESReq, int, error) {
	var v PDUSessionResourceResumeListRESReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeListRESReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeListRESReq{}, 0, err
	}
	return v, len(data), err
}
