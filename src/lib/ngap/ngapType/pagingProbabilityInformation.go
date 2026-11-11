// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingProbabilityInformation represents ASN.1 PagingProbabilityInformation.
const (
	PagingProbabilityInformationPresentP00  Enumerated = 0
	PagingProbabilityInformationPresentP05  Enumerated = 1
	PagingProbabilityInformationPresentP10  Enumerated = 2
	PagingProbabilityInformationPresentP15  Enumerated = 3
	PagingProbabilityInformationPresentP20  Enumerated = 4
	PagingProbabilityInformationPresentP25  Enumerated = 5
	PagingProbabilityInformationPresentP30  Enumerated = 6
	PagingProbabilityInformationPresentP35  Enumerated = 7
	PagingProbabilityInformationPresentP40  Enumerated = 8
	PagingProbabilityInformationPresentP45  Enumerated = 9
	PagingProbabilityInformationPresentP50  Enumerated = 10
	PagingProbabilityInformationPresentP55  Enumerated = 11
	PagingProbabilityInformationPresentP60  Enumerated = 12
	PagingProbabilityInformationPresentP65  Enumerated = 13
	PagingProbabilityInformationPresentP70  Enumerated = 14
	PagingProbabilityInformationPresentP75  Enumerated = 15
	PagingProbabilityInformationPresentP80  Enumerated = 16
	PagingProbabilityInformationPresentP85  Enumerated = 17
	PagingProbabilityInformationPresentP90  Enumerated = 18
	PagingProbabilityInformationPresentP95  Enumerated = 19
	PagingProbabilityInformationPresentP100 Enumerated = 20
)

type PagingProbabilityInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:20"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingProbabilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingProbabilityInformation", v)
}

// DecodePagingProbabilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodePagingProbabilityInformation(data []byte) (PagingProbabilityInformation, int, error) {
	var v PagingProbabilityInformation
	err := ngapcodec.DecodeParameter("PagingProbabilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingProbabilityInformation{}, 0, err
	}
	return v, len(data), err
}
