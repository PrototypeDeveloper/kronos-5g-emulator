// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextModificationRequest represents ASN.1 UEContextModificationRequest.
type UEContextModificationRequest struct {
	AMFUENGAPID                                    ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                                    ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasRANPagingPriority                           bool
	RANPagingPriority                              ngapType.RANPagingPriority `ngap:"id:83,criticality:ignore,presence:optional"`
	HasSecurityKey                                 bool
	SecurityKey                                    ngapType.SecurityKey `ngap:"id:94,criticality:reject,presence:optional"`
	HasIndexToRFSP                                 bool
	IndexToRFSP                                    ngapType.IndexToRFSP `ngap:"id:31,criticality:ignore,presence:optional"`
	HasUEAggregateMaximumBitRate                   bool
	UEAggregateMaximumBitRate                      ngapType.UEAggregateMaximumBitRate `ngap:"id:110,criticality:ignore,presence:optional"`
	HasUESecurityCapabilities                      bool
	UESecurityCapabilities                         ngapType.UESecurityCapabilities `ngap:"id:119,criticality:reject,presence:optional"`
	HasCoreNetworkAssistanceInformationForInactive bool
	CoreNetworkAssistanceInformationForInactive    ngapType.CoreNetworkAssistanceInformationForInactive `ngap:"id:18,criticality:ignore,presence:optional"`
	HasEmergencyFallbackIndicator                  bool
	EmergencyFallbackIndicator                     ngapType.EmergencyFallbackIndicator `ngap:"id:24,criticality:reject,presence:optional"`
	HasNewAMFUENGAPID                              bool
	NewAMFUENGAPID                                 ngapType.AMFUENGAPID `ngap:"id:40,criticality:reject,presence:optional"`
	HasRRCInactiveTransitionReportRequest          bool
	RRCInactiveTransitionReportRequest             ngapType.RRCInactiveTransitionReportRequest `ngap:"id:91,criticality:ignore,presence:optional"`
	HasNewGUAMI                                    bool
	NewGUAMI                                       ngapType.GUAMI `ngap:"id:162,criticality:reject,presence:optional"`
	HasCNAssistedRANTuning                         bool
	CNAssistedRANTuning                            ngapType.CNAssistedRANTuning `ngap:"id:165,criticality:ignore,presence:optional"`
	HasSRVCCOperationPossible                      bool
	SRVCCOperationPossible                         ngapType.SRVCCOperationPossible `ngap:"id:177,criticality:ignore,presence:optional"`
	HasIABAuthorized                               bool
	IABAuthorized                                  ngapType.IABAuthorized `ngap:"id:199,criticality:ignore,presence:optional"`
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
	HasUERadioCapabilityID                         bool
	UERadioCapabilityID                            ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:optional"`
	HasRGLevelWirelineAccessCharacteristics        bool
	RGLevelWirelineAccessCharacteristics           ngapType.RGLevelWirelineAccessCharacteristics `ngap:"id:238,criticality:ignore,presence:optional"`
	HasTimeSyncAssistanceInfo                      bool
	TimeSyncAssistanceInfo                         ngapType.TimeSyncAssistanceInfo `ngap:"id:326,criticality:ignore,presence:optional"`
	HasQMCConfigInfo                               bool
	QMCConfigInfo                                  ngapType.QMCConfigInfo `ngap:"id:328,criticality:ignore,presence:optional"`
	HasQMCDeactivation                             bool
	QMCDeactivation                                ngapType.QMCDeactivation `ngap:"id:329,criticality:ignore,presence:optional"`
	HasUESliceMaximumBitRateList                   bool
	UESliceMaximumBitRateList                      ngapType.UESliceMaximumBitRateList `ngap:"id:335,criticality:ignore,presence:optional"`
	HasManagementBasedMDTPLMNModificationList      bool
	ManagementBasedMDTPLMNModificationList         ngapType.MDTPLMNModificationList `ngap:"id:359,criticality:ignore,presence:optional"`
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
	HasSLPositioningRangingServiceInfo             bool
	SLPositioningRangingServiceInfo                ngapType.SLPositioningRangingServiceInfo `ngap:"id:430,criticality:ignore,presence:optional"`
	AdditionalIEs                                  []ngapType.UnknownIE
	UnknownExtensions                              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextModificationRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextModificationRequest", m)
}

// DecodeUEContextModificationRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextModificationRequest(data []byte) (*UEContextModificationRequest, error) {
	var m UEContextModificationRequest
	err := ngapcodec.Decode("UEContextModificationRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
