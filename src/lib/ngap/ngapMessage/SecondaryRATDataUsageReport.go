// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// SecondaryRATDataUsageReport represents ASN.1 SecondaryRATDataUsageReport.
type SecondaryRATDataUsageReport struct {
	AMFUENGAPID                             ngapType.AMFUENGAPID                             `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                             ngapType.RANUENGAPID                             `ngap:"id:85,criticality:ignore,presence:mandatory"`
	PDUSessionResourceSecondaryRATUsageList ngapType.PDUSessionResourceSecondaryRATUsageList `ngap:"id:142,criticality:ignore,presence:mandatory"`
	HasHandoverFlag                         bool
	HandoverFlag                            ngapType.HandoverFlag `ngap:"id:143,criticality:ignore,presence:optional"`
	HasUserLocationInformation              bool
	UserLocationInformation                 ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                           []ngapType.UnknownIE
	UnknownExtensions                       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *SecondaryRATDataUsageReport) Encode() ([]byte, error) {
	return ngapcodec.Encode("SecondaryRATDataUsageReport", m)
}

// DecodeSecondaryRATDataUsageReport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeSecondaryRATDataUsageReport(data []byte) (*SecondaryRATDataUsageReport, error) {
	var m SecondaryRATDataUsageReport
	err := ngapcodec.Decode("SecondaryRATDataUsageReport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
