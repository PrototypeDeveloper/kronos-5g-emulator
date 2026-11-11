// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverCommandTransfer represents ASN.1 HandoverCommandTransfer.
type HandoverCommandTransfer struct {
	HasDLForwardingUPTNLInformation  bool
	DLForwardingUPTNLInformation     UPTransportLayerInformation
	HasQosFlowToBeForwardedList      bool
	QosFlowToBeForwardedList         QosFlowToBeForwardedList
	HasDataForwardingResponseDRBList bool
	DataForwardingResponseDRBList    DataForwardingResponseDRBList
	HasIEExtensions                  bool
	IEExtensions                     []HandoverCommandTransferExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverCommandTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverCommandTransfer", v)
}

// DecodeHandoverCommandTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverCommandTransfer(data []byte) (HandoverCommandTransfer, int, error) {
	var v HandoverCommandTransfer
	err := ngapcodec.DecodeParameter("HandoverCommandTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverCommandTransfer{}, 0, err
	}
	return v, len(data), err
}
