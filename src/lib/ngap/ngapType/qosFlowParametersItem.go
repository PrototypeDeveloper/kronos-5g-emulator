// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowParametersItem represents ASN.1 QosFlowParametersItem.
type QosFlowParametersItem struct {
	QosFlowIdentifier            QosFlowIdentifier
	HasAlternativeQoSParaSetList bool
	AlternativeQoSParaSetList    AlternativeQoSParaSetList
	HasIEExtensions              bool
	IEExtensions                 []QosFlowParametersItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowParametersItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowParametersItem", v)
}

// DecodeQosFlowParametersItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowParametersItem(data []byte) (QosFlowParametersItem, int, error) {
	var v QosFlowParametersItem
	err := ngapcodec.DecodeParameter("QosFlowParametersItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowParametersItem{}, 0, err
	}
	return v, len(data), err
}
