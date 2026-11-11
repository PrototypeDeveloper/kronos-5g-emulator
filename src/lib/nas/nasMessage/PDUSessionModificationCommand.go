package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionModificationCommand modifies a PDU session.
type PDUSessionModificationCommand struct {
	PDUSessionID, ProcedureTransactionID      uint8
	HasCause5GSM                              bool
	Cause5GSM                                 nasType.Cause5GSM
	HasSessionAMBR                            bool
	SessionAMBR                               nasType.SessionAMBR
	HasRQTimerValue                           bool
	RQTimerValue                              nasType.RQTimerValue
	HasAlwaysOnPDUSessionIndication           bool
	AlwaysOnPDUSessionIndication              nasType.AlwaysOnPDUSessionIndication
	HasAuthorizedQoSRules                     bool
	AuthorizedQoSRules                        nasType.QoSRules
	HasMappedEPSBearerContexts                bool
	MappedEPSBearerContexts                   nasType.MappedEPSBearerContexts
	HasAuthorizedQoSFlowDescriptions          bool
	AuthorizedQoSFlowDescriptions             nasType.QoSFlowDescriptions
	HasExtendedProtocolConfigurationOptions   bool
	ExtendedProtocolConfigurationOptions      nasType.ExtendedProtocolConfigurationOptions
	HasATSSSContainer                         bool
	ATSSSContainer                            nasType.ATSSSContainer
	HasIPHeaderCompressionConfiguration       bool
	IPHeaderCompressionConfiguration          nasType.IPHeaderCompressionConfiguration
	HasPortManagementInformationContainer     bool
	PortManagementInformationContainer        nasType.PortManagementInformationContainer
	HasServingPLMNRateControl                 bool
	ServingPLMNRateControl                    nasType.ServingPLMNRateControl
	HasEthernetHeaderCompressionConfiguration bool
	EthernetHeaderCompressionConfiguration    nasType.EthernetHeaderCompressionConfiguration
	UnknownIEs                                []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionModificationCommandOptions = nascodec.Schema{
	nasType.IECause5GSM:                              {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IESessionAMBR:                            {Format: nascodec.TLV, Min: 6, Max: 6},
	nasType.IERQTimerValue:                           {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEAlwaysOnPDUSessionIndication:           {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEQoSRules:                               {Format: nascodec.TLVE, Min: 4, Max: 65535},
	nasType.IEMappedEPSBearerContexts:                {Format: nascodec.TLVE, Min: 4, Max: 65535},
	nasType.IEQoSFlowDescriptions:                    {Format: nascodec.TLVE, Min: 3, Max: 65535},
	nasType.IEExtendedProtocolConfigurationOptions:   {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEATSSSContainer:                         {Format: nascodec.TLVE, Min: 0, Max: 65535},
	nasType.IEIPHeaderCompressionConfiguration:       {Format: nascodec.TLV, Min: 3, Max: 255},
	nasType.IEPortManagementInformationContainer:     {Format: nascodec.TLVE, Min: 5, Max: 65535},
	nasType.IEServingPLMNRateControl:                 {Format: nascodec.TLV, Min: 2, Max: 2},
	nasType.IEEthernetHeaderCompressionConfiguration: {Format: nascodec.TLV, Min: 1, Max: 1},
}

func (m *PDUSessionModificationCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionModificationCommand")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionModificationCommand, 255)
	var options []nasType.IE
	if m.HasCause5GSM {
		options = append(options, nasType.IE{ID: nasType.IECause5GSM, Value: []byte{m.Cause5GSM.Value}})
	}
	if m.HasSessionAMBR {
		options = append(options, nasType.IE{ID: nasType.IESessionAMBR, Value: m.SessionAMBR.Value})
	}
	if m.HasRQTimerValue {
		options = append(options, nasType.IE{ID: nasType.IERQTimerValue, Value: []byte{m.RQTimerValue.Value}})
	}
	if m.HasAlwaysOnPDUSessionIndication {
		options = append(options, nasType.IE{ID: nasType.IEAlwaysOnPDUSessionIndication, Value: []byte{m.AlwaysOnPDUSessionIndication.Value}})
	}
	if m.HasAuthorizedQoSRules {
		options = append(options, nasType.IE{ID: nasType.IEQoSRules, Value: m.AuthorizedQoSRules.Value})
	}
	if m.HasMappedEPSBearerContexts {
		options = append(options, nasType.IE{ID: nasType.IEMappedEPSBearerContexts, Value: m.MappedEPSBearerContexts.Value})
	}
	if m.HasAuthorizedQoSFlowDescriptions {
		options = append(options, nasType.IE{ID: nasType.IEQoSFlowDescriptions, Value: m.AuthorizedQoSFlowDescriptions.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasATSSSContainer {
		options = append(options, nasType.IE{ID: nasType.IEATSSSContainer, Value: m.ATSSSContainer.Value})
	}
	if m.HasIPHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEIPHeaderCompressionConfiguration, Value: m.IPHeaderCompressionConfiguration.Value})
	}
	if m.HasPortManagementInformationContainer {
		options = append(options, nasType.IE{ID: nasType.IEPortManagementInformationContainer, Value: m.PortManagementInformationContainer.Value})
	}
	if m.HasServingPLMNRateControl {
		options = append(options, nasType.IE{ID: nasType.IEServingPLMNRateControl, Value: m.ServingPLMNRateControl.Value})
	}
	if m.HasEthernetHeaderCompressionConfiguration {
		options = append(options, nasType.IE{ID: nasType.IEEthernetHeaderCompressionConfiguration, Value: []byte{m.EthernetHeaderCompressionConfiguration.Value}})
	}
	w.Options(options, pDUSessionModificationCommandOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionModificationCommandOptions)
	return w.Result()
}
func DecodePDUSessionModificationCommand(data []byte) (*PDUSessionModificationCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionModificationCommand {
		return nil, fmt.Errorf("unexpected header for PDUSessionModificationCommand")
	}
	m := &PDUSessionModificationCommand{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	known, unknown := r.OptionsWithUnknown(pDUSessionModificationCommandOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionModificationCommand: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IECause5GSM:
			m.HasCause5GSM = true
			m.Cause5GSM = nasType.Cause5GSM{Value: ie.Value[0]}
		case nasType.IESessionAMBR:
			m.HasSessionAMBR = true
			m.SessionAMBR = nasType.SessionAMBR{Value: ie.Value}
		case nasType.IERQTimerValue:
			m.HasRQTimerValue = true
			m.RQTimerValue = nasType.RQTimerValue{Value: ie.Value[0]}
		case nasType.IEAlwaysOnPDUSessionIndication:
			m.HasAlwaysOnPDUSessionIndication = true
			m.AlwaysOnPDUSessionIndication = nasType.AlwaysOnPDUSessionIndication{Value: ie.Value[0]}
		case nasType.IEQoSRules:
			m.HasAuthorizedQoSRules = true
			m.AuthorizedQoSRules = nasType.QoSRules{Value: ie.Value}
		case nasType.IEMappedEPSBearerContexts:
			m.HasMappedEPSBearerContexts = true
			m.MappedEPSBearerContexts = nasType.MappedEPSBearerContexts{Value: ie.Value}
		case nasType.IEQoSFlowDescriptions:
			m.HasAuthorizedQoSFlowDescriptions = true
			m.AuthorizedQoSFlowDescriptions = nasType.QoSFlowDescriptions{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEATSSSContainer:
			m.HasATSSSContainer = true
			m.ATSSSContainer = nasType.ATSSSContainer{Value: ie.Value}
		case nasType.IEIPHeaderCompressionConfiguration:
			m.HasIPHeaderCompressionConfiguration = true
			m.IPHeaderCompressionConfiguration = nasType.IPHeaderCompressionConfiguration{Value: ie.Value}
		case nasType.IEPortManagementInformationContainer:
			m.HasPortManagementInformationContainer = true
			m.PortManagementInformationContainer = nasType.PortManagementInformationContainer{Value: ie.Value}
		case nasType.IEServingPLMNRateControl:
			m.HasServingPLMNRateControl = true
			m.ServingPLMNRateControl = nasType.ServingPLMNRateControl{Value: ie.Value}
		case nasType.IEEthernetHeaderCompressionConfiguration:
			m.HasEthernetHeaderCompressionConfiguration = true
			m.EthernetHeaderCompressionConfiguration = nasType.EthernetHeaderCompressionConfiguration{Value: ie.Value[0]}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
