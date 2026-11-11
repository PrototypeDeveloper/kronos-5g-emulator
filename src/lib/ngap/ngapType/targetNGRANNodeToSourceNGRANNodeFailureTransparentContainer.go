// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer represents ASN.1 TargetNGRANNode-ToSourceNGRANNode-FailureTransparentContainer.
type TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer struct {
	HasCellCAGInformation bool
	CellCAGInformation    CellCAGInformation
	HasIEExtensions       bool
	IEExtensions          []TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNGRANNode-ToSourceNGRANNode-FailureTransparentContainer", v)
}

// DecodeTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer(data []byte) (TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer, int, error) {
	var v TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer
	err := ngapcodec.DecodeParameter("TargetNGRANNode-ToSourceNGRANNode-FailureTransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
