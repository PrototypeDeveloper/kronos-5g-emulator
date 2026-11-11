// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEDifferentiationInfoTrafficProfile represents ASN.1 UEDifferentiationInfoTrafficProfile.
const (
	UEDifferentiationInfoTrafficProfilePresentSinglePacket    Enumerated = 0
	UEDifferentiationInfoTrafficProfilePresentDualPackets     Enumerated = 1
	UEDifferentiationInfoTrafficProfilePresentMultiplePackets Enumerated = 2
)

type UEDifferentiationInfoTrafficProfile struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value.
func (v UEDifferentiationInfoTrafficProfile) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:UEDifferentiationInfoTrafficProfile", v)
}

// DecodeUEDifferentiationInfoTrafficProfile decodes a standalone APER value and reports consumed octets.
func DecodeUEDifferentiationInfoTrafficProfile(data []byte) (UEDifferentiationInfoTrafficProfile, int, error) {
	var v UEDifferentiationInfoTrafficProfile
	err := ngapcodec.DecodeParameter("inline:UEDifferentiationInfoTrafficProfile", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfoTrafficProfile{}, 0, err
	}
	return v, len(data), err
}
