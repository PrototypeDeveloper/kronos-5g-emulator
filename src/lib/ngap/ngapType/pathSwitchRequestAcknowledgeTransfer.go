// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PathSwitchRequestAcknowledgeTransfer represents ASN.1 PathSwitchRequestAcknowledgeTransfer.
type PathSwitchRequestAcknowledgeTransfer struct {
	HasULNGUUPTNLInformation bool
	ULNGUUPTNLInformation    UPTransportLayerInformation
	HasSecurityIndication    bool
	SecurityIndication       SecurityIndication
	HasIEExtensions          bool
	IEExtensions             []PathSwitchRequestAcknowledgeTransferExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestAcknowledgeTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestAcknowledgeTransfer", v)
}

// DecodePathSwitchRequestAcknowledgeTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestAcknowledgeTransfer(data []byte) (PathSwitchRequestAcknowledgeTransfer, int, error) {
	var v PathSwitchRequestAcknowledgeTransfer
	err := ngapcodec.DecodeParameter("PathSwitchRequestAcknowledgeTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestAcknowledgeTransfer{}, 0, err
	}
	return v, len(data), err
}
