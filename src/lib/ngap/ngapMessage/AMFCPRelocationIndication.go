// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AMFCPRelocationIndication represents ASN.1 AMFCPRelocationIndication.
type AMFCPRelocationIndication struct {
	AMFUENGAPID              ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID              ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasSNSSAI                bool
	SNSSAI                   ngapType.SNSSAI `ngap:"id:148,criticality:ignore,presence:optional"`
	HasAllowedNSSAI          bool
	AllowedNSSAI             ngapType.AllowedNSSAI `ngap:"id:0,criticality:ignore,presence:optional"`
	HasPartiallyAllowedNSSAI bool
	PartiallyAllowedNSSAI    ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	AdditionalIEs            []ngapType.UnknownIE
	UnknownExtensions        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AMFCPRelocationIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("AMFCPRelocationIndication", m)
}

// DecodeAMFCPRelocationIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAMFCPRelocationIndication(data []byte) (*AMFCPRelocationIndication, error) {
	var m AMFCPRelocationIndication
	err := ngapcodec.Decode("AMFCPRelocationIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
