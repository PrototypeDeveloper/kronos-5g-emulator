// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetToSourceTransparentContainer represents ASN.1 TargetToSource-TransparentContainer.
type TargetToSourceTransparentContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetToSourceTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetToSource-TransparentContainer", v)
}

// DecodeTargetToSourceTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeTargetToSourceTransparentContainer(data []byte) (TargetToSourceTransparentContainer, int, error) {
	var v TargetToSourceTransparentContainer
	err := ngapcodec.DecodeParameter("TargetToSource-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetToSourceTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
