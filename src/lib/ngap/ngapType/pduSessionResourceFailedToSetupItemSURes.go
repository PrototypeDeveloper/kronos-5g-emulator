// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupItemSURes represents ASN.1 PDUSessionResourceFailedToSetupItemSURes.
type PDUSessionResourceFailedToSetupItemSURes struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer []byte
	HasIEExtensions                             bool
	IEExtensions                                []PDUSessionResourceFailedToSetupItemSUResExtIEsEntry
	UnknownExtensions                           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemSURes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemSURes", v)
}

// DecodePDUSessionResourceFailedToSetupItemSURes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemSURes(data []byte) (PDUSessionResourceFailedToSetupItemSURes, int, error) {
	var v PDUSessionResourceFailedToSetupItemSURes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemSURes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemSURes{}, 0, err
	}
	return v, len(data), err
}
