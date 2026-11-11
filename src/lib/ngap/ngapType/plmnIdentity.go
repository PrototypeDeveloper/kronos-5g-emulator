// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PLMNIdentity represents ASN.1 PLMNIdentity.
type PLMNIdentity struct {
	Value []byte `aper:"sizeLB:3,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNIdentity) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PLMNIdentity", v) }

// DecodePLMNIdentity decodes a standalone APER value and reports the consumed octets.
func DecodePLMNIdentity(data []byte) (PLMNIdentity, int, error) {
	var v PLMNIdentity
	err := ngapcodec.DecodeParameter("PLMNIdentity", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNIdentity{}, 0, err
	}
	return v, len(data), err
}
