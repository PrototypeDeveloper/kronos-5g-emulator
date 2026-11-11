// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// TimingSynchronisationStatusRequest represents ASN.1 TimingSynchronisationStatusRequest.
type TimingSynchronisationStatusRequest struct {
	RoutingID         ngapType.RoutingID         `ngap:"id:89,criticality:reject,presence:mandatory"`
	RANTSSRequestType ngapType.RANTSSRequestType `ngap:"id:388,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *TimingSynchronisationStatusRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("TimingSynchronisationStatusRequest", m)
}

// DecodeTimingSynchronisationStatusRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeTimingSynchronisationStatusRequest(data []byte) (*TimingSynchronisationStatusRequest, error) {
	var m TimingSynchronisationStatusRequest
	err := ngapcodec.Decode("TimingSynchronisationStatusRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
