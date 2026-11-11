// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InitialContextSetupRequest represents ASN.1 InitialContextSetupRequest.
type InitialContextSetupRequest struct {
	AMFUENGAPID                                    ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                                    ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasOldAMF                                      bool
	OldAMF                                         ngapType.AMFName `ngap:"id:48,criticality:reject,presence:optional"`
	HasUEAggregateMaximumBitRate                   bool
	UEAggregateMaximumBitRate                      ngapType.UEAggregateMaximumBitRate `ngap:"id:110,criticality:reject,presence:conditional"`
	HasCoreNetworkAssistanceInformationForInactive bool
	CoreNetworkAssistanceInformationForInactive    ngapType.CoreNetworkAssistanceInformationForInactive `ngap:"id:18,criticality:ignore,presence:optional"`
	GUAMI                                          ngapType.GUAMI                                       `ngap:"id:28,criticality:reject,presence:mandatory"`
	HasPDUSessionResourceSetupListCxtReq           bool
	PDUSessionResourceSetupListCxtReq              ngapType.PDUSessionResourceSetupListCxtReq `ngap:"id:71,criticality:reject,presence:optional"`
	AllowedNSSAI                                   ngapType.AllowedNSSAI                      `ngap:"id:0,criticality:reject,presence:mandatory"`
	UESecurityCapabilities                         ngapType.UESecurityCapabilities            `ngap:"id:119,criticality:reject,presence:mandatory"`
	SecurityKey                                    ngapType.SecurityKey                       `ngap:"id:94,criticality:reject,presence:mandatory"`
	HasTraceActivation                             bool
	TraceActivation                                ngapType.TraceActivation `ngap:"id:108,criticality:ignore,presence:optional"`
	HasMobilityRestrictionList                     bool
	MobilityRestrictionList                        ngapType.MobilityRestrictionList `ngap:"id:36,criticality:ignore,presence:optional"`
	HasUERadioCapability                           bool
	UERadioCapability                              ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:optional"`
	HasIndexToRFSP                                 bool
	IndexToRFSP                                    ngapType.IndexToRFSP `ngap:"id:31,criticality:ignore,presence:optional"`
	HasMaskedIMEISV                                bool
	MaskedIMEISV                                   ngapType.MaskedIMEISV `ngap:"id:34,criticality:ignore,presence:optional"`
	HasNASPDU                                      bool
	NASPDU                                         ngapType.NASPDU `ngap:"id:38,criticality:ignore,presence:optional"`
	HasEmergencyFallbackIndicator                  bool
	EmergencyFallbackIndicator                     ngapType.EmergencyFallbackIndicator `ngap:"id:24,criticality:reject,presence:optional"`
	HasRRCInactiveTransitionReportRequest          bool
	RRCInactiveTransitionReportRequest             ngapType.RRCInactiveTransitionReportRequest `ngap:"id:91,criticality:ignore,presence:optional"`
	HasUERadioCapabilityForPaging                  bool
	UERadioCapabilityForPaging                     ngapType.UERadioCapabilityForPaging `ngap:"id:118,criticality:ignore,presence:optional"`
	HasRedirectionVoiceFallback                    bool
	RedirectionVoiceFallback                       ngapType.RedirectionVoiceFallback `ngap:"id:146,criticality:ignore,presence:optional"`
	HasLocationReportingRequestType                bool
	LocationReportingRequestType                   ngapType.LocationReportingRequestType `ngap:"id:33,criticality:ignore,presence:optional"`
	HasCNAssistedRANTuning                         bool
	CNAssistedRANTuning                            ngapType.CNAssistedRANTuning `ngap:"id:165,criticality:ignore,presence:optional"`
	HasSRVCCOperationPossible                      bool
	SRVCCOperationPossible                         ngapType.SRVCCOperationPossible `ngap:"id:177,criticality:ignore,presence:optional"`
	HasIABAuthorized                               bool
	IABAuthorized                                  ngapType.IABAuthorized `ngap:"id:199,criticality:ignore,presence:optional"`
	HasEnhancedCoverageRestriction                 bool
	EnhancedCoverageRestriction                    ngapType.EnhancedCoverageRestriction `ngap:"id:205,criticality:ignore,presence:optional"`
	HasExtendedConnectedTime                       bool
	ExtendedConnectedTime                          ngapType.ExtendedConnectedTime `ngap:"id:206,criticality:ignore,presence:optional"`
	HasUEDifferentiationInfo                       bool
	UEDifferentiationInfo                          ngapType.UEDifferentiationInfo `ngap:"id:209,criticality:ignore,presence:optional"`
	HasNRV2XServicesAuthorized                     bool
	NRV2XServicesAuthorized                        ngapType.NRV2XServicesAuthorized `ngap:"id:216,criticality:ignore,presence:optional"`
	HasLTEV2XServicesAuthorized                    bool
	LTEV2XServicesAuthorized                       ngapType.LTEV2XServicesAuthorized `ngap:"id:215,criticality:ignore,presence:optional"`
	HasNRUESidelinkAggregateMaximumBitrate         bool
	NRUESidelinkAggregateMaximumBitrate            ngapType.NRUESidelinkAggregateMaximumBitrate `ngap:"id:218,criticality:ignore,presence:optional"`
	HasLTEUESidelinkAggregateMaximumBitrate        bool
	LTEUESidelinkAggregateMaximumBitrate           ngapType.LTEUESidelinkAggregateMaximumBitrate `ngap:"id:217,criticality:ignore,presence:optional"`
	HasPC5QoSParameters                            bool
	PC5QoSParameters                               ngapType.PC5QoSParameters `ngap:"id:219,criticality:ignore,presence:optional"`
	HasCEmodeBrestricted                           bool
	CEmodeBrestricted                              ngapType.CEmodeBrestricted `ngap:"id:222,criticality:ignore,presence:optional"`
	HasUEUPCIoTSupport                             bool
	UEUPCIoTSupport                                ngapType.UEUPCIoTSupport `ngap:"id:234,criticality:ignore,presence:optional"`
	HasRGLevelWirelineAccessCharacteristics        bool
	RGLevelWirelineAccessCharacteristics           ngapType.RGLevelWirelineAccessCharacteristics `ngap:"id:238,criticality:ignore,presence:optional"`
	HasManagementBasedMDTPLMNList                  bool
	ManagementBasedMDTPLMNList                     ngapType.MDTPLMNList `ngap:"id:254,criticality:ignore,presence:optional"`
	HasUERadioCapabilityID                         bool
	UERadioCapabilityID                            ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:optional"`
	HasTimeSyncAssistanceInfo                      bool
	TimeSyncAssistanceInfo                         ngapType.TimeSyncAssistanceInfo `ngap:"id:326,criticality:ignore,presence:optional"`
	HasQMCConfigInfo                               bool
	QMCConfigInfo                                  ngapType.QMCConfigInfo `ngap:"id:328,criticality:ignore,presence:optional"`
	HasTargetNSSAIInformation                      bool
	TargetNSSAIInformation                         ngapType.TargetNSSAIInformation `ngap:"id:334,criticality:ignore,presence:optional"`
	HasUESliceMaximumBitRateList                   bool
	UESliceMaximumBitRateList                      ngapType.UESliceMaximumBitRateList `ngap:"id:335,criticality:ignore,presence:optional"`
	HasFiveGProSeAuthorized                        bool
	FiveGProSeAuthorized                           ngapType.FiveGProSeAuthorized `ngap:"id:345,criticality:ignore,presence:optional"`
	HasFiveGProSeUEPC5AggregateMaximumBitRate      bool
	FiveGProSeUEPC5AggregateMaximumBitRate         ngapType.NRUESidelinkAggregateMaximumBitrate `ngap:"id:346,criticality:ignore,presence:optional"`
	HasFiveGProSePC5QoSParameters                  bool
	FiveGProSePC5QoSParameters                     ngapType.FiveGProSePC5QoSParameters `ngap:"id:347,criticality:ignore,presence:optional"`
	HasNetworkControlledRepeaterAuthorized         bool
	NetworkControlledRepeaterAuthorized            ngapType.NetworkControlledRepeaterAuthorized `ngap:"id:367,criticality:ignore,presence:optional"`
	HasAerialUEsubscriptionInformation             bool
	AerialUEsubscriptionInformation                ngapType.AerialUEsubscriptionInformation `ngap:"id:373,criticality:ignore,presence:optional"`
	HasNRA2XServicesAuthorized                     bool
	NRA2XServicesAuthorized                        ngapType.NRA2XServicesAuthorized `ngap:"id:374,criticality:ignore,presence:optional"`
	HasLTEA2XServicesAuthorized                    bool
	LTEA2XServicesAuthorized                       ngapType.LTEA2XServicesAuthorized `ngap:"id:375,criticality:ignore,presence:optional"`
	HasNRA2XUEPC5AggregateMaximumBitRate           bool
	NRA2XUEPC5AggregateMaximumBitRate              ngapType.NRUESidelinkAggregateMaximumBitrate `ngap:"id:376,criticality:ignore,presence:optional"`
	HasLTEA2XUEPC5AggregateMaximumBitRate          bool
	LTEA2XUEPC5AggregateMaximumBitRate             ngapType.LTEUESidelinkAggregateMaximumBitrate `ngap:"id:377,criticality:ignore,presence:optional"`
	HasA2XPC5QoSParameters                         bool
	A2XPC5QoSParameters                            ngapType.A2XPC5QoSParameters `ngap:"id:378,criticality:ignore,presence:optional"`
	HasMobileIABAuthorized                         bool
	MobileIABAuthorized                            ngapType.MobileIABAuthorized `ngap:"id:400,criticality:ignore,presence:optional"`
	HasPartiallyAllowedNSSAI                       bool
	PartiallyAllowedNSSAI                          ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	HasSLPositioningRangingServiceInfo             bool
	SLPositioningRangingServiceInfo                ngapType.SLPositioningRangingServiceInfo `ngap:"id:430,criticality:ignore,presence:optional"`
	HasExtendedOldAMF                              bool
	ExtendedOldAMF                                 ngapType.ExtendedAMFName `ngap:"id:443,criticality:ignore,presence:optional"`
	HasAMFUENGAPID2                                bool
	AMFUENGAPID2                                   ngapType.AMFUENGAPID `ngap:"id:450,criticality:ignore,presence:optional"`
	AdditionalIEs                                  []ngapType.UnknownIE
	UnknownExtensions                              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InitialContextSetupRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("InitialContextSetupRequest", m)
}

// DecodeInitialContextSetupRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInitialContextSetupRequest(data []byte) (*InitialContextSetupRequest, error) {
	var m InitialContextSetupRequest
	err := ngapcodec.Decode("InitialContextSetupRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
