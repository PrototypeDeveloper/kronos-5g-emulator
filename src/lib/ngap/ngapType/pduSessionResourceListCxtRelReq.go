// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceListCxtRelReq represents ASN.1 PDUSessionResourceListCxtRelReq.
type PDUSessionResourceListCxtRelReq struct {
	List []PDUSessionResourceItemCxtRelReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceListCxtRelReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceListCxtRelReq", v)
}

// DecodePDUSessionResourceListCxtRelReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceListCxtRelReq(data []byte) (PDUSessionResourceListCxtRelReq, int, error) {
	var v PDUSessionResourceListCxtRelReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceListCxtRelReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceListCxtRelReq{}, 0, err
	}
	return v, len(data), err
}
