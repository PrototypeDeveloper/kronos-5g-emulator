// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowPerTNLInformation represents ASN.1 QosFlowPerTNLInformation.
type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation
	AssociatedQosFlowList       AssociatedQosFlowList
	HasIEExtensions             bool
	IEExtensions                []QosFlowPerTNLInformationExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowPerTNLInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowPerTNLInformation", v)
}

// DecodeQosFlowPerTNLInformation decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowPerTNLInformation(data []byte) (QosFlowPerTNLInformation, int, error) {
	var v QosFlowPerTNLInformation
	err := ngapcodec.DecodeParameter("QosFlowPerTNLInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowPerTNLInformation{}, 0, err
	}
	return v, len(data), err
}
