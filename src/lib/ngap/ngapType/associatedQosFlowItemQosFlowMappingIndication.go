// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedQosFlowItemQosFlowMappingIndication represents ASN.1 AssociatedQosFlowItemQosFlowMappingIndication.
const (
	AssociatedQosFlowItemQosFlowMappingIndicationPresentUl Enumerated = 0
	AssociatedQosFlowItemQosFlowMappingIndicationPresentDl Enumerated = 1
)

type AssociatedQosFlowItemQosFlowMappingIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AssociatedQosFlowItemQosFlowMappingIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AssociatedQosFlowItemQosFlowMappingIndication", v)
}

// DecodeAssociatedQosFlowItemQosFlowMappingIndication decodes a standalone APER value and reports consumed octets.
func DecodeAssociatedQosFlowItemQosFlowMappingIndication(data []byte) (AssociatedQosFlowItemQosFlowMappingIndication, int, error) {
	var v AssociatedQosFlowItemQosFlowMappingIndication
	err := ngapcodec.DecodeParameter("inline:AssociatedQosFlowItemQosFlowMappingIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedQosFlowItemQosFlowMappingIndication{}, 0, err
	}
	return v, len(data), err
}
