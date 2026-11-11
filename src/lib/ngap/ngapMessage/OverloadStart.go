// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// OverloadStart represents ASN.1 OverloadStart.
type OverloadStart struct {
	HasAMFOverloadResponse               bool
	AMFOverloadResponse                  ngapType.OverloadResponse `ngap:"id:2,criticality:reject,presence:optional"`
	HasAMFTrafficLoadReductionIndication bool
	AMFTrafficLoadReductionIndication    ngapType.TrafficLoadReductionIndication `ngap:"id:9,criticality:ignore,presence:optional"`
	HasOverloadStartNSSAIList            bool
	OverloadStartNSSAIList               ngapType.OverloadStartNSSAIList `ngap:"id:49,criticality:ignore,presence:optional"`
	HasGUAMIList                         bool
	GUAMIList                            ngapType.GUAMIList `ngap:"id:451,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *OverloadStart) Encode() ([]byte, error) { return ngapcodec.Encode("OverloadStart", m) }

// DecodeOverloadStart decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeOverloadStart(data []byte) (*OverloadStart, error) {
	var m OverloadStart
	err := ngapcodec.Decode("OverloadStart", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
