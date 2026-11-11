// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionActivationRequestTransferOCTETSTRING represents ASN.1 MulticastSessionActivationRequestTransfer-OCTET-STRING.
type MulticastSessionActivationRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionActivationRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionActivationRequestTransfer-OCTET-STRING", v)
}

// DecodeMulticastSessionActivationRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionActivationRequestTransferOCTETSTRING(data []byte) (MulticastSessionActivationRequestTransferOCTETSTRING, int, error) {
	var v MulticastSessionActivationRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MulticastSessionActivationRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionActivationRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
