// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionUpdateRequestTransferOCTETSTRING represents ASN.1 MulticastSessionUpdateRequestTransfer-OCTET-STRING.
type MulticastSessionUpdateRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionUpdateRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionUpdateRequestTransfer-OCTET-STRING", v)
}

// DecodeMulticastSessionUpdateRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionUpdateRequestTransferOCTETSTRING(data []byte) (MulticastSessionUpdateRequestTransferOCTETSTRING, int, error) {
	var v MulticastSessionUpdateRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MulticastSessionUpdateRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionUpdateRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
