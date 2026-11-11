// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceResumeListRESRes represents ASN.1 PDUSessionResourceResumeListRESRes.
type PDUSessionResourceResumeListRESRes struct {
	List []PDUSessionResourceResumeItemRESRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeListRESRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeListRESRes", v)
}

// DecodePDUSessionResourceResumeListRESRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeListRESRes(data []byte) (PDUSessionResourceResumeListRESRes, int, error) {
	var v PDUSessionResourceResumeListRESRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeListRESRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeListRESRes{}, 0, err
	}
	return v, len(data), err
}
