// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ResourceStatusReportingSystem represents ASN.1 ResourceStatusReportingSystem.
type ResourceStatusReportingSystemPresent uint32

const (
	ResourceStatusReportingSystemPresentNothing ResourceStatusReportingSystemPresent = iota
	ResourceStatusReportingSystemPresentEUTRANReportingStatus
	ResourceStatusReportingSystemPresentNGRANReportingStatus
	ResourceStatusReportingSystemPresentChoiceExtensions
)

type ResourceStatusReportingSystem struct {
	Present               ResourceStatusReportingSystemPresent
	EUTRANReportingStatus EUTRANReportingStatusIEs
	NGRANReportingStatus  NGRANReportingStatusIEs
	ChoiceExtensions      ResourceStatusReportingSystemExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ResourceStatusReportingSystem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ResourceStatusReportingSystem", v)
}

// DecodeResourceStatusReportingSystem decodes a standalone APER value and reports the consumed octets.
func DecodeResourceStatusReportingSystem(data []byte) (ResourceStatusReportingSystem, int, error) {
	var v ResourceStatusReportingSystem
	err := ngapcodec.DecodeParameter("ResourceStatusReportingSystem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ResourceStatusReportingSystem{}, 0, err
	}
	return v, len(data), err
}
