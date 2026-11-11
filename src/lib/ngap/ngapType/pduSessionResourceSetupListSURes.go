// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupListSURes represents ASN.1 PDUSessionResourceSetupListSURes.
type PDUSessionResourceSetupListSURes struct {
	List []PDUSessionResourceSetupItemSURes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupListSURes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupListSURes", v)
}

// DecodePDUSessionResourceSetupListSURes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupListSURes(data []byte) (PDUSessionResourceSetupListSURes, int, error) {
	var v PDUSessionResourceSetupListSURes
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupListSURes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupListSURes{}, 0, err
	}
	return v, len(data), err
}
