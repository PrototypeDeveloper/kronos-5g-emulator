// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionDeactivationRequestTransferOCTETSTRING represents ASN.1 MulticastSessionDeactivationRequestTransfer-OCTET-STRING.
type MulticastSessionDeactivationRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionDeactivationRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionDeactivationRequestTransfer-OCTET-STRING", v)
}

// DecodeMulticastSessionDeactivationRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionDeactivationRequestTransferOCTETSTRING(data []byte) (MulticastSessionDeactivationRequestTransferOCTETSTRING, int, error) {
	var v MulticastSessionDeactivationRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MulticastSessionDeactivationRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionDeactivationRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
