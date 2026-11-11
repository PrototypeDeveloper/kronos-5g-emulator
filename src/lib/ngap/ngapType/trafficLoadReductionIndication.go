// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TrafficLoadReductionIndication represents ASN.1 TrafficLoadReductionIndication.
type TrafficLoadReductionIndication struct {
	Value int64 `aper:"valueLB:1,valueUB:99"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TrafficLoadReductionIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TrafficLoadReductionIndication", v)
}

// DecodeTrafficLoadReductionIndication decodes a standalone APER value and reports the consumed octets.
func DecodeTrafficLoadReductionIndication(data []byte) (TrafficLoadReductionIndication, int, error) {
	var v TrafficLoadReductionIndication
	err := ngapcodec.DecodeParameter("TrafficLoadReductionIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return TrafficLoadReductionIndication{}, 0, err
	}
	return v, len(data), err
}
