// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupItemCxtRes represents ASN.1 PDUSessionResourceFailedToSetupItemCxtRes.
type PDUSessionResourceFailedToSetupItemCxtRes struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer []byte
	HasIEExtensions                             bool
	IEExtensions                                []PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry
	UnknownExtensions                           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemCxtRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemCxtRes", v)
}

// DecodePDUSessionResourceFailedToSetupItemCxtRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemCxtRes(data []byte) (PDUSessionResourceFailedToSetupItemCxtRes, int, error) {
	var v PDUSessionResourceFailedToSetupItemCxtRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemCxtRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemCxtRes{}, 0, err
	}
	return v, len(data), err
}
