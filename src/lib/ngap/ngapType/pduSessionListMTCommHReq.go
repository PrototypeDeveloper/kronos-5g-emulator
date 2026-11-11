// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionListMTCommHReq represents ASN.1 PDUSessionListMTCommHReq.
type PDUSessionListMTCommHReq struct {
	List []PDUSessionListMTCommHReqItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionListMTCommHReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionListMTCommHReq", v)
}

// DecodePDUSessionListMTCommHReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionListMTCommHReq(data []byte) (PDUSessionListMTCommHReq, int, error) {
	var v PDUSessionListMTCommHReq
	err := ngapcodec.DecodeParameter("PDUSessionListMTCommHReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionListMTCommHReq{}, 0, err
	}
	return v, len(data), err
}
