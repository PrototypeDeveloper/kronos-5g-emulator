// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleaseResponseTransferOCTETSTRING represents ASN.1 PDUSessionResourceReleaseResponseTransfer-OCTET-STRING.
type PDUSessionResourceReleaseResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleaseResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleaseResponseTransfer-OCTET-STRING", v)
}

// DecodePDUSessionResourceReleaseResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleaseResponseTransferOCTETSTRING(data []byte) (PDUSessionResourceReleaseResponseTransferOCTETSTRING, int, error) {
	var v PDUSessionResourceReleaseResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleaseResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleaseResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
