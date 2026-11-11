// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupRequestTransfer represents ASN.1 PDUSessionResourceSetupRequestTransfer.
type PDUSessionResourceSetupRequestTransfer struct {
	ProtocolIEs       []PDUSessionResourceSetupRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupRequestTransfer", v)
}

// DecodePDUSessionResourceSetupRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupRequestTransfer(data []byte) (PDUSessionResourceSetupRequestTransfer, int, error) {
	var v PDUSessionResourceSetupRequestTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
