// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANStatusTransferTransparentContainer represents ASN.1 RANStatusTransfer-TransparentContainer.
type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList DRBsSubjectToStatusTransferList
	HasIEExtensions                 bool
	IEExtensions                    []RANStatusTransferTransparentContainerExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANStatusTransferTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANStatusTransfer-TransparentContainer", v)
}

// DecodeRANStatusTransferTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeRANStatusTransferTransparentContainer(data []byte) (RANStatusTransferTransparentContainer, int, error) {
	var v RANStatusTransferTransparentContainer
	err := ngapcodec.DecodeParameter("RANStatusTransfer-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANStatusTransferTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
