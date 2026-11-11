// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EarlyStatusTransferTransparentContainer represents ASN.1 EarlyStatusTransfer-TransparentContainer.
type EarlyStatusTransferTransparentContainer struct {
	ProcedureStage    ProcedureStageChoice
	HasIEExtensions   bool
	IEExtensions      []EarlyStatusTransferTransparentContainerExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EarlyStatusTransferTransparentContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EarlyStatusTransfer-TransparentContainer", v)
}

// DecodeEarlyStatusTransferTransparentContainer decodes a standalone APER value and reports the consumed octets.
func DecodeEarlyStatusTransferTransparentContainer(data []byte) (EarlyStatusTransferTransparentContainer, int, error) {
	var v EarlyStatusTransferTransparentContainer
	err := ngapcodec.DecodeParameter("EarlyStatusTransfer-TransparentContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return EarlyStatusTransferTransparentContainer{}, 0, err
	}
	return v, len(data), err
}
