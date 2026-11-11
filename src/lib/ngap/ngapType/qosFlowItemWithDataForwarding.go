// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowItemWithDataForwarding represents ASN.1 QosFlowItemWithDataForwarding.
type QosFlowItemWithDataForwarding struct {
	QosFlowIdentifier         QosFlowIdentifier
	HasDataForwardingAccepted bool
	DataForwardingAccepted    DataForwardingAccepted
	HasIEExtensions           bool
	IEExtensions              []QosFlowItemWithDataForwardingExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowItemWithDataForwarding) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowItemWithDataForwarding", v)
}

// DecodeQosFlowItemWithDataForwarding decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowItemWithDataForwarding(data []byte) (QosFlowItemWithDataForwarding, int, error) {
	var v QosFlowItemWithDataForwarding
	err := ngapcodec.DecodeParameter("QosFlowItemWithDataForwarding", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowItemWithDataForwarding{}, 0, err
	}
	return v, len(data), err
}
