// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGSetupResponse represents ASN.1 NGSetupResponse.
type NGSetupResponse struct {
	AMFName                   ngapType.AMFName             `ngap:"id:1,criticality:reject,presence:mandatory"`
	ServedGUAMIList           ngapType.ServedGUAMIList     `ngap:"id:96,criticality:reject,presence:mandatory"`
	RelativeAMFCapacity       ngapType.RelativeAMFCapacity `ngap:"id:86,criticality:ignore,presence:mandatory"`
	PLMNSupportList           ngapType.PLMNSupportList     `ngap:"id:80,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasUERetentionInformation bool
	UERetentionInformation    ngapType.UERetentionInformation `ngap:"id:147,criticality:ignore,presence:optional"`
	HasIABSupported           bool
	IABSupported              ngapType.IABSupported `ngap:"id:200,criticality:ignore,presence:optional"`
	HasExtendedAMFName        bool
	ExtendedAMFName           ngapType.ExtendedAMFName `ngap:"id:274,criticality:ignore,presence:optional"`
	HasMobileIABSupported     bool
	MobileIABSupported        ngapType.MobileIABSupported `ngap:"id:404,criticality:ignore,presence:optional"`
	HasAIOTFIdentifier        bool
	AIOTFIdentifier           ngapType.AIOTFIdentifier `ngap:"id:467,criticality:reject,presence:optional"`
	HasAIOTFName              bool
	AIOTFName                 ngapType.AIOTFName `ngap:"id:476,criticality:reject,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGSetupResponse) Encode() ([]byte, error) { return ngapcodec.Encode("NGSetupResponse", m) }

// DecodeNGSetupResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGSetupResponse(data []byte) (*NGSetupResponse, error) {
	var m NGSetupResponse
	err := ngapcodec.Decode("NGSetupResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
