// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyRequestTransfer represents ASN.1 PDUSessionResourceModifyRequestTransfer.
type PDUSessionResourceModifyRequestTransfer struct {
	ProtocolIEs       []PDUSessionResourceModifyRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyRequestTransfer", v)
}

// DecodePDUSessionResourceModifyRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyRequestTransfer(data []byte) (PDUSessionResourceModifyRequestTransfer, int, error) {
	var v PDUSessionResourceModifyRequestTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
