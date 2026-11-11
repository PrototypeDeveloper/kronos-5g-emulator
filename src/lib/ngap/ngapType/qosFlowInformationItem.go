// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowInformationItem represents ASN.1 QosFlowInformationItem.
type QosFlowInformationItem struct {
	QosFlowIdentifier QosFlowIdentifier
	HasDLForwarding   bool
	DLForwarding      DLForwarding
	HasIEExtensions   bool
	IEExtensions      []QosFlowInformationItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowInformationItem", v)
}

// DecodeQosFlowInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowInformationItem(data []byte) (QosFlowInformationItem, int, error) {
	var v QosFlowInformationItem
	err := ngapcodec.DecodeParameter("QosFlowInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowInformationItem{}, 0, err
	}
	return v, len(data), err
}
