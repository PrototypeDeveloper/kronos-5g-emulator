// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyIndicationUnsuccessfulTransfer represents ASN.1 PDUSessionResourceModifyIndicationUnsuccessfulTransfer.
type PDUSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyIndicationUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyIndicationUnsuccessfulTransfer", v)
}

// DecodePDUSessionResourceModifyIndicationUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyIndicationUnsuccessfulTransfer(data []byte) (PDUSessionResourceModifyIndicationUnsuccessfulTransfer, int, error) {
	var v PDUSessionResourceModifyIndicationUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyIndicationUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
