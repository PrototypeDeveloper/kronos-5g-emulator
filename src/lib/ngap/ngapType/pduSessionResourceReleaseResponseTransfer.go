// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleaseResponseTransfer represents ASN.1 PDUSessionResourceReleaseResponseTransfer.
type PDUSessionResourceReleaseResponseTransfer struct {
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceReleaseResponseTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleaseResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleaseResponseTransfer", v)
}

// DecodePDUSessionResourceReleaseResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleaseResponseTransfer(data []byte) (PDUSessionResourceReleaseResponseTransfer, int, error) {
	var v PDUSessionResourceReleaseResponseTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleaseResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleaseResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
