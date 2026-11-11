// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupItemCxtRes represents ASN.1 PDUSessionResourceSetupItemCxtRes.
type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer []byte
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceSetupItemCxtResExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemCxtRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemCxtRes", v)
}

// DecodePDUSessionResourceSetupItemCxtRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemCxtRes(data []byte) (PDUSessionResourceSetupItemCxtRes, int, error) {
	var v PDUSessionResourceSetupItemCxtRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemCxtRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemCxtRes{}, 0, err
	}
	return v, len(data), err
}
