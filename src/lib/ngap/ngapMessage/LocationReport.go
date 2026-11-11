// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// LocationReport represents ASN.1 LocationReport.
type LocationReport struct {
	AMFUENGAPID                       ngapType.AMFUENGAPID             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                       ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	UserLocationInformation           ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:mandatory"`
	HasUEPresenceInAreaOfInterestList bool
	UEPresenceInAreaOfInterestList    ngapType.UEPresenceInAreaOfInterestList `ngap:"id:116,criticality:ignore,presence:optional"`
	LocationReportingRequestType      ngapType.LocationReportingRequestType   `ngap:"id:33,criticality:ignore,presence:mandatory"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *LocationReport) Encode() ([]byte, error) { return ngapcodec.Encode("LocationReport", m) }

// DecodeLocationReport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeLocationReport(data []byte) (*LocationReport, error) {
	var m LocationReport
	err := ngapcodec.Decode("LocationReport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
