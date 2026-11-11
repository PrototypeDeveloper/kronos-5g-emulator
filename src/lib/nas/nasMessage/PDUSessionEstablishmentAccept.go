package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionEstablishmentAccept contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type PDUSessionEstablishmentAccept struct {
	PDUSessionID                            uint8
	ProcedureTransactionID                  uint8
	SelectedSSCModeAndPDUSessionType        uint8
	AuthorizedQoSRules                      []byte
	SessionAMBR                             []byte
	HasAuthorizedQoSFlowDescriptions        bool
	AuthorizedQoSFlowDescriptions           nasType.QoSFlowDescriptions
	HasCause5GSM                            bool
	Cause5GSM                               nasType.Cause5GSM
	HasPDUAddress                           bool
	PDUAddress                              nasType.PDUAddress
	HasRQTimerValue                         bool
	RQTimerValue                            nasType.RQTimerValue
	HasSNSSAI                               bool
	SNSSAI                                  nasType.SNSSAI
	HasAlwaysOnPDUSessionIndication         bool
	AlwaysOnPDUSessionIndication            nasType.AlwaysOnPDUSessionIndication
	HasMappedEPSBearerContexts              bool
	MappedEPSBearerContexts                 nasType.MappedEPSBearerContexts
	HasEAPMessage                           bool
	EAPMessage                              nasType.EAPMessage
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	HasDNN                                  bool
	DNN                                     nasType.DNN
}

var pDUSessionEstablishmentAcceptOptions = nascodec.Schema{
	nasType.IEQoSFlowDescriptions:                  {Format: nascodec.TLVE, Min: 3, Max: 65535},
	nasType.IECause5GSM:                            {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEPDUAddress:                           {Format: nascodec.TLV, Min: 5, Max: 29},
	nasType.IERQTimerValue:                         {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IESNSSAI:                               {Format: nascodec.TLV, Min: 1, Max: 8},
	nasType.IEAlwaysOnPDUSessionIndication:         {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEMappedEPSBearerContexts:              {Format: nascodec.TLVE, Min: 4, Max: 65535},
	nasType.IEEAPMessage:                           {Format: nascodec.TLVE, Min: 4, Max: 1500},
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEDNN:                                  {Format: nascodec.TLV, Min: 1, Max: 100},
}

func (m *PDUSessionEstablishmentAccept) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionEstablishmentAccept")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionEstablishmentAccept, 255)
	w.U8(m.SelectedSSCModeAndPDUSessionType, 255)
	w.LV(m.AuthorizedQoSRules, 2, 4, 65535)
	w.LV(m.SessionAMBR, 1, 6, 6)
	var options []nasType.IE
	if m.HasAuthorizedQoSFlowDescriptions {
		options = append(options, nasType.IE{ID: nasType.IEQoSFlowDescriptions, Value: m.AuthorizedQoSFlowDescriptions.Value})
	}
	if m.HasCause5GSM {
		options = append(options, nasType.IE{ID: nasType.IECause5GSM, Value: []byte{m.Cause5GSM.Value}})
	}
	if m.HasPDUAddress {
		options = append(options, nasType.IE{ID: nasType.IEPDUAddress, Value: m.PDUAddress.Value})
	}
	if m.HasRQTimerValue {
		options = append(options, nasType.IE{ID: nasType.IERQTimerValue, Value: []byte{m.RQTimerValue.Value}})
	}
	if m.HasSNSSAI {
		options = append(options, nasType.IE{ID: nasType.IESNSSAI, Value: m.SNSSAI.Value})
	}
	if m.HasAlwaysOnPDUSessionIndication {
		options = append(options, nasType.IE{ID: nasType.IEAlwaysOnPDUSessionIndication, Value: []byte{m.AlwaysOnPDUSessionIndication.Value}})
	}
	if m.HasMappedEPSBearerContexts {
		options = append(options, nasType.IE{ID: nasType.IEMappedEPSBearerContexts, Value: m.MappedEPSBearerContexts.Value})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasDNN {
		options = append(options, nasType.IE{ID: nasType.IEDNN, Value: m.DNN.Value})
	}
	w.Options(options, pDUSessionEstablishmentAcceptOptions)
	return w.Result()
}
func DecodePDUSessionEstablishmentAccept(data []byte) (*PDUSessionEstablishmentAccept, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionEstablishmentAccept {
		return nil, fmt.Errorf("unexpected header for PDUSessionEstablishmentAccept")
	}
	m := &PDUSessionEstablishmentAccept{}
	r := nascodec.Reader{Data: data[4:]}
	m.PDUSessionID = h.PDUSessionID
	m.ProcedureTransactionID = h.ProcedureTransactionID
	m.SelectedSSCModeAndPDUSessionType = r.U8(255)
	m.AuthorizedQoSRules = r.LV(2, 4, 65535)
	m.SessionAMBR = r.LV(1, 6, 6)
	for _, ie := range r.Options(pDUSessionEstablishmentAcceptOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEQoSFlowDescriptions:
			m.HasAuthorizedQoSFlowDescriptions = true
			m.AuthorizedQoSFlowDescriptions = nasType.QoSFlowDescriptions{Value: ie.Value}
		case nasType.IECause5GSM:
			m.HasCause5GSM = true
			m.Cause5GSM = nasType.Cause5GSM{Value: ie.Value[0]}
		case nasType.IEPDUAddress:
			m.HasPDUAddress = true
			m.PDUAddress = nasType.PDUAddress{Value: ie.Value}
		case nasType.IERQTimerValue:
			m.HasRQTimerValue = true
			m.RQTimerValue = nasType.RQTimerValue{Value: ie.Value[0]}
		case nasType.IESNSSAI:
			m.HasSNSSAI = true
			m.SNSSAI = nasType.SNSSAI{Value: ie.Value}
		case nasType.IEAlwaysOnPDUSessionIndication:
			m.HasAlwaysOnPDUSessionIndication = true
			m.AlwaysOnPDUSessionIndication = nasType.AlwaysOnPDUSessionIndication{Value: ie.Value[0]}
		case nasType.IEMappedEPSBearerContexts:
			m.HasMappedEPSBearerContexts = true
			m.MappedEPSBearerContexts = nasType.MappedEPSBearerContexts{Value: ie.Value}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEDNN:
			m.HasDNN = true
			m.DNN = nasType.DNN{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionEstablishmentAccept: %w", r.Err)
	}
	return m, nil
}
