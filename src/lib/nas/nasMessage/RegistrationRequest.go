package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// RegistrationRequest contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type RegistrationRequest struct {
	RegistrationTypeAndNgKSI               uint8
	MobileIdentity                         []byte
	HasNonCurrentNativeNASKeySetIdentifier bool
	NonCurrentNativeNASKeySetIdentifier    nasType.NonCurrentNativeNASKeySetIdentifier
	HasCapability5GMM                      bool
	Capability5GMM                         nasType.Capability5GMM
	HasUESecurityCapability                bool
	UESecurityCapability                   nasType.UESecurityCapability
	HasRequestedNSSAI                      bool
	RequestedNSSAI                         nasType.RequestedNSSAI
	HasLastVisitedRegisteredTAI            bool
	LastVisitedRegisteredTAI               nasType.LastVisitedRegisteredTAI
	HasS1UENetworkCapability               bool
	S1UENetworkCapability                  nasType.S1UENetworkCapability
	HasUplinkDataStatus                    bool
	UplinkDataStatus                       nasType.UplinkDataStatus
	HasPDUSessionStatus                    bool
	PDUSessionStatus                       nasType.PDUSessionStatus
	HasMICOIndication                      bool
	MICOIndication                         nasType.MICOIndication
	HasUEStatus                            bool
	UEStatus                               nasType.UEStatus
	HasAdditionalGUTI                      bool
	AdditionalGUTI                         nasType.AdditionalGUTI
	HasAllowedPDUSessionStatus             bool
	AllowedPDUSessionStatus                nasType.AllowedPDUSessionStatus
	HasUEUsageSetting                      bool
	UEUsageSetting                         nasType.UEUsageSetting
	HasRequestedDRXParameters              bool
	RequestedDRXParameters                 nasType.RequestedDRXParameters
	HasEPSNASMessageContainer              bool
	EPSNASMessageContainer                 nasType.EPSNASMessageContainer
	HasLADNIndication                      bool
	LADNIndication                         nasType.LADNIndication
	HasPayloadContainerType                bool
	PayloadContainerType                   nasType.PayloadContainerType
	HasPayloadContainer                    bool
	PayloadContainer                       nasType.PayloadContainer
	HasNetworkSlicingIndication            bool
	NetworkSlicingIndication               nasType.NetworkSlicingIndication
	HasUpdateType5GS                       bool
	UpdateType5GS                          nasType.UpdateType5GS
	HasNASMessageContainer                 bool
	NASMessageContainer                    nasType.NASMessageContainer
}

var registrationRequestOptions = nascodec.Schema{
	nasType.IENonCurrentNativeNASKeySetIdentifier: {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IECapability5GMM:                      {Format: nascodec.TLV, Min: 1, Max: 13},
	nasType.IEUESecurityCapability:                {Format: nascodec.TLV, Min: 2, Max: 8},
	nasType.IERequestedNSSAI:                      {Format: nascodec.TLV, Min: 2, Max: 72},
	nasType.IELastVisitedRegisteredTAI:            {Format: nascodec.TV, Min: 6, Max: 6},
	nasType.IES1UENetworkCapability:               {Format: nascodec.TLV, Min: 2, Max: 13},
	nasType.IEUplinkDataStatus:                    {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEPDUSessionStatus:                    {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEMICOIndication:                      {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEUEStatus:                            {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEAdditionalGUTI:                      {Format: nascodec.TLVE, Min: 11, Max: 11},
	nasType.IEAllowedPDUSessionStatus:             {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEUEUsageSetting:                      {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IERequestedDRXParameters:              {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEEPSNASMessageContainer:              {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IELADNIndication:                      {Format: nascodec.TLVE, Min: 0, Max: 808},
	nasType.IEPayloadContainerType:                {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEPayloadContainer:                    {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IENetworkSlicingIndication:            {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEUpdateType5GS:                       {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IENASMessageContainer:                 {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *RegistrationRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil RegistrationRequest")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgRegistrationRequest}
	w.U8(m.RegistrationTypeAndNgKSI, 255)
	w.LV(m.MobileIdentity, 2, 4, 65535)
	var options []nasType.IE
	if m.HasNonCurrentNativeNASKeySetIdentifier {
		options = append(options, nasType.IE{ID: nasType.IENonCurrentNativeNASKeySetIdentifier, Value: []byte{m.NonCurrentNativeNASKeySetIdentifier.Value}})
	}
	if m.HasCapability5GMM {
		options = append(options, nasType.IE{ID: nasType.IECapability5GMM, Value: m.Capability5GMM.Value})
	}
	if m.HasUESecurityCapability {
		options = append(options, nasType.IE{ID: nasType.IEUESecurityCapability, Value: m.UESecurityCapability.Value})
	}
	if m.HasRequestedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IERequestedNSSAI, Value: m.RequestedNSSAI.Value})
	}
	if m.HasLastVisitedRegisteredTAI {
		options = append(options, nasType.IE{ID: nasType.IELastVisitedRegisteredTAI, Value: m.LastVisitedRegisteredTAI.Value})
	}
	if m.HasS1UENetworkCapability {
		options = append(options, nasType.IE{ID: nasType.IES1UENetworkCapability, Value: m.S1UENetworkCapability.Value})
	}
	if m.HasUplinkDataStatus {
		options = append(options, nasType.IE{ID: nasType.IEUplinkDataStatus, Value: m.UplinkDataStatus.Value})
	}
	if m.HasPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionStatus, Value: m.PDUSessionStatus.Value})
	}
	if m.HasMICOIndication {
		options = append(options, nasType.IE{ID: nasType.IEMICOIndication, Value: []byte{m.MICOIndication.Value}})
	}
	if m.HasUEStatus {
		options = append(options, nasType.IE{ID: nasType.IEUEStatus, Value: []byte{m.UEStatus.Value}})
	}
	if m.HasAdditionalGUTI {
		options = append(options, nasType.IE{ID: nasType.IEAdditionalGUTI, Value: m.AdditionalGUTI.Value})
	}
	if m.HasAllowedPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEAllowedPDUSessionStatus, Value: m.AllowedPDUSessionStatus.Value})
	}
	if m.HasUEUsageSetting {
		options = append(options, nasType.IE{ID: nasType.IEUEUsageSetting, Value: []byte{m.UEUsageSetting.Value}})
	}
	if m.HasRequestedDRXParameters {
		options = append(options, nasType.IE{ID: nasType.IERequestedDRXParameters, Value: []byte{m.RequestedDRXParameters.Value}})
	}
	if m.HasEPSNASMessageContainer {
		options = append(options, nasType.IE{ID: nasType.IEEPSNASMessageContainer, Value: m.EPSNASMessageContainer.Value})
	}
	if m.HasLADNIndication {
		options = append(options, nasType.IE{ID: nasType.IELADNIndication, Value: m.LADNIndication.Value})
	}
	if m.HasPayloadContainerType {
		options = append(options, nasType.IE{ID: nasType.IEPayloadContainerType, Value: []byte{m.PayloadContainerType.Value}})
	}
	if m.HasPayloadContainer {
		options = append(options, nasType.IE{ID: nasType.IEPayloadContainer, Value: m.PayloadContainer.Value})
	}
	if m.HasNetworkSlicingIndication {
		options = append(options, nasType.IE{ID: nasType.IENetworkSlicingIndication, Value: []byte{m.NetworkSlicingIndication.Value}})
	}
	if m.HasUpdateType5GS {
		options = append(options, nasType.IE{ID: nasType.IEUpdateType5GS, Value: []byte{m.UpdateType5GS.Value}})
	}
	if m.HasNASMessageContainer {
		options = append(options, nasType.IE{ID: nasType.IENASMessageContainer, Value: m.NASMessageContainer.Value})
	}
	w.Options(options, registrationRequestOptions)
	return w.Result()
}
func DecodeRegistrationRequest(data []byte) (*RegistrationRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgRegistrationRequest {
		return nil, fmt.Errorf("unexpected header for RegistrationRequest")
	}
	m := &RegistrationRequest{}
	r := nascodec.Reader{Data: data[3:]}
	m.RegistrationTypeAndNgKSI = r.U8(255)
	m.MobileIdentity = r.LV(2, 4, 65535)
	for _, ie := range r.Options(registrationRequestOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IENonCurrentNativeNASKeySetIdentifier:
			m.HasNonCurrentNativeNASKeySetIdentifier = true
			m.NonCurrentNativeNASKeySetIdentifier = nasType.NonCurrentNativeNASKeySetIdentifier{Value: ie.Value[0]}
		case nasType.IECapability5GMM:
			m.HasCapability5GMM = true
			m.Capability5GMM = nasType.Capability5GMM{Value: ie.Value}
		case nasType.IEUESecurityCapability:
			m.HasUESecurityCapability = true
			m.UESecurityCapability = nasType.UESecurityCapability{Value: ie.Value}
		case nasType.IERequestedNSSAI:
			m.HasRequestedNSSAI = true
			m.RequestedNSSAI = nasType.RequestedNSSAI{Value: ie.Value}
		case nasType.IELastVisitedRegisteredTAI:
			m.HasLastVisitedRegisteredTAI = true
			m.LastVisitedRegisteredTAI = nasType.LastVisitedRegisteredTAI{Value: ie.Value}
		case nasType.IES1UENetworkCapability:
			m.HasS1UENetworkCapability = true
			m.S1UENetworkCapability = nasType.S1UENetworkCapability{Value: ie.Value}
		case nasType.IEUplinkDataStatus:
			m.HasUplinkDataStatus = true
			m.UplinkDataStatus = nasType.UplinkDataStatus{Value: ie.Value}
		case nasType.IEPDUSessionStatus:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = nasType.PDUSessionStatus{Value: ie.Value}
		case nasType.IEMICOIndication:
			m.HasMICOIndication = true
			m.MICOIndication = nasType.MICOIndication{Value: ie.Value[0]}
		case nasType.IEUEStatus:
			m.HasUEStatus = true
			m.UEStatus = nasType.UEStatus{Value: ie.Value[0]}
		case nasType.IEAdditionalGUTI:
			m.HasAdditionalGUTI = true
			m.AdditionalGUTI = nasType.AdditionalGUTI{Value: ie.Value}
		case nasType.IEAllowedPDUSessionStatus:
			m.HasAllowedPDUSessionStatus = true
			m.AllowedPDUSessionStatus = nasType.AllowedPDUSessionStatus{Value: ie.Value}
		case nasType.IEUEUsageSetting:
			m.HasUEUsageSetting = true
			m.UEUsageSetting = nasType.UEUsageSetting{Value: ie.Value[0]}
		case nasType.IERequestedDRXParameters:
			m.HasRequestedDRXParameters = true
			m.RequestedDRXParameters = nasType.RequestedDRXParameters{Value: ie.Value[0]}
		case nasType.IEEPSNASMessageContainer:
			m.HasEPSNASMessageContainer = true
			m.EPSNASMessageContainer = nasType.EPSNASMessageContainer{Value: ie.Value}
		case nasType.IELADNIndication:
			m.HasLADNIndication = true
			m.LADNIndication = nasType.LADNIndication{Value: ie.Value}
		case nasType.IEPayloadContainerType:
			m.HasPayloadContainerType = true
			m.PayloadContainerType = nasType.PayloadContainerType{Value: ie.Value[0]}
		case nasType.IEPayloadContainer:
			m.HasPayloadContainer = true
			m.PayloadContainer = nasType.PayloadContainer{Value: ie.Value}
		case nasType.IENetworkSlicingIndication:
			m.HasNetworkSlicingIndication = true
			m.NetworkSlicingIndication = nasType.NetworkSlicingIndication{Value: ie.Value[0]}
		case nasType.IEUpdateType5GS:
			m.HasUpdateType5GS = true
			m.UpdateType5GS = nasType.UpdateType5GS{Value: ie.Value[0]}
		case nasType.IENASMessageContainer:
			m.HasNASMessageContainer = true
			m.NASMessageContainer = nasType.NASMessageContainer{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode RegistrationRequest: %w", r.Err)
	}
	return m, nil
}
