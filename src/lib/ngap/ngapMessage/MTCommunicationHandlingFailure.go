// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MTCommunicationHandlingFailure represents ASN.1 MTCommunicationHandlingFailure.
type MTCommunicationHandlingFailure struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	Cause                     ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MTCommunicationHandlingFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("MTCommunicationHandlingFailure", m)
}

// DecodeMTCommunicationHandlingFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMTCommunicationHandlingFailure(data []byte) (*MTCommunicationHandlingFailure, error) {
	var m MTCommunicationHandlingFailure
	err := ngapcodec.Decode("MTCommunicationHandlingFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
