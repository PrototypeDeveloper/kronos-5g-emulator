package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ControlPlaneServiceRequest requests a service over the control plane (CIoT).
type ControlPlaneServiceRequest struct {
	ControlPlaneServiceTypeAndNgKSI uint8 // low nibble control plane service type, high nibble ngKSI
	HasCIoTSmallDataContainer       bool
	CIoTSmallDataContainer          nasType.CIoTSmallDataContainer
	HasPayloadContainerType         bool
	PayloadContainerType            nasType.PayloadContainerType
	HasPayloadContainer             bool
	PayloadContainer                nasType.PayloadContainer
	HasPDUSessionID                 bool
	PDUSessionID                    nasType.PDUSessionID
	HasPDUSessionStatus             bool
	PDUSessionStatus                nasType.PDUSessionStatus
	HasReleaseAssistanceIndication  bool
	ReleaseAssistanceIndication     nasType.ReleaseAssistanceIndication
	HasUplinkDataStatus             bool
	UplinkDataStatus                nasType.UplinkDataStatus
	HasNASMessageContainer          bool
	NASMessageContainer             nasType.NASMessageContainer
	HasAdditionalInformation        bool
	AdditionalInformation           nasType.AdditionalInformation
	UnknownIEs                      []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var controlPlaneServiceRequestOptions = nascodec.Schema{
	nasType.IECIoTSmallDataContainer:      {Format: nascodec.TLV, Min: 2, Max: 255},
	nasType.IEPayloadContainerType:        {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEPayloadContainer:            {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEPDUSessionID:                {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEPDUSessionStatus:            {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEReleaseAssistanceIndication: {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEUplinkDataStatus:            {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IENASMessageContainer:         {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEAdditionalInformation:       {Format: nascodec.TLV, Min: 1, Max: 255},
}

func (m *ControlPlaneServiceRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ControlPlaneServiceRequest")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgControlPlaneServiceRequest}
	w.U8(m.ControlPlaneServiceTypeAndNgKSI, 255)
	var options []nasType.IE
	if m.HasCIoTSmallDataContainer {
		options = append(options, nasType.IE{ID: nasType.IECIoTSmallDataContainer, Value: m.CIoTSmallDataContainer.Value})
	}
	if m.HasPayloadContainerType {
		options = append(options, nasType.IE{ID: nasType.IEPayloadContainerType, Value: []byte{m.PayloadContainerType.Value}})
	}
	if m.HasPayloadContainer {
		options = append(options, nasType.IE{ID: nasType.IEPayloadContainer, Value: m.PayloadContainer.Value})
	}
	if m.HasPDUSessionID {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionID, Value: []byte{m.PDUSessionID.Value}})
	}
	if m.HasPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionStatus, Value: m.PDUSessionStatus.Value})
	}
	if m.HasReleaseAssistanceIndication {
		options = append(options, nasType.IE{ID: nasType.IEReleaseAssistanceIndication, Value: []byte{m.ReleaseAssistanceIndication.Value}})
	}
	if m.HasUplinkDataStatus {
		options = append(options, nasType.IE{ID: nasType.IEUplinkDataStatus, Value: m.UplinkDataStatus.Value})
	}
	if m.HasNASMessageContainer {
		options = append(options, nasType.IE{ID: nasType.IENASMessageContainer, Value: m.NASMessageContainer.Value})
	}
	if m.HasAdditionalInformation {
		options = append(options, nasType.IE{ID: nasType.IEAdditionalInformation, Value: m.AdditionalInformation.Value})
	}
	w.Options(options, controlPlaneServiceRequestOptions)
	w.UnknownOptions(m.UnknownIEs, controlPlaneServiceRequestOptions)
	return w.Result()
}
func DecodeControlPlaneServiceRequest(data []byte) (*ControlPlaneServiceRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgControlPlaneServiceRequest {
		return nil, fmt.Errorf("unexpected header for ControlPlaneServiceRequest")
	}
	m := &ControlPlaneServiceRequest{}
	r := nascodec.Reader{Data: data[3:]}
	m.ControlPlaneServiceTypeAndNgKSI = r.U8(255)
	known, unknown := r.OptionsWithUnknown(controlPlaneServiceRequestOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode ControlPlaneServiceRequest: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IECIoTSmallDataContainer:
			m.HasCIoTSmallDataContainer = true
			m.CIoTSmallDataContainer = nasType.CIoTSmallDataContainer{Value: ie.Value}
		case nasType.IEPayloadContainerType:
			m.HasPayloadContainerType = true
			m.PayloadContainerType = nasType.PayloadContainerType{Value: ie.Value[0]}
		case nasType.IEPayloadContainer:
			m.HasPayloadContainer = true
			m.PayloadContainer = nasType.PayloadContainer{Value: ie.Value}
		case nasType.IEPDUSessionID:
			m.HasPDUSessionID = true
			m.PDUSessionID = nasType.PDUSessionID{Value: ie.Value[0]}
		case nasType.IEPDUSessionStatus:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = nasType.PDUSessionStatus{Value: ie.Value}
		case nasType.IEReleaseAssistanceIndication:
			m.HasReleaseAssistanceIndication = true
			m.ReleaseAssistanceIndication = nasType.ReleaseAssistanceIndication{Value: ie.Value[0]}
		case nasType.IEUplinkDataStatus:
			m.HasUplinkDataStatus = true
			m.UplinkDataStatus = nasType.UplinkDataStatus{Value: ie.Value}
		case nasType.IENASMessageContainer:
			m.HasNASMessageContainer = true
			m.NASMessageContainer = nasType.NASMessageContainer{Value: ie.Value}
		case nasType.IEAdditionalInformation:
			m.HasAdditionalInformation = true
			m.AdditionalInformation = nasType.AdditionalInformation{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
