package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionModificationRequest asks to modify a PDU session.
type PDUSessionModificationRequest struct {
	PDUSessionID, ProcedureTransactionID      uint8
	HasCapability5GSM                         bool
	Capability5GSM                            nasType.Capability5GSM
	HasCause5GSM                              bool
	Cause5GSM                                 nasType.Cause5GSM
	HasMaximumNumberOfSupportedPacketFilters  bool
	MaximumNumberOfSupportedPacketFilters     nasType.MaximumNumberOfSupportedPacketFilters
	HasAlwaysOnPDUSessionRequested            bool
	AlwaysOnPDUSessionRequested               nasType.AlwaysOnPDUSessionRequested
	HasIntegrityProtectionMaximumDataRate     bool
	IntegrityProtectionMaximumDataRate        nasType.IntegrityProtectionMaximumDataRate
	HasRequestedQoSRules                      bool
	RequestedQoSRules                         nasType.QoSRules
	HasRequestedQoSFlowDescriptions           bool
	RequestedQoSFlowDescriptions              nasType.QoSFlowDescriptions
	HasMappedEPSBearerContexts                bool
	MappedEPSBearerContexts                   nasType.MappedEPSBearerContexts
	HasExtendedProtocolConfigurationOptions   bool
	ExtendedProtocolConfigurationOptions      nasType.ExtendedProtocolConfigurationOptions
	HasPortManagementInformationContainer     bool
	PortManagementInformationContainer        nasType.PortManagementInformationContainer
	HasIPHeaderCompressionConfiguration       bool
	IPHeaderCompressionConfiguration          nasType.IPHeaderCompressionConfiguration
	HasEthernetHeaderCompressionConfiguration bool
	EthernetHeaderCompressionConfiguration    nasType.EthernetHeaderCompressionConfiguration
	UnknownIEs                                []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionModificationRequestOptions = nascodec.Schema{
	nasType.IECapability5GSM:                         {Format: nascodec.TLV, Min: 1, Max: 13},
	nasType.IECause5GSM:                              {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEMaximumNumberOfSupportedPacketFilters:  {Format: nascodec.TV, Min: 2, Max: 2},
	nasType.IEAlwaysOnPDUSessionRequested:            {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEIntegrityProtectionMaximumDataRate:     {Format: nascodec.TV, Min: 2, Max: 2},
	nasType.IEQoSRules:                               {Format: nascodec.TLVE, Min: 4, Max: 65535},
	nasType.IEQoSFlowDescriptions:                    {Format: nascodec.TLVE, Min: 3, Max: 65535},
	nasType.IEMappedEPSBearerContexts:                {Format: nascodec.TLVE, Min: 4, Max: 65535},
	nasType.IEExtendedProtocolConfigurationOptions:   {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEPortManagementInformationContainer:     {Format: nascodec.TLVE, Min: 5, Max: 65535},
	nasType.IEIPHeaderCompressionConfiguration:       {Format: nascodec.TLV, Min: 3, Max: 255},
	nasType.IEEthernetHeaderCompressionConfiguration: {Format: nascodec.TLV, Min: 1, Max: 1},
}

func (m *PDUSessionModificationRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionModificationRequest")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionModificationRequest, 255)
	var options []nasType.IE
	if m.HasCapability5GSM {
		options = append(options, nasType.IE{ID: nasType.IECapability5GSM, Value: m.Capability5GSM.Value})
	}
	if m.HasCause5GSM {
		options = append(options, nasType.IE{ID: nasType.IECause5GSM, Value: []byte{m.Cause5GSM.Value}})
	}
	if m.HasMaximumNumberOfSupportedPacketFilters {
		options = append(options, nasType.IE{ID: nasType.IEMaximumNumberOfSupportedPacketFilters, Value: m.MaximumNumberOfSupportedPacketFilters.Value})
	}
	if m.HasAlwaysOnPDUSessionRequested {
		options = append(options, nasType.IE{ID: nasType.IEAlwaysOnPDUSessionRequested, Value: []byte{m.AlwaysOnPDUSessionRequested.Value}})
	}
	if m.HasIntegrityProtectionMaximumDataRate {
		options = append(options, nasType.IE{ID: nasType.IEIntegrityProtectionMaximumDataRate, Value: m.IntegrityProtectionMaximumDataRate.Value})
	}
	if m.HasRequestedQoSRules {
		options = append(options, nasType.IE{ID: nasType.IEQoSRules, Value: m.RequestedQoSRules.Value})
	}
	if m.HasRequestedQoSFlowDescriptions {
		options = append(options, nasType.IE{ID: nasType.IEQoSFlowDescriptions, Value: m.RequestedQoSFlowDescriptions.Value})
	}
	if m.HasMappedEPSBearerContexts {
		options = append(options, nasType.IE{ID: nasType.IEMappedEPSBearerContexts, Value: m.MappedEPSBearerContexts.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasPortManagementInformationContainer {
		options = append(options, nasType.IE{ID: nasType.IEPortManagementInformationContainer, Value: m.PortManagementInformationContainer.Value})
	}
	if m.HasIPHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEIPHeaderCompressionConfiguration, Value: m.IPHeaderCompressionConfiguration.Value})
	}
	if m.HasEthernetHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEEthernetHeaderCompressionConfiguration, Value: []byte{m.EthernetHeaderCompressionConfiguration.Value}})
	}
	w.Options(options, pDUSessionModificationRequestOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionModificationRequestOptions)
	return w.Result()
}
func DecodePDUSessionModificationRequest(data []byte) (*PDUSessionModificationRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionModificationRequest {
		return nil, fmt.Errorf("unexpected header for PDUSessionModificationRequest")
	}
	m := &PDUSessionModificationRequest{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	known, unknown := r.OptionsWithUnknown(pDUSessionModificationRequestOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionModificationRequest: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IECapability5GSM:
			m.HasCapability5GSM = true
			m.Capability5GSM = nasType.Capability5GSM{Value: ie.Value}
		case nasType.IECause5GSM:
			m.HasCause5GSM = true
			m.Cause5GSM = nasType.Cause5GSM{Value: ie.Value[0]}
		case nasType.IEMaximumNumberOfSupportedPacketFilters:
			m.HasMaximumNumberOfSupportedPacketFilters = true
			m.MaximumNumberOfSupportedPacketFilters = nasType.MaximumNumberOfSupportedPacketFilters{Value: ie.Value}
		case nasType.IEAlwaysOnPDUSessionRequested:
			m.HasAlwaysOnPDUSessionRequested = true
			m.AlwaysOnPDUSessionRequested = nasType.AlwaysOnPDUSessionRequested{Value: ie.Value[0]}
		case nasType.IEIntegrityProtectionMaximumDataRate:
			m.HasIntegrityProtectionMaximumDataRate = true
			m.IntegrityProtectionMaximumDataRate = nasType.IntegrityProtectionMaximumDataRate{Value: ie.Value}
		case nasType.IEQoSRules:
			m.HasRequestedQoSRules = true
			m.RequestedQoSRules = nasType.QoSRules{Value: ie.Value}
		case nasType.IEQoSFlowDescriptions:
			m.HasRequestedQoSFlowDescriptions = true
			m.RequestedQoSFlowDescriptions = nasType.QoSFlowDescriptions{Value: ie.Value}
		case nasType.IEMappedEPSBearerContexts:
			m.HasMappedEPSBearerContexts = true
			m.MappedEPSBearerContexts = nasType.MappedEPSBearerContexts{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEPortManagementInformationContainer:
			m.HasPortManagementInformationContainer = true
			m.PortManagementInformationContainer = nasType.PortManagementInformationContainer{Value: ie.Value}
		case nasType.IEIPHeaderCompressionConfiguration:
			m.HasIPHeaderCompressionConfiguration = true
			m.IPHeaderCompressionConfiguration = nasType.IPHeaderCompressionConfiguration{Value: ie.Value}
		case nasType.IEEthernetHeaderCompressionConfiguration:
			m.HasEthernetHeaderCompressionConfiguration = true
			m.EthernetHeaderCompressionConfiguration = nasType.EthernetHeaderCompressionConfiguration{Value: ie.Value[0]}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
