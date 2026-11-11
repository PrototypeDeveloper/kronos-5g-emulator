// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyListModRes represents ASN.1 PDUSessionResourceModifyListModRes.
type PDUSessionResourceModifyListModRes struct {
	List []PDUSessionResourceModifyItemModRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyListModRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyListModRes", v)
}

// DecodePDUSessionResourceModifyListModRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyListModRes(data []byte) (PDUSessionResourceModifyListModRes, int, error) {
	var v PDUSessionResourceModifyListModRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyListModRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyListModRes{}, 0, err
	}
	return v, len(data), err
}
