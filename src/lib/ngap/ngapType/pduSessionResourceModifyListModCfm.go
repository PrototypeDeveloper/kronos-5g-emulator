// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyListModCfm represents ASN.1 PDUSessionResourceModifyListModCfm.
type PDUSessionResourceModifyListModCfm struct {
	List []PDUSessionResourceModifyItemModCfm `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyListModCfm) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyListModCfm", v)
}

// DecodePDUSessionResourceModifyListModCfm decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyListModCfm(data []byte) (PDUSessionResourceModifyListModCfm, int, error) {
	var v PDUSessionResourceModifyListModCfm
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyListModCfm", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyListModCfm{}, 0, err
	}
	return v, len(data), err
}
