// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToModifyListModCfm represents ASN.1 PDUSessionResourceFailedToModifyListModCfm.
type PDUSessionResourceFailedToModifyListModCfm struct {
	List []PDUSessionResourceFailedToModifyItemModCfm `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyListModCfm) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyListModCfm", v)
}

// DecodePDUSessionResourceFailedToModifyListModCfm decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyListModCfm(data []byte) (PDUSessionResourceFailedToModifyListModCfm, int, error) {
	var v PDUSessionResourceFailedToModifyListModCfm
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyListModCfm", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyListModCfm{}, 0, err
	}
	return v, len(data), err
}
