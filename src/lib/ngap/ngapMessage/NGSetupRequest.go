// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGSetupRequest represents ASN.1 NGSetupRequest.
type NGSetupRequest struct {
	GlobalRANNodeID           ngapType.GlobalRANNodeID `ngap:"id:27,criticality:reject,presence:mandatory"`
	HasRANNodeName            bool
	RANNodeName               ngapType.RANNodeName     `ngap:"id:82,criticality:ignore,presence:optional"`
	SupportedTAList           ngapType.SupportedTAList `ngap:"id:102,criticality:reject,presence:mandatory"`
	DefaultPagingDRX          ngapType.PagingDRX       `ngap:"id:21,criticality:ignore,presence:mandatory"`
	HasUERetentionInformation bool
	UERetentionInformation    ngapType.UERetentionInformation `ngap:"id:147,criticality:ignore,presence:optional"`
	HasNBIoTDefaultPagingDRX  bool
	NBIoTDefaultPagingDRX     ngapType.NBIoTDefaultPagingDRX `ngap:"id:204,criticality:ignore,presence:optional"`
	HasExtendedRANNodeName    bool
	ExtendedRANNodeName       ngapType.ExtendedRANNodeName `ngap:"id:273,criticality:ignore,presence:optional"`
	HasAIoTSupport            bool
	AIoTSupport               ngapType.AIoTSupport `ngap:"id:475,criticality:reject,presence:optional"`
	HasAdditionalULI          bool
	AdditionalULI             ngapType.AdditionalULI `ngap:"id:483,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGSetupRequest) Encode() ([]byte, error) { return ngapcodec.Encode("NGSetupRequest", m) }

// DecodeNGSetupRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGSetupRequest(data []byte) (*NGSetupRequest, error) {
	var m NGSetupRequest
	err := ngapcodec.Decode("NGSetupRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
