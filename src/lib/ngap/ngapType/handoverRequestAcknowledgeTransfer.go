// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverRequestAcknowledgeTransfer represents ASN.1 HandoverRequestAcknowledgeTransfer.
type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation            UPTransportLayerInformation
	HasDLForwardingUPTNLInformation  bool
	DLForwardingUPTNLInformation     UPTransportLayerInformation
	HasSecurityResult                bool
	SecurityResult                   SecurityResult
	QosFlowSetupResponseList         QosFlowListWithDataForwarding
	HasQosFlowFailedToSetupList      bool
	QosFlowFailedToSetupList         QosFlowListWithCause
	HasDataForwardingResponseDRBList bool
	DataForwardingResponseDRBList    DataForwardingResponseDRBList
	HasIEExtensions                  bool
	IEExtensions                     []HandoverRequestAcknowledgeTransferExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverRequestAcknowledgeTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverRequestAcknowledgeTransfer", v)
}

// DecodeHandoverRequestAcknowledgeTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverRequestAcknowledgeTransfer(data []byte) (HandoverRequestAcknowledgeTransfer, int, error) {
	var v HandoverRequestAcknowledgeTransfer
	err := ngapcodec.DecodeParameter("HandoverRequestAcknowledgeTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverRequestAcknowledgeTransfer{}, 0, err
	}
	return v, len(data), err
}
