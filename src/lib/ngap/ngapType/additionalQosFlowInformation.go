// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalQosFlowInformation represents ASN.1 AdditionalQosFlowInformation.
const (
	AdditionalQosFlowInformationPresentMoreLikely Enumerated = 0
)

type AdditionalQosFlowInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalQosFlowInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalQosFlowInformation", v)
}

// DecodeAdditionalQosFlowInformation decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalQosFlowInformation(data []byte) (AdditionalQosFlowInformation, int, error) {
	var v AdditionalQosFlowInformation
	err := ngapcodec.DecodeParameter("AdditionalQosFlowInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalQosFlowInformation{}, 0, err
	}
	return v, len(data), err
}
