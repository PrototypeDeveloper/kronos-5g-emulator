// Package libngap encodes and decodes NGAP PDUs using aligned PER.
package libngap

import (
	"fmt"
	"libngap/aper"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NotifyError accompanies a decoded message whose unknown IE requires notification.
type NotifyError = aper.NotifyError
type Diagnostic = aper.Diagnostic
type CriticalityError = aper.CriticalityError

// Encode encodes a supported signal pointer as a complete NGAP PDU.
func Encode(value any) ([]byte, error) {
	switch m := value.(type) {
	case *ngapMessage.AIOTSessionReleaseCommand:
		return m.Encode()
	case *ngapMessage.AIOTSessionReleaseComplete:
		return m.Encode()
	case *ngapMessage.AIOTSessionReleaseRequest:
		return m.Encode()
	case *ngapMessage.AMFCPRelocationIndication:
		return m.Encode()
	case *ngapMessage.AMFConfigurationUpdate:
		return m.Encode()
	case *ngapMessage.AMFConfigurationUpdateAcknowledge:
		return m.Encode()
	case *ngapMessage.AMFConfigurationUpdateFailure:
		return m.Encode()
	case *ngapMessage.AMFStatusIndication:
		return m.Encode()
	case *ngapMessage.BroadcastSessionModificationRequest:
		return m.Encode()
	case *ngapMessage.BroadcastSessionModificationResponse:
		return m.Encode()
	case *ngapMessage.BroadcastSessionModificationFailure:
		return m.Encode()
	case *ngapMessage.BroadcastSessionReleaseRequest:
		return m.Encode()
	case *ngapMessage.BroadcastSessionReleaseResponse:
		return m.Encode()
	case *ngapMessage.BroadcastSessionReleaseRequired:
		return m.Encode()
	case *ngapMessage.BroadcastSessionSetupRequest:
		return m.Encode()
	case *ngapMessage.BroadcastSessionSetupResponse:
		return m.Encode()
	case *ngapMessage.BroadcastSessionSetupFailure:
		return m.Encode()
	case *ngapMessage.BroadcastSessionTransportRequest:
		return m.Encode()
	case *ngapMessage.BroadcastSessionTransportResponse:
		return m.Encode()
	case *ngapMessage.BroadcastSessionTransportFailure:
		return m.Encode()
	case *ngapMessage.CellTrafficTrace:
		return m.Encode()
	case *ngapMessage.CommandRequest:
		return m.Encode()
	case *ngapMessage.CommandResponse:
		return m.Encode()
	case *ngapMessage.CommandFailure:
		return m.Encode()
	case *ngapMessage.ConnectionEstablishmentIndication:
		return m.Encode()
	case *ngapMessage.DeactivateTrace:
		return m.Encode()
	case *ngapMessage.DistributionReleaseRequest:
		return m.Encode()
	case *ngapMessage.DistributionReleaseResponse:
		return m.Encode()
	case *ngapMessage.DistributionSetupRequest:
		return m.Encode()
	case *ngapMessage.DistributionSetupResponse:
		return m.Encode()
	case *ngapMessage.DistributionSetupFailure:
		return m.Encode()
	case *ngapMessage.DownlinkNASTransport:
		return m.Encode()
	case *ngapMessage.DownlinkNonUEAssociatedNRPPaTransport:
		return m.Encode()
	case *ngapMessage.DownlinkRANConfigurationTransfer:
		return m.Encode()
	case *ngapMessage.DownlinkRANEarlyStatusTransfer:
		return m.Encode()
	case *ngapMessage.DownlinkRANStatusTransfer:
		return m.Encode()
	case *ngapMessage.DownlinkRIMInformationTransfer:
		return m.Encode()
	case *ngapMessage.DownlinkUEAssociatedNRPPaTransport:
		return m.Encode()
	case *ngapMessage.ErrorIndication:
		return m.Encode()
	case *ngapMessage.HandoverCancel:
		return m.Encode()
	case *ngapMessage.HandoverCancelAcknowledge:
		return m.Encode()
	case *ngapMessage.HandoverNotify:
		return m.Encode()
	case *ngapMessage.HandoverRequired:
		return m.Encode()
	case *ngapMessage.HandoverCommand:
		return m.Encode()
	case *ngapMessage.HandoverPreparationFailure:
		return m.Encode()
	case *ngapMessage.HandoverRequest:
		return m.Encode()
	case *ngapMessage.HandoverRequestAcknowledge:
		return m.Encode()
	case *ngapMessage.HandoverFailure:
		return m.Encode()
	case *ngapMessage.HandoverSuccess:
		return m.Encode()
	case *ngapMessage.InitialContextSetupRequest:
		return m.Encode()
	case *ngapMessage.InitialContextSetupResponse:
		return m.Encode()
	case *ngapMessage.InitialContextSetupFailure:
		return m.Encode()
	case *ngapMessage.InitialUEMessage:
		return m.Encode()
	case *ngapMessage.InventoryReport:
		return m.Encode()
	case *ngapMessage.InventoryRequest:
		return m.Encode()
	case *ngapMessage.InventoryResponse:
		return m.Encode()
	case *ngapMessage.InventoryFailure:
		return m.Encode()
	case *ngapMessage.LocationReport:
		return m.Encode()
	case *ngapMessage.LocationReportingControl:
		return m.Encode()
	case *ngapMessage.LocationReportingFailureIndication:
		return m.Encode()
	case *ngapMessage.MTCommunicationHandlingRequest:
		return m.Encode()
	case *ngapMessage.MTCommunicationHandlingResponse:
		return m.Encode()
	case *ngapMessage.MTCommunicationHandlingFailure:
		return m.Encode()
	case *ngapMessage.MulticastGroupPaging:
		return m.Encode()
	case *ngapMessage.MulticastSessionActivationRequest:
		return m.Encode()
	case *ngapMessage.MulticastSessionActivationResponse:
		return m.Encode()
	case *ngapMessage.MulticastSessionActivationFailure:
		return m.Encode()
	case *ngapMessage.MulticastSessionDeactivationRequest:
		return m.Encode()
	case *ngapMessage.MulticastSessionDeactivationResponse:
		return m.Encode()
	case *ngapMessage.MulticastSessionUpdateRequest:
		return m.Encode()
	case *ngapMessage.MulticastSessionUpdateResponse:
		return m.Encode()
	case *ngapMessage.MulticastSessionUpdateFailure:
		return m.Encode()
	case *ngapMessage.NASNonDeliveryIndication:
		return m.Encode()
	case *ngapMessage.NGRemovalRequest:
		return m.Encode()
	case *ngapMessage.NGRemovalResponse:
		return m.Encode()
	case *ngapMessage.NGRemovalFailure:
		return m.Encode()
	case *ngapMessage.NGReset:
		return m.Encode()
	case *ngapMessage.NGResetAcknowledge:
		return m.Encode()
	case *ngapMessage.NGSetupRequest:
		return m.Encode()
	case *ngapMessage.NGSetupResponse:
		return m.Encode()
	case *ngapMessage.NGSetupFailure:
		return m.Encode()
	case *ngapMessage.OverloadStart:
		return m.Encode()
	case *ngapMessage.OverloadStop:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceModifyRequest:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceModifyResponse:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceModifyIndication:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceModifyConfirm:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceNotify:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceReleaseCommand:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceReleaseResponse:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceSetupRequest:
		return m.Encode()
	case *ngapMessage.PDUSessionResourceSetupResponse:
		return m.Encode()
	case *ngapMessage.PWSCancelRequest:
		return m.Encode()
	case *ngapMessage.PWSCancelResponse:
		return m.Encode()
	case *ngapMessage.PWSFailureIndication:
		return m.Encode()
	case *ngapMessage.PWSRestartIndication:
		return m.Encode()
	case *ngapMessage.Paging:
		return m.Encode()
	case *ngapMessage.PathSwitchRequest:
		return m.Encode()
	case *ngapMessage.PathSwitchRequestAcknowledge:
		return m.Encode()
	case *ngapMessage.PathSwitchRequestFailure:
		return m.Encode()
	case *ngapMessage.PrivateMessage:
		return m.Encode()
	case *ngapMessage.RANCPRelocationIndication:
		return m.Encode()
	case *ngapMessage.RANConfigurationUpdate:
		return m.Encode()
	case *ngapMessage.RANConfigurationUpdateAcknowledge:
		return m.Encode()
	case *ngapMessage.RANConfigurationUpdateFailure:
		return m.Encode()
	case *ngapMessage.RANPagingRequest:
		return m.Encode()
	case *ngapMessage.RRCInactiveTransitionReport:
		return m.Encode()
	case *ngapMessage.RerouteNASRequest:
		return m.Encode()
	case *ngapMessage.RetrieveUEInformation:
		return m.Encode()
	case *ngapMessage.SecondaryRATDataUsageReport:
		return m.Encode()
	case *ngapMessage.TimingSynchronisationStatusRequest:
		return m.Encode()
	case *ngapMessage.TimingSynchronisationStatusResponse:
		return m.Encode()
	case *ngapMessage.TimingSynchronisationStatusFailure:
		return m.Encode()
	case *ngapMessage.TimingSynchronisationStatusReport:
		return m.Encode()
	case *ngapMessage.TraceFailureIndication:
		return m.Encode()
	case *ngapMessage.TraceStart:
		return m.Encode()
	case *ngapMessage.UEContextModificationRequest:
		return m.Encode()
	case *ngapMessage.UEContextModificationResponse:
		return m.Encode()
	case *ngapMessage.UEContextModificationFailure:
		return m.Encode()
	case *ngapMessage.UEContextReleaseCommand:
		return m.Encode()
	case *ngapMessage.UEContextReleaseComplete:
		return m.Encode()
	case *ngapMessage.UEContextReleaseRequest:
		return m.Encode()
	case *ngapMessage.UEContextResumeRequest:
		return m.Encode()
	case *ngapMessage.UEContextResumeResponse:
		return m.Encode()
	case *ngapMessage.UEContextResumeFailure:
		return m.Encode()
	case *ngapMessage.UEContextSuspendRequest:
		return m.Encode()
	case *ngapMessage.UEContextSuspendResponse:
		return m.Encode()
	case *ngapMessage.UEContextSuspendFailure:
		return m.Encode()
	case *ngapMessage.UEInformationTransfer:
		return m.Encode()
	case *ngapMessage.UERadioCapabilityCheckRequest:
		return m.Encode()
	case *ngapMessage.UERadioCapabilityCheckResponse:
		return m.Encode()
	case *ngapMessage.UERadioCapabilityIDMappingRequest:
		return m.Encode()
	case *ngapMessage.UERadioCapabilityIDMappingResponse:
		return m.Encode()
	case *ngapMessage.UERadioCapabilityInfoIndication:
		return m.Encode()
	case *ngapMessage.UETNLABindingReleaseRequest:
		return m.Encode()
	case *ngapMessage.UplinkNASTransport:
		return m.Encode()
	case *ngapMessage.UplinkNonUEAssociatedNRPPaTransport:
		return m.Encode()
	case *ngapMessage.UplinkRANConfigurationTransfer:
		return m.Encode()
	case *ngapMessage.UplinkRANEarlyStatusTransfer:
		return m.Encode()
	case *ngapMessage.UplinkRANStatusTransfer:
		return m.Encode()
	case *ngapMessage.UplinkRIMInformationTransfer:
		return m.Encode()
	case *ngapMessage.UplinkUEAssociatedNRPPaTransport:
		return m.Encode()
	case *ngapMessage.WriteReplaceWarningRequest:
		return m.Encode()
	case *ngapMessage.WriteReplaceWarningResponse:
		return m.Encode()
	default:
		return nil, fmt.Errorf("unsupported NGAP message type: %T", value)
	}
}

// Decode returns a concrete *ngapMessage signal. A NotifyError accompanies a
// usable decoded value; fatal errors return nil. NAS-PDU is never decoded here.
func Decode(data []byte) (any, error) {
	h, err := ngapcodec.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	switch [2]uint64{h.Kind, h.ProcedureCode} {
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdAIOTSessionRelease}:
		m, e := ngapMessage.DecodeAIOTSessionReleaseCommand(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdAIOTSessionRelease}:
		m, e := ngapMessage.DecodeAIOTSessionReleaseComplete(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdAIOTSessionReleaseRequest}:
		m, e := ngapMessage.DecodeAIOTSessionReleaseRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdAMFCPRelocationIndication}:
		m, e := ngapMessage.DecodeAMFCPRelocationIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdAMFConfigurationUpdate}:
		m, e := ngapMessage.DecodeAMFConfigurationUpdate(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdAMFConfigurationUpdate}:
		m, e := ngapMessage.DecodeAMFConfigurationUpdateAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdAMFConfigurationUpdate}:
		m, e := ngapMessage.DecodeAMFConfigurationUpdateFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdAMFStatusIndication}:
		m, e := ngapMessage.DecodeAMFStatusIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdBroadcastSessionModification}:
		m, e := ngapMessage.DecodeBroadcastSessionModificationRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdBroadcastSessionModification}:
		m, e := ngapMessage.DecodeBroadcastSessionModificationResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdBroadcastSessionModification}:
		m, e := ngapMessage.DecodeBroadcastSessionModificationFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdBroadcastSessionRelease}:
		m, e := ngapMessage.DecodeBroadcastSessionReleaseRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdBroadcastSessionRelease}:
		m, e := ngapMessage.DecodeBroadcastSessionReleaseResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdBroadcastSessionReleaseRequired}:
		m, e := ngapMessage.DecodeBroadcastSessionReleaseRequired(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdBroadcastSessionSetup}:
		m, e := ngapMessage.DecodeBroadcastSessionSetupRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdBroadcastSessionSetup}:
		m, e := ngapMessage.DecodeBroadcastSessionSetupResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdBroadcastSessionSetup}:
		m, e := ngapMessage.DecodeBroadcastSessionSetupFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdBroadcastSessionTransport}:
		m, e := ngapMessage.DecodeBroadcastSessionTransportRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdBroadcastSessionTransport}:
		m, e := ngapMessage.DecodeBroadcastSessionTransportResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdBroadcastSessionTransport}:
		m, e := ngapMessage.DecodeBroadcastSessionTransportFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdCellTrafficTrace}:
		m, e := ngapMessage.DecodeCellTrafficTrace(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdCommandRequest}:
		m, e := ngapMessage.DecodeCommandRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdCommandRequest}:
		m, e := ngapMessage.DecodeCommandResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdCommandRequest}:
		m, e := ngapMessage.DecodeCommandFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdConnectionEstablishmentIndication}:
		m, e := ngapMessage.DecodeConnectionEstablishmentIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDeactivateTrace}:
		m, e := ngapMessage.DecodeDeactivateTrace(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDistributionRelease}:
		m, e := ngapMessage.DecodeDistributionReleaseRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdDistributionRelease}:
		m, e := ngapMessage.DecodeDistributionReleaseResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDistributionSetup}:
		m, e := ngapMessage.DecodeDistributionSetupRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdDistributionSetup}:
		m, e := ngapMessage.DecodeDistributionSetupResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdDistributionSetup}:
		m, e := ngapMessage.DecodeDistributionSetupFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkNASTransport}:
		m, e := ngapMessage.DecodeDownlinkNASTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkNonUEAssociatedNRPPaTransport}:
		m, e := ngapMessage.DecodeDownlinkNonUEAssociatedNRPPaTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkRANConfigurationTransfer}:
		m, e := ngapMessage.DecodeDownlinkRANConfigurationTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkRANEarlyStatusTransfer}:
		m, e := ngapMessage.DecodeDownlinkRANEarlyStatusTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkRANStatusTransfer}:
		m, e := ngapMessage.DecodeDownlinkRANStatusTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkRIMInformationTransfer}:
		m, e := ngapMessage.DecodeDownlinkRIMInformationTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdDownlinkUEAssociatedNRPPaTransport}:
		m, e := ngapMessage.DecodeDownlinkUEAssociatedNRPPaTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdErrorIndication}:
		m, e := ngapMessage.DecodeErrorIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdHandoverCancel}:
		m, e := ngapMessage.DecodeHandoverCancel(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdHandoverCancel}:
		m, e := ngapMessage.DecodeHandoverCancelAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdHandoverNotification}:
		m, e := ngapMessage.DecodeHandoverNotify(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdHandoverPreparation}:
		m, e := ngapMessage.DecodeHandoverRequired(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdHandoverPreparation}:
		m, e := ngapMessage.DecodeHandoverCommand(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdHandoverPreparation}:
		m, e := ngapMessage.DecodeHandoverPreparationFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdHandoverResourceAllocation}:
		m, e := ngapMessage.DecodeHandoverRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdHandoverResourceAllocation}:
		m, e := ngapMessage.DecodeHandoverRequestAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdHandoverResourceAllocation}:
		m, e := ngapMessage.DecodeHandoverFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdHandoverSuccess}:
		m, e := ngapMessage.DecodeHandoverSuccess(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdInitialContextSetup}:
		m, e := ngapMessage.DecodeInitialContextSetupRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdInitialContextSetup}:
		m, e := ngapMessage.DecodeInitialContextSetupResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdInitialContextSetup}:
		m, e := ngapMessage.DecodeInitialContextSetupFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdInitialUEMessage}:
		m, e := ngapMessage.DecodeInitialUEMessage(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdInventoryReport}:
		m, e := ngapMessage.DecodeInventoryReport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdInventoryRequest}:
		m, e := ngapMessage.DecodeInventoryRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdInventoryRequest}:
		m, e := ngapMessage.DecodeInventoryResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdInventoryRequest}:
		m, e := ngapMessage.DecodeInventoryFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdLocationReport}:
		m, e := ngapMessage.DecodeLocationReport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdLocationReportingControl}:
		m, e := ngapMessage.DecodeLocationReportingControl(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdLocationReportingFailureIndication}:
		m, e := ngapMessage.DecodeLocationReportingFailureIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdMTCommunicationHandling}:
		m, e := ngapMessage.DecodeMTCommunicationHandlingRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdMTCommunicationHandling}:
		m, e := ngapMessage.DecodeMTCommunicationHandlingResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdMTCommunicationHandling}:
		m, e := ngapMessage.DecodeMTCommunicationHandlingFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdMulticastGroupPaging}:
		m, e := ngapMessage.DecodeMulticastGroupPaging(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdMulticastSessionActivation}:
		m, e := ngapMessage.DecodeMulticastSessionActivationRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdMulticastSessionActivation}:
		m, e := ngapMessage.DecodeMulticastSessionActivationResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdMulticastSessionActivation}:
		m, e := ngapMessage.DecodeMulticastSessionActivationFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdMulticastSessionDeactivation}:
		m, e := ngapMessage.DecodeMulticastSessionDeactivationRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdMulticastSessionDeactivation}:
		m, e := ngapMessage.DecodeMulticastSessionDeactivationResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdMulticastSessionUpdate}:
		m, e := ngapMessage.DecodeMulticastSessionUpdateRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdMulticastSessionUpdate}:
		m, e := ngapMessage.DecodeMulticastSessionUpdateResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdMulticastSessionUpdate}:
		m, e := ngapMessage.DecodeMulticastSessionUpdateFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdNASNonDeliveryIndication}:
		m, e := ngapMessage.DecodeNASNonDeliveryIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdNGRemoval}:
		m, e := ngapMessage.DecodeNGRemovalRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdNGRemoval}:
		m, e := ngapMessage.DecodeNGRemovalResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdNGRemoval}:
		m, e := ngapMessage.DecodeNGRemovalFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdNGReset}:
		m, e := ngapMessage.DecodeNGReset(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdNGReset}:
		m, e := ngapMessage.DecodeNGResetAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdNGSetup}:
		m, e := ngapMessage.DecodeNGSetupRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdNGSetup}:
		m, e := ngapMessage.DecodeNGSetupResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdNGSetup}:
		m, e := ngapMessage.DecodeNGSetupFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdOverloadStart}:
		m, e := ngapMessage.DecodeOverloadStart(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdOverloadStop}:
		m, e := ngapMessage.DecodeOverloadStop(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPDUSessionResourceModify}:
		m, e := ngapMessage.DecodePDUSessionResourceModifyRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPDUSessionResourceModify}:
		m, e := ngapMessage.DecodePDUSessionResourceModifyResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPDUSessionResourceModifyIndication}:
		m, e := ngapMessage.DecodePDUSessionResourceModifyIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPDUSessionResourceModifyIndication}:
		m, e := ngapMessage.DecodePDUSessionResourceModifyConfirm(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPDUSessionResourceNotify}:
		m, e := ngapMessage.DecodePDUSessionResourceNotify(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPDUSessionResourceRelease}:
		m, e := ngapMessage.DecodePDUSessionResourceReleaseCommand(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPDUSessionResourceRelease}:
		m, e := ngapMessage.DecodePDUSessionResourceReleaseResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPDUSessionResourceSetup}:
		m, e := ngapMessage.DecodePDUSessionResourceSetupRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPDUSessionResourceSetup}:
		m, e := ngapMessage.DecodePDUSessionResourceSetupResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPWSCancel}:
		m, e := ngapMessage.DecodePWSCancelRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPWSCancel}:
		m, e := ngapMessage.DecodePWSCancelResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPWSFailureIndication}:
		m, e := ngapMessage.DecodePWSFailureIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPWSRestartIndication}:
		m, e := ngapMessage.DecodePWSRestartIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPaging}:
		m, e := ngapMessage.DecodePaging(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPathSwitchRequest}:
		m, e := ngapMessage.DecodePathSwitchRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdPathSwitchRequest}:
		m, e := ngapMessage.DecodePathSwitchRequestAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdPathSwitchRequest}:
		m, e := ngapMessage.DecodePathSwitchRequestFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdPrivateMessage}:
		m, e := ngapMessage.DecodePrivateMessage(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRANCPRelocationIndication}:
		m, e := ngapMessage.DecodeRANCPRelocationIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRANConfigurationUpdate}:
		m, e := ngapMessage.DecodeRANConfigurationUpdate(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdRANConfigurationUpdate}:
		m, e := ngapMessage.DecodeRANConfigurationUpdateAcknowledge(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdRANConfigurationUpdate}:
		m, e := ngapMessage.DecodeRANConfigurationUpdateFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRANPagingRequest}:
		m, e := ngapMessage.DecodeRANPagingRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRRCInactiveTransitionReport}:
		m, e := ngapMessage.DecodeRRCInactiveTransitionReport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRerouteNASRequest}:
		m, e := ngapMessage.DecodeRerouteNASRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdRetrieveUEInformation}:
		m, e := ngapMessage.DecodeRetrieveUEInformation(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdSecondaryRATDataUsageReport}:
		m, e := ngapMessage.DecodeSecondaryRATDataUsageReport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdTimingSynchronisationStatus}:
		m, e := ngapMessage.DecodeTimingSynchronisationStatusRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdTimingSynchronisationStatus}:
		m, e := ngapMessage.DecodeTimingSynchronisationStatusResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdTimingSynchronisationStatus}:
		m, e := ngapMessage.DecodeTimingSynchronisationStatusFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdTimingSynchronisationStatusReport}:
		m, e := ngapMessage.DecodeTimingSynchronisationStatusReport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdTraceFailureIndication}:
		m, e := ngapMessage.DecodeTraceFailureIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdTraceStart}:
		m, e := ngapMessage.DecodeTraceStart(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEContextModification}:
		m, e := ngapMessage.DecodeUEContextModificationRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUEContextModification}:
		m, e := ngapMessage.DecodeUEContextModificationResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdUEContextModification}:
		m, e := ngapMessage.DecodeUEContextModificationFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEContextRelease}:
		m, e := ngapMessage.DecodeUEContextReleaseCommand(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUEContextRelease}:
		m, e := ngapMessage.DecodeUEContextReleaseComplete(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEContextReleaseRequest}:
		m, e := ngapMessage.DecodeUEContextReleaseRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEContextResume}:
		m, e := ngapMessage.DecodeUEContextResumeRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUEContextResume}:
		m, e := ngapMessage.DecodeUEContextResumeResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdUEContextResume}:
		m, e := ngapMessage.DecodeUEContextResumeFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEContextSuspend}:
		m, e := ngapMessage.DecodeUEContextSuspendRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUEContextSuspend}:
		m, e := ngapMessage.DecodeUEContextSuspendResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUUnsuccessfulOutcome, ngapType.IdUEContextSuspend}:
		m, e := ngapMessage.DecodeUEContextSuspendFailure(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUEInformationTransfer}:
		m, e := ngapMessage.DecodeUEInformationTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUERadioCapabilityCheck}:
		m, e := ngapMessage.DecodeUERadioCapabilityCheckRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUERadioCapabilityCheck}:
		m, e := ngapMessage.DecodeUERadioCapabilityCheckResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUERadioCapabilityIDMapping}:
		m, e := ngapMessage.DecodeUERadioCapabilityIDMappingRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdUERadioCapabilityIDMapping}:
		m, e := ngapMessage.DecodeUERadioCapabilityIDMappingResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUERadioCapabilityInfoIndication}:
		m, e := ngapMessage.DecodeUERadioCapabilityInfoIndication(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUETNLABindingRelease}:
		m, e := ngapMessage.DecodeUETNLABindingReleaseRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkNASTransport}:
		m, e := ngapMessage.DecodeUplinkNASTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkNonUEAssociatedNRPPaTransport}:
		m, e := ngapMessage.DecodeUplinkNonUEAssociatedNRPPaTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkRANConfigurationTransfer}:
		m, e := ngapMessage.DecodeUplinkRANConfigurationTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkRANEarlyStatusTransfer}:
		m, e := ngapMessage.DecodeUplinkRANEarlyStatusTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkRANStatusTransfer}:
		m, e := ngapMessage.DecodeUplinkRANStatusTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkRIMInformationTransfer}:
		m, e := ngapMessage.DecodeUplinkRIMInformationTransfer(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdUplinkUEAssociatedNRPPaTransport}:
		m, e := ngapMessage.DecodeUplinkUEAssociatedNRPPaTransport(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUInitiatingMessage, ngapType.IdWriteReplaceWarning}:
		m, e := ngapMessage.DecodeWriteReplaceWarningRequest(data)
		if m == nil {
			return nil, e
		}
		return m, e
	case [2]uint64{ngapType.PDUSuccessfulOutcome, ngapType.IdWriteReplaceWarning}:
		m, e := ngapMessage.DecodeWriteReplaceWarningResponse(data)
		if m == nil {
			return nil, e
		}
		return m, e
	default:
		return nil, fmt.Errorf("unsupported NGAP PDU kind %d, procedure code %d", h.Kind, h.ProcedureCode)
	}
}
