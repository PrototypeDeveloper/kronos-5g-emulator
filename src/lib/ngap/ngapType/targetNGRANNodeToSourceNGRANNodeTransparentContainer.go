// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetNGRANNodeToSourceNGRANNodeTransparentContainer represents ASN.1 TargetNGRANNode-ToSourceNGRANNode-TransparentContainer.
type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer      RRCContainer
	HasIEExtensions   bool
	IEExtensions      []TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNGRANNode-ToSourceNGRANNode-TransparentContainer", v)
}

// DecodeTargetNGRANNodeToSourceNGRANNodeTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNGRANNodeToSourceNGRANNodeTransparentContainer(data []byte) (TargetNGRANNodeToSourceNGRANNodeTransparentContainer, int, error) {
	var v TargetNGRANNodeToSourceNGRANNodeTransparentContainer
	err := ngapcodec.DecodeParameter("TargetNGRANNode-ToSourceNGRANNode-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNGRANNodeToSourceNGRANNodeTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
