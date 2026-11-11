// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupItemSURes represents ASN.1 PDUSessionResourceSetupItemSURes.
type PDUSessionResourceSetupItemSURes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer []byte
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceSetupItemSUResExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemSURes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemSURes", v)
}

// DecodePDUSessionResourceSetupItemSURes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemSURes(data []byte) (PDUSessionResourceSetupItemSURes, int, error) {
	var v PDUSessionResourceSetupItemSURes
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemSURes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemSURes{}, 0, err
	}
	return v, len(data), err
}
