// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToResumeListRESRes represents ASN.1 PDUSessionResourceFailedToResumeListRESRes.
type PDUSessionResourceFailedToResumeListRESRes struct {
	List []PDUSessionResourceFailedToResumeItemRESRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeListRESRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeListRESRes", v)
}

// DecodePDUSessionResourceFailedToResumeListRESRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeListRESRes(data []byte) (PDUSessionResourceFailedToResumeListRESRes, int, error) {
	var v PDUSessionResourceFailedToResumeListRESRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeListRESRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeListRESRes{}, 0, err
	}
	return v, len(data), err
}
