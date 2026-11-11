// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MRBProgressInformation represents ASN.1 MRB-ProgressInformation.
type MRBProgressInformationPresent uint32

const (
	MRBProgressInformationPresentNothing MRBProgressInformationPresent = iota
	MRBProgressInformationPresentPDCPSNLength12
	MRBProgressInformationPresentPDCPSNLength18
	MRBProgressInformationPresentChoiceExtensions
)

type MRBProgressInformation struct {
	Present          MRBProgressInformationPresent
	PDCPSNLength12   int64 `aper:"valueLB:0,valueUB:4095"`
	PDCPSNLength18   int64 `aper:"valueLB:0,valueUB:262143"`
	ChoiceExtensions MRBProgressInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MRBProgressInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MRB-ProgressInformation", v)
}

// DecodeMRBProgressInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMRBProgressInformation(data []byte) (MRBProgressInformation, int, error) {
	var v MRBProgressInformation
	err := ngapcodec.DecodeParameter("MRB-ProgressInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MRBProgressInformation{}, 0, err
	}
	return v, len(data), err
}
