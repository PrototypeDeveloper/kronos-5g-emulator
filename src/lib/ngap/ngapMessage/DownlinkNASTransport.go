// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkNASTransport represents ASN.1 DownlinkNASTransport.
type DownlinkNASTransport struct {
	AMFUENGAPID                    ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                    ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasOldAMF                      bool
	OldAMF                         ngapType.AMFName `ngap:"id:48,criticality:reject,presence:optional"`
	HasRANPagingPriority           bool
	RANPagingPriority              ngapType.RANPagingPriority `ngap:"id:83,criticality:ignore,presence:optional"`
	NASPDU                         ngapType.NASPDU            `ngap:"id:38,criticality:reject,presence:mandatory"`
	HasMobilityRestrictionList     bool
	MobilityRestrictionList        ngapType.MobilityRestrictionList `ngap:"id:36,criticality:ignore,presence:optional"`
	HasIndexToRFSP                 bool
	IndexToRFSP                    ngapType.IndexToRFSP `ngap:"id:31,criticality:ignore,presence:optional"`
	HasUEAggregateMaximumBitRate   bool
	UEAggregateMaximumBitRate      ngapType.UEAggregateMaximumBitRate `ngap:"id:110,criticality:ignore,presence:optional"`
	HasAllowedNSSAI                bool
	AllowedNSSAI                   ngapType.AllowedNSSAI `ngap:"id:0,criticality:reject,presence:optional"`
	HasSRVCCOperationPossible      bool
	SRVCCOperationPossible         ngapType.SRVCCOperationPossible `ngap:"id:177,criticality:ignore,presence:optional"`
	HasEnhancedCoverageRestriction bool
	EnhancedCoverageRestriction    ngapType.EnhancedCoverageRestriction `ngap:"id:205,criticality:ignore,presence:optional"`
	HasExtendedConnectedTime       bool
	ExtendedConnectedTime          ngapType.ExtendedConnectedTime `ngap:"id:206,criticality:ignore,presence:optional"`
	HasUEDifferentiationInfo       bool
	UEDifferentiationInfo          ngapType.UEDifferentiationInfo `ngap:"id:209,criticality:ignore,presence:optional"`
	HasCEmodeBrestricted           bool
	CEmodeBrestricted              ngapType.CEmodeBrestricted `ngap:"id:222,criticality:ignore,presence:optional"`
	HasUERadioCapability           bool
	UERadioCapability              ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:optional"`
	HasUECapabilityInfoRequest     bool
	UECapabilityInfoRequest        ngapType.UECapabilityInfoRequest `ngap:"id:228,criticality:ignore,presence:optional"`
	HasEndIndication               bool
	EndIndication                  ngapType.EndIndication `ngap:"id:226,criticality:ignore,presence:optional"`
	HasUERadioCapabilityID         bool
	UERadioCapabilityID            ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:optional"`
	HasTargetNSSAIInformation      bool
	TargetNSSAIInformation         ngapType.TargetNSSAIInformation `ngap:"id:334,criticality:ignore,presence:optional"`
	HasMaskedIMEISV                bool
	MaskedIMEISV                   ngapType.MaskedIMEISV `ngap:"id:34,criticality:ignore,presence:optional"`
	HasPartiallyAllowedNSSAI       bool
	PartiallyAllowedNSSAI          ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	HasMobileIABAuthorized         bool
	MobileIABAuthorized            ngapType.MobileIABAuthorized `ngap:"id:400,criticality:ignore,presence:optional"`
	HasExtendedOldAMF              bool
	ExtendedOldAMF                 ngapType.ExtendedAMFName `ngap:"id:443,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkNASTransport) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkNASTransport", m)
}

// DecodeDownlinkNASTransport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkNASTransport(data []byte) (*DownlinkNASTransport, error) {
	var m DownlinkNASTransport
	err := ngapcodec.Decode("DownlinkNASTransport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
