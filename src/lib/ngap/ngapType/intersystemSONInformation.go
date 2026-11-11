// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONInformation represents ASN.1 IntersystemSONInformation.
type IntersystemSONInformationPresent uint32

const (
	IntersystemSONInformationPresentNothing IntersystemSONInformationPresent = iota
	IntersystemSONInformationPresentIntersystemSONInformationReport
	IntersystemSONInformationPresentChoiceExtensions
)

type IntersystemSONInformation struct {
	Present                         IntersystemSONInformationPresent
	IntersystemSONInformationReport IntersystemSONInformationReport
	ChoiceExtensions                IntersystemSONInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformation", v)
}

// DecodeIntersystemSONInformation decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformation(data []byte) (IntersystemSONInformation, int, error) {
	var v IntersystemSONInformation
	err := ngapcodec.DecodeParameter("IntersystemSONInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformation{}, 0, err
	}
	return v, len(data), err
}
