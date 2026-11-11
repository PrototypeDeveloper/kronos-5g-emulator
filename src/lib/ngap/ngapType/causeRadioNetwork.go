// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CauseRadioNetwork represents ASN.1 CauseRadioNetwork.
const (
	CauseRadioNetworkPresentUnspecified                                              Enumerated = 0
	CauseRadioNetworkPresentTxnrelocoverallExpiry                                    Enumerated = 1
	CauseRadioNetworkPresentSuccessfulHandover                                       Enumerated = 2
	CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason                         Enumerated = 3
	CauseRadioNetworkPresentReleaseDueTo5gcGeneratedReason                           Enumerated = 4
	CauseRadioNetworkPresentHandoverCancelled                                        Enumerated = 5
	CauseRadioNetworkPresentPartialHandover                                          Enumerated = 6
	CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem              Enumerated = 7
	CauseRadioNetworkPresentHoTargetNotAllowed                                       Enumerated = 8
	CauseRadioNetworkPresentTngrelocoverallExpiry                                    Enumerated = 9
	CauseRadioNetworkPresentTngrelocprepExpiry                                       Enumerated = 10
	CauseRadioNetworkPresentCellNotAvailable                                         Enumerated = 11
	CauseRadioNetworkPresentUnknownTargetID                                          Enumerated = 12
	CauseRadioNetworkPresentNoRadioResourcesAvailableInTargetCell                    Enumerated = 13
	CauseRadioNetworkPresentUnknownLocalUENGAPID                                     Enumerated = 14
	CauseRadioNetworkPresentInconsistentRemoteUENGAPID                               Enumerated = 15
	CauseRadioNetworkPresentHandoverDesirableForRadioReason                          Enumerated = 16
	CauseRadioNetworkPresentTimeCriticalHandover                                     Enumerated = 17
	CauseRadioNetworkPresentResourceOptimisationHandover                             Enumerated = 18
	CauseRadioNetworkPresentReduceLoadInServingCell                                  Enumerated = 19
	CauseRadioNetworkPresentUserInactivity                                           Enumerated = 20
	CauseRadioNetworkPresentRadioConnectionWithUeLost                                Enumerated = 21
	CauseRadioNetworkPresentRadioResourcesNotAvailable                               Enumerated = 22
	CauseRadioNetworkPresentInvalidQosCombination                                    Enumerated = 23
	CauseRadioNetworkPresentFailureInRadioInterfaceProcedure                         Enumerated = 24
	CauseRadioNetworkPresentInteractionWithOtherProcedure                            Enumerated = 25
	CauseRadioNetworkPresentUnknownPDUSessionID                                      Enumerated = 26
	CauseRadioNetworkPresentUnkownQosFlowID                                          Enumerated = 27
	CauseRadioNetworkPresentMultiplePDUSessionIDInstances                            Enumerated = 28
	CauseRadioNetworkPresentMultipleQosFlowIDInstances                               Enumerated = 29
	CauseRadioNetworkPresentEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported Enumerated = 30
	CauseRadioNetworkPresentNgIntraSystemHandoverTriggered                           Enumerated = 31
	CauseRadioNetworkPresentNgInterSystemHandoverTriggered                           Enumerated = 32
	CauseRadioNetworkPresentXnHandoverTriggered                                      Enumerated = 33
	CauseRadioNetworkPresentNotSupported5QIValue                                     Enumerated = 34
	CauseRadioNetworkPresentUeContextTransfer                                        Enumerated = 35
	CauseRadioNetworkPresentImsVoiceEpsFallbackOrRatFallbackTriggered                Enumerated = 36
	CauseRadioNetworkPresentUpIntegrityProtectionNotPossible                         Enumerated = 37
	CauseRadioNetworkPresentUpConfidentialityProtectionNotPossible                   Enumerated = 38
	CauseRadioNetworkPresentSliceNotSupported                                        Enumerated = 39
	CauseRadioNetworkPresentUeInRrcInactiveStateNotReachable                         Enumerated = 40
	CauseRadioNetworkPresentRedirection                                              Enumerated = 41
	CauseRadioNetworkPresentResourcesNotAvailableForTheSlice                         Enumerated = 42
	CauseRadioNetworkPresentUeMaxIntegrityProtectedDataRateReason                    Enumerated = 43
	CauseRadioNetworkPresentReleaseDueToCnDetectedMobility                           Enumerated = 44
	CauseRadioNetworkPresentN26InterfaceNotAvailable                                 Enumerated = 45
	CauseRadioNetworkPresentReleaseDueToPreEmption                                   Enumerated = 46
	CauseRadioNetworkPresentMultipleLocationReportingReferenceIDInstances            Enumerated = 47
	CauseRadioNetworkPresentRsnNotAvailableForTheUp                                  Enumerated = 48
	CauseRadioNetworkPresentNpnAccessDenied                                          Enumerated = 49
	CauseRadioNetworkPresentCagOnlyAccessDenied                                      Enumerated = 50
	CauseRadioNetworkPresentInsufficientUeCapabilities                               Enumerated = 51
	CauseRadioNetworkPresentRedcapUeNotSupported                                     Enumerated = 52
	CauseRadioNetworkPresentUnknownMBSSessionID                                      Enumerated = 53
	CauseRadioNetworkPresentIndicatedMBSSessionAreaInformationNotServedByTheGNB      Enumerated = 54
	CauseRadioNetworkPresentInconsistentSliceInfoForTheSession                       Enumerated = 55
	CauseRadioNetworkPresentMisalignedAssociationForMulticastUnicast                 Enumerated = 56
	CauseRadioNetworkPresentEredcapUeNotSupported                                    Enumerated = 57
	CauseRadioNetworkPresentTwoRxXrUeNotSupported                                    Enumerated = 58
	CauseRadioNetworkPresentAerialUeFlightInformationReportingInitiationFailure      Enumerated = 59
	CauseRadioNetworkPresentUnknownRANAIoTDeviceNGAPID                               Enumerated = 60
	CauseRadioNetworkPresentRequestedAIoTServiceAreaInformationNotServedByTheGNB     Enumerated = 61
	CauseRadioNetworkPresentUnknownAIoTSession                                       Enumerated = 62
	CauseRadioNetworkPresentAiotDeviceNotReachable                                   Enumerated = 63
	CauseRadioNetworkPresentMultipleAIoTSessionIDInstances                           Enumerated = 64
)

type CauseRadioNetwork struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:44"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseRadioNetwork) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CauseRadioNetwork", v)
}

// DecodeCauseRadioNetwork decodes a standalone APER value and reports the consumed octets.
func DecodeCauseRadioNetwork(data []byte) (CauseRadioNetwork, int, error) {
	var v CauseRadioNetwork
	err := ngapcodec.DecodeParameter("CauseRadioNetwork", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseRadioNetwork{}, 0, err
	}
	return v, len(data), err
}
