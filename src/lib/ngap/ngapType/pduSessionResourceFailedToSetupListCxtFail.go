// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupListCxtFail represents ASN.1 PDUSessionResourceFailedToSetupListCxtFail.
type PDUSessionResourceFailedToSetupListCxtFail struct {
	List []PDUSessionResourceFailedToSetupItemCxtFail `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupListCxtFail) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupListCxtFail", v)
}

// DecodePDUSessionResourceFailedToSetupListCxtFail decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupListCxtFail(data []byte) (PDUSessionResourceFailedToSetupListCxtFail, int, error) {
	var v PDUSessionResourceFailedToSetupListCxtFail
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupListCxtFail", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupListCxtFail{}, 0, err
	}
	return v, len(data), err
}
