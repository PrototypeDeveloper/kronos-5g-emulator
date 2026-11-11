// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceNGRANNodeToTargetNGRANNodeTransparentContainer represents ASN.1 SourceNGRANNode-ToTargetNGRANNode-TransparentContainer.
type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                         RRCContainer
	HasPDUSessionResourceInformationList bool
	PDUSessionResourceInformationList    PDUSessionResourceInformationList
	HasERABInformationList               bool
	ERABInformationList                  ERABInformationList
	TargetCellID                         NGRANCGI
	HasIndexToRFSP                       bool
	IndexToRFSP                          IndexToRFSP
	UEHistoryInformation                 UEHistoryInformation
	HasIEExtensions                      bool
	IEExtensions                         []SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceNGRANNode-ToTargetNGRANNode-TransparentContainer", v)
}

// DecodeSourceNGRANNodeToTargetNGRANNodeTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeSourceNGRANNodeToTargetNGRANNodeTransparentContainer(data []byte) (SourceNGRANNodeToTargetNGRANNodeTransparentContainer, int, error) {
	var v SourceNGRANNodeToTargetNGRANNodeTransparentContainer
	err := ngapcodec.DecodeParameter("SourceNGRANNode-ToTargetNGRANNode-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceNGRANNodeToTargetNGRANNodeTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
