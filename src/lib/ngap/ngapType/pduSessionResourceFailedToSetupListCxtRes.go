// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupListCxtRes represents ASN.1 PDUSessionResourceFailedToSetupListCxtRes.
type PDUSessionResourceFailedToSetupListCxtRes struct {
	List []PDUSessionResourceFailedToSetupItemCxtRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupListCxtRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupListCxtRes", v)
}

// DecodePDUSessionResourceFailedToSetupListCxtRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupListCxtRes(data []byte) (PDUSessionResourceFailedToSetupListCxtRes, int, error) {
	var v PDUSessionResourceFailedToSetupListCxtRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupListCxtRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupListCxtRes{}, 0, err
	}
	return v, len(data), err
}
