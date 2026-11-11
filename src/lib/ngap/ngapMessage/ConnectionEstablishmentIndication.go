// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// ConnectionEstablishmentIndication represents ASN.1 ConnectionEstablishmentIndication.
type ConnectionEstablishmentIndication struct {
	AMFUENGAPID                    ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                    ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasUERadioCapability           bool
	UERadioCapability              ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:optional"`
	HasEndIndication               bool
	EndIndication                  ngapType.EndIndication `ngap:"id:226,criticality:ignore,presence:optional"`
	HasSNSSAI                      bool
	SNSSAI                         ngapType.SNSSAI `ngap:"id:148,criticality:ignore,presence:optional"`
	HasAllowedNSSAI                bool
	AllowedNSSAI                   ngapType.AllowedNSSAI `ngap:"id:0,criticality:ignore,presence:optional"`
	HasUEDifferentiationInfo       bool
	UEDifferentiationInfo          ngapType.UEDifferentiationInfo `ngap:"id:209,criticality:ignore,presence:optional"`
	HasDLCPSecurityInformation     bool
	DLCPSecurityInformation        ngapType.DLCPSecurityInformation `ngap:"id:212,criticality:ignore,presence:optional"`
	HasNBIoTUEPriority             bool
	NBIoTUEPriority                ngapType.NBIoTUEPriority `ngap:"id:210,criticality:ignore,presence:optional"`
	HasEnhancedCoverageRestriction bool
	EnhancedCoverageRestriction    ngapType.EnhancedCoverageRestriction `ngap:"id:205,criticality:ignore,presence:optional"`
	HasCEmodeBrestricted           bool
	CEmodeBrestricted              ngapType.CEmodeBrestricted `ngap:"id:222,criticality:ignore,presence:optional"`
	HasUERadioCapabilityID         bool
	UERadioCapabilityID            ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:optional"`
	HasMaskedIMEISV                bool
	MaskedIMEISV                   ngapType.MaskedIMEISV `ngap:"id:34,criticality:ignore,presence:optional"`
	HasOldAMF                      bool
	OldAMF                         ngapType.AMFName `ngap:"id:48,criticality:reject,presence:optional"`
	HasPartiallyAllowedNSSAI       bool
	PartiallyAllowedNSSAI          ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	HasExtendedOldAMF              bool
	ExtendedOldAMF                 ngapType.ExtendedAMFName `ngap:"id:443,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *ConnectionEstablishmentIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("ConnectionEstablishmentIndication", m)
}

// DecodeConnectionEstablishmentIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeConnectionEstablishmentIndication(data []byte) (*ConnectionEstablishmentIndication, error) {
	var m ConnectionEstablishmentIndication
	err := ngapcodec.Decode("ConnectionEstablishmentIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
