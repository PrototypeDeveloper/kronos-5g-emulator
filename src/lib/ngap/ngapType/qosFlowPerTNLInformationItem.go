// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowPerTNLInformationItem represents ASN.1 QosFlowPerTNLInformationItem.
type QosFlowPerTNLInformationItem struct {
	QosFlowPerTNLInformation QosFlowPerTNLInformation
	HasIEExtensions          bool
	IEExtensions             []QosFlowPerTNLInformationItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowPerTNLInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowPerTNLInformationItem", v)
}

// DecodeQosFlowPerTNLInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowPerTNLInformationItem(data []byte) (QosFlowPerTNLInformationItem, int, error) {
	var v QosFlowPerTNLInformationItem
	err := ngapcodec.DecodeParameter("QosFlowPerTNLInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowPerTNLInformationItem{}, 0, err
	}
	return v, len(data), err
}
