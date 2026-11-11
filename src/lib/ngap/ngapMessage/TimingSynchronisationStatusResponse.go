// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// TimingSynchronisationStatusResponse represents ASN.1 TimingSynchronisationStatusResponse.
type TimingSynchronisationStatusResponse struct {
	RoutingID                 ngapType.RoutingID `ngap:"id:89,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *TimingSynchronisationStatusResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("TimingSynchronisationStatusResponse", m)
}

// DecodeTimingSynchronisationStatusResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeTimingSynchronisationStatusResponse(data []byte) (*TimingSynchronisationStatusResponse, error) {
	var m TimingSynchronisationStatusResponse
	err := ngapcodec.Decode("TimingSynchronisationStatusResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
