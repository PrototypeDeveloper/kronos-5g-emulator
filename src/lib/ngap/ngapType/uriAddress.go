// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// URIAddress represents ASN.1 URI-address.
type URIAddress struct {
	Value string
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v URIAddress) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("URI-address", v) }

// DecodeURIAddress decodes a standalone APER value and reports the consumed octets.
func DecodeURIAddress(data []byte) (URIAddress, int, error) {
	var v URIAddress
	err := ngapcodec.DecodeParameter("URI-address", data, &v)
	if ngapcodec.IsFatal(err) {
		return URIAddress{}, 0, err
	}
	return v, len(data), err
}
