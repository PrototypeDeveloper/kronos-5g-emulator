// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// TimingSynchronisationStatusReport represents ASN.1 TimingSynchronisationStatusReport.
type TimingSynchronisationStatusReport struct {
	RoutingID                          ngapType.RoutingID                          `ngap:"id:89,criticality:ignore,presence:mandatory"`
	RANTimingSynchronisationStatusInfo ngapType.RANTimingSynchronisationStatusInfo `ngap:"id:387,criticality:ignore,presence:mandatory"`
	RANTSSScope                        ngapType.RANTSSScope                        `ngap:"id:389,criticality:ignore,presence:mandatory"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *TimingSynchronisationStatusReport) Encode() ([]byte, error) {
	return ngapcodec.Encode("TimingSynchronisationStatusReport", m)
}

// DecodeTimingSynchronisationStatusReport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeTimingSynchronisationStatusReport(data []byte) (*TimingSynchronisationStatusReport, error) {
	var m TimingSynchronisationStatusReport
	err := ngapcodec.Decode("TimingSynchronisationStatusReport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
