// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RANCPRelocationIndication represents ASN.1 RANCPRelocationIndication.
type RANCPRelocationIndication struct {
	RANUENGAPID             ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	FiveGSTMSI              ngapType.FiveGSTMSI              `ngap:"id:26,criticality:reject,presence:mandatory"`
	EUTRACGI                ngapType.EUTRACGI                `ngap:"id:25,criticality:ignore,presence:mandatory"`
	TAI                     ngapType.TAI                     `ngap:"id:213,criticality:ignore,presence:mandatory"`
	ULCPSecurityInformation ngapType.ULCPSecurityInformation `ngap:"id:211,criticality:reject,presence:mandatory"`
	AdditionalIEs           []ngapType.UnknownIE
	UnknownExtensions       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RANCPRelocationIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("RANCPRelocationIndication", m)
}

// DecodeRANCPRelocationIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRANCPRelocationIndication(data []byte) (*RANCPRelocationIndication, error) {
	var m RANCPRelocationIndication
	err := ngapcodec.Decode("RANCPRelocationIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
