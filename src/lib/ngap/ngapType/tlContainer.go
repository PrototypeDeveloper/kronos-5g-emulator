// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TLContainer represents ASN.1 TLContainer.
type TLContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TLContainer) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TLContainer", v) }

// DecodeTLContainer decodes a standalone APER value and reports the consumed octets.
func DecodeTLContainer(data []byte) (TLContainer, int, error) {
	var v TLContainer
	err := ngapcodec.DecodeParameter("TLContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return TLContainer{}, 0, err
	}
	return v, len(data), err
}
