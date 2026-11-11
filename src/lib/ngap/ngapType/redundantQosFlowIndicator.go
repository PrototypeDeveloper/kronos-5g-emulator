// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RedundantQosFlowIndicator represents ASN.1 RedundantQosFlowIndicator.
const (
	RedundantQosFlowIndicatorPresentTrue  Enumerated = 0
	RedundantQosFlowIndicatorPresentFalse Enumerated = 1
)

type RedundantQosFlowIndicator struct {
	Value Enumerated `aper:"valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RedundantQosFlowIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RedundantQosFlowIndicator", v)
}

// DecodeRedundantQosFlowIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeRedundantQosFlowIndicator(data []byte) (RedundantQosFlowIndicator, int, error) {
	var v RedundantQosFlowIndicator
	err := ngapcodec.DecodeParameter("RedundantQosFlowIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return RedundantQosFlowIndicator{}, 0, err
	}
	return v, len(data), err
}
