// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InterSystemHandoverReportType represents ASN.1 InterSystemHandoverReportType.
type InterSystemHandoverReportTypePresent uint32

const (
	InterSystemHandoverReportTypePresentNothing InterSystemHandoverReportTypePresent = iota
	InterSystemHandoverReportTypePresentTooearlyIntersystemHO
	InterSystemHandoverReportTypePresentIntersystemUnnecessaryHO
	InterSystemHandoverReportTypePresentChoiceExtensions
)

type InterSystemHandoverReportType struct {
	Present                  InterSystemHandoverReportTypePresent
	TooearlyIntersystemHO    TooearlyIntersystemHO
	IntersystemUnnecessaryHO IntersystemUnnecessaryHO
	ChoiceExtensions         InterSystemHandoverReportTypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemHandoverReportType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemHandoverReportType", v)
}

// DecodeInterSystemHandoverReportType decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemHandoverReportType(data []byte) (InterSystemHandoverReportType, int, error) {
	var v InterSystemHandoverReportType
	err := ngapcodec.DecodeParameter("InterSystemHandoverReportType", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemHandoverReportType{}, 0, err
	}
	return v, len(data), err
}
