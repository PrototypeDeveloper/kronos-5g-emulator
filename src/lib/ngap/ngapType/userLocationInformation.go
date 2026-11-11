// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformation represents ASN.1 UserLocationInformation.
type UserLocationInformationPresent uint32

const (
	UserLocationInformationPresentNothing UserLocationInformationPresent = iota
	UserLocationInformationPresentUserLocationInformationEUTRA
	UserLocationInformationPresentUserLocationInformationNR
	UserLocationInformationPresentUserLocationInformationN3IWFWithPortNumber
	UserLocationInformationPresentChoiceExtensions
)

type UserLocationInformation struct {
	Present                                    UserLocationInformationPresent
	UserLocationInformationEUTRA               UserLocationInformationEUTRA
	UserLocationInformationNR                  UserLocationInformationNR
	UserLocationInformationN3IWFWithPortNumber UserLocationInformationN3IWFWithPortNumber
	ChoiceExtensions                           UserLocationInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformation", v)
}

// DecodeUserLocationInformation decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformation(data []byte) (UserLocationInformation, int, error) {
	var v UserLocationInformation
	err := ngapcodec.DecodeParameter("UserLocationInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformation{}, 0, err
	}
	return v, len(data), err
}
