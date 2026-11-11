// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToModifyListModRes represents ASN.1 PDUSessionResourceFailedToModifyListModRes.
type PDUSessionResourceFailedToModifyListModRes struct {
	List []PDUSessionResourceFailedToModifyItemModRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyListModRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyListModRes", v)
}

// DecodePDUSessionResourceFailedToModifyListModRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyListModRes(data []byte) (PDUSessionResourceFailedToModifyListModRes, int, error) {
	var v PDUSessionResourceFailedToModifyListModRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyListModRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyListModRes{}, 0, err
	}
	return v, len(data), err
}
