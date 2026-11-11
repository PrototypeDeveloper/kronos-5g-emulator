// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleaseCommandTransfer represents ASN.1 PDUSessionResourceReleaseCommandTransfer.
type PDUSessionResourceReleaseCommandTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceReleaseCommandTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleaseCommandTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleaseCommandTransfer", v)
}

// DecodePDUSessionResourceReleaseCommandTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleaseCommandTransfer(data []byte) (PDUSessionResourceReleaseCommandTransfer, int, error) {
	var v PDUSessionResourceReleaseCommandTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleaseCommandTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleaseCommandTransfer{}, 0, err
	}
	return v, len(data), err
}
