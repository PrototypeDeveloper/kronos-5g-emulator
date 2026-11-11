// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceToTargetTransparentContainer represents ASN.1 SourceToTarget-TransparentContainer.
type SourceToTargetTransparentContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceToTargetTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceToTarget-TransparentContainer", v)
}

// DecodeSourceToTargetTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeSourceToTargetTransparentContainer(data []byte) (SourceToTargetTransparentContainer, int, error) {
	var v SourceToTargetTransparentContainer
	err := ngapcodec.DecodeParameter("SourceToTarget-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceToTargetTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
