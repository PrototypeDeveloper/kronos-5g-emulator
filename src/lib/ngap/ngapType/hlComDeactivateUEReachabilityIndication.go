// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HLComDeactivateUEReachabilityIndication represents ASN.1 HLComDeactivateUEReachabilityIndication.
const (
	HLComDeactivateUEReachabilityIndicationPresentTrue Enumerated = 0
)

type HLComDeactivateUEReachabilityIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v HLComDeactivateUEReachabilityIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:HLComDeactivateUEReachabilityIndication", v)
}

// DecodeHLComDeactivateUEReachabilityIndication decodes a standalone APER value and reports consumed octets.
func DecodeHLComDeactivateUEReachabilityIndication(data []byte) (HLComDeactivateUEReachabilityIndication, int, error) {
	var v HLComDeactivateUEReachabilityIndication
	err := ngapcodec.DecodeParameter("inline:HLComDeactivateUEReachabilityIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return HLComDeactivateUEReachabilityIndication{}, 0, err
	}
	return v, len(data), err
}
