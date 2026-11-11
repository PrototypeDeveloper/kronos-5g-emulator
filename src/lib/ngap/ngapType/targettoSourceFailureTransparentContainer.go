// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargettoSourceFailureTransparentContainer represents ASN.1 TargettoSource-Failure-TransparentContainer.
type TargettoSourceFailureTransparentContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargettoSourceFailureTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargettoSource-Failure-TransparentContainer", v)
}

// DecodeTargettoSourceFailureTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeTargettoSourceFailureTransparentContainer(data []byte) (TargettoSourceFailureTransparentContainer, int, error) {
	var v TargettoSourceFailureTransparentContainer
	err := ngapcodec.DecodeParameter("TargettoSource-Failure-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargettoSourceFailureTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
