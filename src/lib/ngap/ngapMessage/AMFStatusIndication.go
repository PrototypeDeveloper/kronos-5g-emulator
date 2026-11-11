// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AMFStatusIndication represents ASN.1 AMFStatusIndication.
type AMFStatusIndication struct {
	UnavailableGUAMIList ngapType.UnavailableGUAMIList `ngap:"id:120,criticality:reject,presence:mandatory"`
	AdditionalIEs        []ngapType.UnknownIE
	UnknownExtensions    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AMFStatusIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("AMFStatusIndication", m)
}

// DecodeAMFStatusIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAMFStatusIndication(data []byte) (*AMFStatusIndication, error) {
	var m AMFStatusIndication
	err := ngapcodec.Decode("AMFStatusIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
