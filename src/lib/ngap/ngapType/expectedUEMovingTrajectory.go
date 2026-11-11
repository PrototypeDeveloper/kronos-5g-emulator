// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedUEMovingTrajectory represents ASN.1 ExpectedUEMovingTrajectory.
type ExpectedUEMovingTrajectory struct {
	List []ExpectedUEMovingTrajectoryItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEMovingTrajectory) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEMovingTrajectory", v)
}

// DecodeExpectedUEMovingTrajectory decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEMovingTrajectory(data []byte) (ExpectedUEMovingTrajectory, int, error) {
	var v ExpectedUEMovingTrajectory
	err := ngapcodec.DecodeParameter("ExpectedUEMovingTrajectory", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEMovingTrajectory{}, 0, err
	}
	return v, len(data), err
}
