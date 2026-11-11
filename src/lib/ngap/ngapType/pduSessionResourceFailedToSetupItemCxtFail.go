// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupItemCxtFail represents ASN.1 PDUSessionResourceFailedToSetupItemCxtFail.
type PDUSessionResourceFailedToSetupItemCxtFail struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer []byte
	HasIEExtensions                             bool
	IEExtensions                                []PDUSessionResourceFailedToSetupItemCxtFailExtIEsEntry
	UnknownExtensions                           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemCxtFail) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemCxtFail", v)
}

// DecodePDUSessionResourceFailedToSetupItemCxtFail decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemCxtFail(data []byte) (PDUSessionResourceFailedToSetupItemCxtFail, int, error) {
	var v PDUSessionResourceFailedToSetupItemCxtFail
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemCxtFail", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemCxtFail{}, 0, err
	}
	return v, len(data), err
}
