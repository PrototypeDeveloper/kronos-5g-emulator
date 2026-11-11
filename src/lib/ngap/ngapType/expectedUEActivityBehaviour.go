// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedUEActivityBehaviour represents ASN.1 ExpectedUEActivityBehaviour.
type ExpectedUEActivityBehaviour struct {
	HasExpectedActivityPeriod                 bool
	ExpectedActivityPeriod                    ExpectedActivityPeriod
	HasExpectedIdlePeriod                     bool
	ExpectedIdlePeriod                        ExpectedIdlePeriod
	HasSourceOfUEActivityBehaviourInformation bool
	SourceOfUEActivityBehaviourInformation    SourceOfUEActivityBehaviourInformation
	HasIEExtensions                           bool
	IEExtensions                              []ExpectedUEActivityBehaviourExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEActivityBehaviour) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEActivityBehaviour", v)
}

// DecodeExpectedUEActivityBehaviour decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEActivityBehaviour(data []byte) (ExpectedUEActivityBehaviour, int, error) {
	var v ExpectedUEActivityBehaviour
	err := ngapcodec.DecodeParameter("ExpectedUEActivityBehaviour", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEActivityBehaviour{}, 0, err
	}
	return v, len(data), err
}
