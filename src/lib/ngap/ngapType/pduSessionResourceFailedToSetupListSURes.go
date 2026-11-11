// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupListSURes represents ASN.1 PDUSessionResourceFailedToSetupListSURes.
type PDUSessionResourceFailedToSetupListSURes struct {
	List []PDUSessionResourceFailedToSetupItemSURes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupListSURes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupListSURes", v)
}

// DecodePDUSessionResourceFailedToSetupListSURes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupListSURes(data []byte) (PDUSessionResourceFailedToSetupListSURes, int, error) {
	var v PDUSessionResourceFailedToSetupListSURes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupListSURes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupListSURes{}, 0, err
	}
	return v, len(data), err
}
