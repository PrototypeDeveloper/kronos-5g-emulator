// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTimingSynchronisationStatusInfoSynchronisationState represents ASN.1 RANTimingSynchronisationStatusInfoSynchronisationState.
const (
	RANTimingSynchronisationStatusInfoSynchronisationStatePresentLocked   Enumerated = 0
	RANTimingSynchronisationStatusInfoSynchronisationStatePresentHoldover Enumerated = 1
	RANTimingSynchronisationStatusInfoSynchronisationStatePresentFreerun  Enumerated = 2
)

type RANTimingSynchronisationStatusInfoSynchronisationState struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value.
func (v RANTimingSynchronisationStatusInfoSynchronisationState) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:RANTimingSynchronisationStatusInfoSynchronisationState", v)
}

// DecodeRANTimingSynchronisationStatusInfoSynchronisationState decodes a standalone APER value and reports consumed octets.
func DecodeRANTimingSynchronisationStatusInfoSynchronisationState(data []byte) (RANTimingSynchronisationStatusInfoSynchronisationState, int, error) {
	var v RANTimingSynchronisationStatusInfoSynchronisationState
	err := ngapcodec.DecodeParameter("inline:RANTimingSynchronisationStatusInfoSynchronisationState", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTimingSynchronisationStatusInfoSynchronisationState{}, 0, err
	}
	return v, len(data), err
}
