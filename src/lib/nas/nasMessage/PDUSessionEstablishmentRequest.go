package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionEstablishmentRequest contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type PDUSessionEstablishmentRequest struct {
	PDUSessionID                              uint8
	ProcedureTransactionID                    uint8
	IntegrityProtectionMaximumDataRate        []byte
	HasPDUSessionType                         bool
	PDUSessionType                            nasType.PDUSessionType
	HasSSCMode                                bool
	SSCMode                                   nasType.SSCMode
	HasCapability5GSM                         bool
	Capability5GSM                            nasType.Capability5GSM
	HasMaximumNumberOfSupportedPacketFilters  bool
	MaximumNumberOfSupportedPacketFilters     nasType.MaximumNumberOfSupportedPacketFilters
	HasAlwaysOnPDUSessionRequested            bool
	AlwaysOnPDUSessionRequested               nasType.AlwaysOnPDUSessionRequested
	HasSMPDUDNRequestContainer                bool
	SMPDUDNRequestContainer                   nasType.SMPDUDNRequestContainer
	HasExtendedProtocolConfigurationOptions   bool
	ExtendedProtocolConfigurationOptions      nasType.ExtendedProtocolConfigurationOptions
	HasIPHeaderCompressionConfiguration       bool
	IPHeaderCompressionConfiguration          nasType.IPHeaderCompressionConfiguration
	HasDSTTEthernetPortMACAddress             bool
	DSTTEthernetPortMACAddress                nasType.DSTTEthernetPortMACAddress
	HasUEDSTTResidenceTime                    bool
	UEDSTTResidenceTime                       nasType.UEDSTTResidenceTime
	HasPortManagementInformationContainer     bool
	PortManagementInformationContainer        nasType.PortManagementInformationContainer
	HasEthernetHeaderCompressionConfiguration bool
	EthernetHeaderCompressionConfiguration    nasType.EthernetHeaderCompressionConfiguration
	HasSuggestedInterfaceIdentifier           bool
	SuggestedInterfaceIdentifier              nasType.PDUAddress
}

var pDUSessionEstablishmentRequestOptions = nascodec.Schema{
	nasType.IEPDUSessionType:                         {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IESSCMode:                                {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IECapability5GSM:                         {Format: nascodec.TLV, Min: 1, Max: 13},
	nasType.IEMaximumNumberOfSupportedPacketFilters:  {Format: nascodec.TV, Min: 2, Max: 2},
	nasType.IEAlwaysOnPDUSessionRequested:            {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IESMPDUDNRequestContainer:                {Format: nascodec.TLV, Min: 1, Max: 253},
	nasType.IEExtendedProtocolConfigurationOptions:   {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEIPHeaderCompressionConfiguration:       {Format: nascodec.TLV, Min: 3, Max: 255},
	nasType.IEDSTTEthernetPortMACAddress:             {Format: nascodec.TLV, Min: 6, Max: 6},
	nasType.IEUEDSTTResidenceTime:                    {Format: nascodec.TLV, Min: 8, Max: 8},
	nasType.IEPortManagementInformationContainer:     {Format: nascodec.TLVE, Min: 5, Max: 65535},
	nasType.IEEthernetHeaderCompressionConfiguration: {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IESuggestedInterfaceIdentifier:           {Format: nascodec.TLV, Min: 9, Max: 9},
}

func (m *PDUSessionEstablishmentRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionEstablishmentRequest")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionEstablishmentRequest, 255)
	w.Fixed(m.IntegrityProtectionMaximumDataRate, 2)
	var options []nasType.IE
	if m.HasPDUSessionType {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionType, Value: []byte{m.PDUSessionType.Value}})
	}
	if m.HasSSCMode {
		options = append(options, nasType.IE{ID: nasType.IESSCMode, Value: []byte{m.SSCMode.Value}})
	}
	if m.HasCapability5GSM {
		options = append(options, nasType.IE{ID: nasType.IECapability5GSM, Value: m.Capability5GSM.Value})
	}
	if m.HasMaximumNumberOfSupportedPacketFilters {
		options = append(options, nasType.IE{ID: nasType.IEMaximumNumberOfSupportedPacketFilters, Value: m.MaximumNumberOfSupportedPacketFilters.Value})
	}
	if m.HasAlwaysOnPDUSessionRequested {
		options = append(options, nasType.IE{ID: nasType.IEAlwaysOnPDUSessionRequested, Value: []byte{m.AlwaysOnPDUSessionRequested.Value}})
	}
	if m.HasSMPDUDNRequestContainer {
		options = append(options, nasType.IE{ID: nasType.IESMPDUDNRequestContainer, Value: m.SMPDUDNRequestContainer.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasIPHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEIPHeaderCompressionConfiguration, Value: m.IPHeaderCompressionConfiguration.Value})
	}
	if m.HasDSTTEthernetPortMACAddress {
		options = append(options, nasType.IE{ID: nasType.IEDSTTEthernetPortMACAddress, Value: m.DSTTEthernetPortMACAddress.Value})
	}
	if m.HasUEDSTTResidenceTime {
		options = append(options, nasType.IE{ID: nasType.IEUEDSTTResidenceTime, Value: m.UEDSTTResidenceTime.Value})
	}
	if m.HasPortManagementInformationContainer {
		options = append(options, nasType.IE{ID: nasType.IEPortManagementInformationContainer, Value: m.PortManagementInformationContainer.Value})
	}
	if m.HasEthernetHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEEthernetHeaderCompressionConfiguration, Value: []byte{m.EthernetHeaderCompressionConfiguration.Value}})
	}
	if m.HasSuggestedInterfaceIdentifier {
		options = append(options, nasType.IE{ID: nasType.IESuggestedInterfaceIdentifier, Value: m.SuggestedInterfaceIdentifier.Value})
	}
	w.Options(options, pDUSessionEstablishmentRequestOptions)
	return w.Result()
}
func DecodePDUSessionEstablishmentRequest(data []byte) (*PDUSessionEstablishmentRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionEstablishmentRequest {
		return nil, fmt.Errorf("unexpected header for PDUSessionEstablishmentRequest")
	}
	m := &PDUSessionEstablishmentRequest{}
	r := nascodec.Reader{Data: data[4:]}
	m.PDUSessionID = h.PDUSessionID
	m.ProcedureTransactionID = h.ProcedureTransactionID
	m.IntegrityProtectionMaximumDataRate = r.Take(2)
	for _, ie := range r.Options(pDUSessionEstablishmentRequestOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEPDUSessionType:
			m.HasPDUSessionType = true
			m.PDUSessionType = nasType.PDUSessionType{Value: ie.Value[0]}
		case nasType.IESSCMode:
			m.HasSSCMode = true
			m.SSCMode = nasType.SSCMode{Value: ie.Value[0]}
		case nasType.IECapability5GSM:
			m.HasCapability5GSM = true
			m.Capability5GSM = nasType.Capability5GSM{Value: ie.Value}
		case nasType.IEMaximumNumberOfSupportedPacketFilters:
			m.HasMaximumNumberOfSupportedPacketFilters = true
			m.MaximumNumberOfSupportedPacketFilters = nasType.MaximumNumberOfSupportedPacketFilters{Value: ie.Value}
		case nasType.IEAlwaysOnPDUSessionRequested:
			m.HasAlwaysOnPDUSessionRequested = true
			m.AlwaysOnPDUSessionRequested = nasType.AlwaysOnPDUSessionRequested{Value: ie.Value[0]}
		case nasType.IESMPDUDNRequestContainer:
			m.HasSMPDUDNRequestContainer = true
			m.SMPDUDNRequestContainer = nasType.SMPDUDNRequestContainer{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEIPHeaderCompressionConfiguration:
			m.HasIPHeaderCompressionConfiguration = true
			m.IPHeaderCompressionConfiguration = nasType.IPHeaderCompressionConfiguration{Value: ie.Value}
		case nasType.IEDSTTEthernetPortMACAddress:
			m.HasDSTTEthernetPortMACAddress = true
			m.DSTTEthernetPortMACAddress = nasType.DSTTEthernetPortMACAddress{Value: ie.Value}
		case nasType.IEUEDSTTResidenceTime:
			m.HasUEDSTTResidenceTime = true
			m.UEDSTTResidenceTime = nasType.UEDSTTResidenceTime{Value: ie.Value}
		case nasType.IEPortManagementInformationContainer:
			m.HasPortManagementInformationContainer = true
			m.PortManagementInformationContainer = nasType.PortManagementInformationContainer{Value: ie.Value}
		case nasType.IEEthernetHeaderCompressionConfiguration:
			m.HasEthernetHeaderCompressionConfiguration = true
			m.EthernetHeaderCompressionConfiguration = nasType.EthernetHeaderCompressionConfiguration{Value: ie.Value[0]}
		case nasType.IESuggestedInterfaceIdentifier:
			m.HasSuggestedInterfaceIdentifier = true
			m.SuggestedInterfaceIdentifier = nasType.PDUAddress{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionEstablishmentRequest: %w", r.Err)
	}
	return m, nil
}
