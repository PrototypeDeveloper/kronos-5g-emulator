// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowLevelQosParameters represents ASN.1 QosFlowLevelQosParameters.
type QosFlowLevelQosParameters struct {
	QosCharacteristics              QosCharacteristics
	AllocationAndRetentionPriority  AllocationAndRetentionPriority
	HasGBRQosInformation            bool
	GBRQosInformation               GBRQosInformation
	HasReflectiveQosAttribute       bool
	ReflectiveQosAttribute          ReflectiveQosAttribute
	HasAdditionalQosFlowInformation bool
	AdditionalQosFlowInformation    AdditionalQosFlowInformation
	HasIEExtensions                 bool
	IEExtensions                    []QosFlowLevelQosParametersExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowLevelQosParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowLevelQosParameters", v)
}

// DecodeQosFlowLevelQosParameters decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowLevelQosParameters(data []byte) (QosFlowLevelQosParameters, int, error) {
	var v QosFlowLevelQosParameters
	err := ngapcodec.DecodeParameter("QosFlowLevelQosParameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowLevelQosParameters{}, 0, err
	}
	return v, len(data), err
}
