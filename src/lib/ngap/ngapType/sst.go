// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SST represents ASN.1 SST.
type SST struct {
	Value []byte `aper:"sizeLB:1,sizeUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SST) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SST", v) }

// DecodeSST decodes a standalone APER value and reports the consumed octets.
func DecodeSST(data []byte) (SST, int, error) {
	var v SST
	err := ngapcodec.DecodeParameter("SST", data, &v)
	if ngapcodec.IsFatal(err) {
		return SST{}, 0, err
	}
	return v, len(data), err
}
