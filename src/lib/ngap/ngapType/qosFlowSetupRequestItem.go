// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowSetupRequestItem represents ASN.1 QosFlowSetupRequestItem.
type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters QosFlowLevelQosParameters
	HasERABID                 bool
	ERABID                    ERABID
	HasIEExtensions           bool
	IEExtensions              []QosFlowSetupRequestItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowSetupRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowSetupRequestItem", v)
}

// DecodeQosFlowSetupRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowSetupRequestItem(data []byte) (QosFlowSetupRequestItem, int, error) {
	var v QosFlowSetupRequestItem
	err := ngapcodec.DecodeParameter("QosFlowSetupRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowSetupRequestItem{}, 0, err
	}
	return v, len(data), err
}
