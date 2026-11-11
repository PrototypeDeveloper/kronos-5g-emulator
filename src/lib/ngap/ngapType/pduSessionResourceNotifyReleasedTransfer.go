// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceNotifyReleasedTransfer represents ASN.1 PDUSessionResourceNotifyReleasedTransfer.
type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceNotifyReleasedTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyReleasedTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyReleasedTransfer", v)
}

// DecodePDUSessionResourceNotifyReleasedTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyReleasedTransfer(data []byte) (PDUSessionResourceNotifyReleasedTransfer, int, error) {
	var v PDUSessionResourceNotifyReleasedTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyReleasedTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyReleasedTransfer{}, 0, err
	}
	return v, len(data), err
}
