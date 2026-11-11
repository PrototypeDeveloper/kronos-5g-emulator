// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceOfUEActivityBehaviourInformation represents ASN.1 SourceOfUEActivityBehaviourInformation.
const (
	SourceOfUEActivityBehaviourInformationPresentSubscriptionInformation Enumerated = 0
	SourceOfUEActivityBehaviourInformationPresentStatistics              Enumerated = 1
)

type SourceOfUEActivityBehaviourInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceOfUEActivityBehaviourInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceOfUEActivityBehaviourInformation", v)
}

// DecodeSourceOfUEActivityBehaviourInformation decodes a standalone APER value and reports the consumed octets.
func DecodeSourceOfUEActivityBehaviourInformation(data []byte) (SourceOfUEActivityBehaviourInformation, int, error) {
	var v SourceOfUEActivityBehaviourInformation
	err := ngapcodec.DecodeParameter("SourceOfUEActivityBehaviourInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceOfUEActivityBehaviourInformation{}, 0, err
	}
	return v, len(data), err
}
