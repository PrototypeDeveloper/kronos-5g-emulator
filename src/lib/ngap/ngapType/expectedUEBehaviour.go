// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedUEBehaviour represents ASN.1 ExpectedUEBehaviour.
type ExpectedUEBehaviour struct {
	HasExpectedUEActivityBehaviour bool
	ExpectedUEActivityBehaviour    ExpectedUEActivityBehaviour
	HasExpectedHOInterval          bool
	ExpectedHOInterval             ExpectedHOInterval
	HasExpectedUEMobility          bool
	ExpectedUEMobility             ExpectedUEMobility
	HasExpectedUEMovingTrajectory  bool
	ExpectedUEMovingTrajectory     ExpectedUEMovingTrajectory
	HasIEExtensions                bool
	IEExtensions                   []ExpectedUEBehaviourExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEBehaviour) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEBehaviour", v)
}

// DecodeExpectedUEBehaviour decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEBehaviour(data []byte) (ExpectedUEBehaviour, int, error) {
	var v ExpectedUEBehaviour
	err := ngapcodec.DecodeParameter("ExpectedUEBehaviour", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEBehaviour{}, 0, err
	}
	return v, len(data), err
}
