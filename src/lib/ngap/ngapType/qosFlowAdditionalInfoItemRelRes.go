// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAdditionalInfoItemRelRes represents ASN.1 QosFlowAdditionalInfoItemRelRes.
type QosFlowAdditionalInfoItemRelRes struct {
	QosFlowIdentifier    QosFlowIdentifier
	HasUplinkTLContainer bool
	UplinkTLContainer    TLContainer
	HasIEExtensions      bool
	IEExtensions         []QosFlowAdditionalInfoItemRelResExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoItemRelRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoItemRelRes", v)
}

// DecodeQosFlowAdditionalInfoItemRelRes decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoItemRelRes(data []byte) (QosFlowAdditionalInfoItemRelRes, int, error) {
	var v QosFlowAdditionalInfoItemRelRes
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoItemRelRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoItemRelRes{}, 0, err
	}
	return v, len(data), err
}
