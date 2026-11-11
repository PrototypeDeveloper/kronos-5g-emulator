// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExpectedUEMovingTrajectoryItem represents ASN.1 ExpectedUEMovingTrajectoryItem.
type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI            NGRANCGI
	HasTimeStayedInCell bool
	TimeStayedInCell    int64 `aper:"valueLB:0,valueUB:4095"`
	HasIEExtensions     bool
	IEExtensions        []ExpectedUEMovingTrajectoryItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEMovingTrajectoryItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEMovingTrajectoryItem", v)
}

// DecodeExpectedUEMovingTrajectoryItem decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEMovingTrajectoryItem(data []byte) (ExpectedUEMovingTrajectoryItem, int, error) {
	var v ExpectedUEMovingTrajectoryItem
	err := ngapcodec.DecodeParameter("ExpectedUEMovingTrajectoryItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEMovingTrajectoryItem{}, 0, err
	}
	return v, len(data), err
}
