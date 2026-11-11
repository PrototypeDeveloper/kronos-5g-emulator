// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RATInformation represents ASN.1 RAT-Information.
const (
	RATInformationPresentUnlicensed Enumerated = 0
	RATInformationPresentNbIoT      Enumerated = 1
	RATInformationPresentNRLEO      Enumerated = 2
	RATInformationPresentNRMEO      Enumerated = 3
	RATInformationPresentNRGEO      Enumerated = 4
	RATInformationPresentNROTHERSAT Enumerated = 5
)

type RATInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RATInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RAT-Information", v)
}

// DecodeRATInformation decodes a standalone APER value and reports the consumed octets.
func DecodeRATInformation(data []byte) (RATInformation, int, error) {
	var v RATInformation
	err := ngapcodec.DecodeParameter("RAT-Information", data, &v)
	if ngapcodec.IsFatal(err) {
		return RATInformation{}, 0, err
	}
	return v, len(data), err
}
