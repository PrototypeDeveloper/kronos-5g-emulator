// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupListCxtRes represents ASN.1 PDUSessionResourceSetupListCxtRes.
type PDUSessionResourceSetupListCxtRes struct {
	List []PDUSessionResourceSetupItemCxtRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupListCxtRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupListCxtRes", v)
}

// DecodePDUSessionResourceSetupListCxtRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupListCxtRes(data []byte) (PDUSessionResourceSetupListCxtRes, int, error) {
	var v PDUSessionResourceSetupListCxtRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupListCxtRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupListCxtRes{}, 0, err
	}
	return v, len(data), err
}
