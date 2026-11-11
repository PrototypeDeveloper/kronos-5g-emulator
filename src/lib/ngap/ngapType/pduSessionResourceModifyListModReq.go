// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyListModReq represents ASN.1 PDUSessionResourceModifyListModReq.
type PDUSessionResourceModifyListModReq struct {
	List []PDUSessionResourceModifyItemModReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyListModReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyListModReq", v)
}

// DecodePDUSessionResourceModifyListModReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyListModReq(data []byte) (PDUSessionResourceModifyListModReq, int, error) {
	var v PDUSessionResourceModifyListModReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyListModReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyListModReq{}, 0, err
	}
	return v, len(data), err
}
