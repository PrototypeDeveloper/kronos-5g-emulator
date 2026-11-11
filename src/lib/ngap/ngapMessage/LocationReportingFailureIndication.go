// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// LocationReportingFailureIndication represents ASN.1 LocationReportingFailureIndication.
type LocationReportingFailureIndication struct {
	AMFUENGAPID                                        ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                                        ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	Cause                                              ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasAerialUEFlightInformationOngoingReportingFailed bool
	AerialUEFlightInformationOngoingReportingFailed    ngapType.AerialUEFlightInformationOngoingReportingFailed `ngap:"id:446,criticality:ignore,presence:optional"`
	AdditionalIEs                                      []ngapType.UnknownIE
	UnknownExtensions                                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *LocationReportingFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("LocationReportingFailureIndication", m)
}

// DecodeLocationReportingFailureIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeLocationReportingFailureIndication(data []byte) (*LocationReportingFailureIndication, error) {
	var m LocationReportingFailureIndication
	err := ngapcodec.Decode("LocationReportingFailureIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
