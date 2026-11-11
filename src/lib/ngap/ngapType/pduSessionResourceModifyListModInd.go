// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyListModInd represents ASN.1 PDUSessionResourceModifyListModInd.
type PDUSessionResourceModifyListModInd struct {
	List []PDUSessionResourceModifyItemModInd `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyListModInd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyListModInd", v)
}

// DecodePDUSessionResourceModifyListModInd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyListModInd(data []byte) (PDUSessionResourceModifyListModInd, int, error) {
	var v PDUSessionResourceModifyListModInd
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyListModInd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyListModInd{}, 0, err
	}
	return v, len(data), err
}
