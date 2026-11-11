package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// RegistrationAccept contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type RegistrationAccept struct {
	RegistrationResult                        []byte
	HasGUTI5GS                                bool
	GUTI5GS                                   nasType.GUTI5GS
	HasEquivalentPLMNs                        bool
	EquivalentPLMNs                           nasType.EquivalentPLMNs
	HasTAIList                                bool
	TAIList                                   nasType.TAIList
	HasAllowedNSSAI                           bool
	AllowedNSSAI                              nasType.AllowedNSSAI
	HasRejectedNSSAI                          bool
	RejectedNSSAI                             nasType.RejectedNSSAI
	HasConfiguredNSSAI                        bool
	ConfiguredNSSAI                           nasType.ConfiguredNSSAI
	HasNetworkFeatureSupport5GS               bool
	NetworkFeatureSupport5GS                  nasType.NetworkFeatureSupport5GS
	HasPDUSessionStatus                       bool
	PDUSessionStatus                          nasType.PDUSessionStatus
	HasPDUSessionReactivationResult           bool
	PDUSessionReactivationResult              nasType.PDUSessionReactivationResult
	HasPDUSessionReactivationResultErrorCause bool
	PDUSessionReactivationResultErrorCause    nasType.PDUSessionReactivationResultErrorCause
	HasT3512Value                             bool
	T3512Value                                nasType.T3512Value
	HasT3502Value                             bool
	T3502Value                                nasType.T3502Value
	HasNon3GPPDeregistrationTimerValue        bool
	Non3GPPDeregistrationTimerValue           nasType.Non3GPPDeregistrationTimerValue
	HasServiceAreaList                        bool
	ServiceAreaList                           nasType.ServiceAreaList
	HasT3324Value                             bool
	T3324Value                                nasType.T3324Value
	HasMICOIndication                         bool
	MICOIndication                            nasType.MICOIndication
	HasNetworkSlicingIndication               bool
	NetworkSlicingIndication                  nasType.NetworkSlicingIndication
}

var registrationAcceptOptions = nascodec.Schema{
	nasType.IEGUTI5GS:                                {Format: nascodec.TLVE, Min: 11, Max: 11},
	nasType.IEEquivalentPLMNs:                        {Format: nascodec.TLV, Min: 3, Max: 45},
	nasType.IETAIList:                                {Format: nascodec.TLV, Min: 7, Max: 112},
	nasType.IEAllowedNSSAI:                           {Format: nascodec.TLV, Min: 2, Max: 72},
	nasType.IERejectedNSSAI:                          {Format: nascodec.TLV, Min: 2, Max: 40},
	nasType.IEConfiguredNSSAI:                        {Format: nascodec.TLV, Min: 2, Max: 144},
	nasType.IENetworkFeatureSupport5GS:               {Format: nascodec.TLV, Min: 1, Max: 4},
	nasType.IEPDUSessionStatus:                       {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEPDUSessionReactivationResult:           {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IEPDUSessionReactivationResultErrorCause: {Format: nascodec.TLVE, Min: 2, Max: 512},
	nasType.IET3512Value:                             {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IET3502Value:                             {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IENon3GPPDeregistrationTimerValue:        {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEServiceAreaList:                        {Format: nascodec.TLV, Min: 4, Max: 112},
	nasType.IET3324Value:                             {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEMICOIndication:                         {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IENetworkSlicingIndication:               {Format: nascodec.HalfTV, Min: 1, Max: 1},
}

func (m *RegistrationAccept) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil RegistrationAccept")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgRegistrationAccept}
	w.LV(m.RegistrationResult, 1, 1, 1)
	var options []nasType.IE
	if m.HasGUTI5GS {
		options = append(options, nasType.IE{ID: nasType.IEGUTI5GS, Value: m.GUTI5GS.Value})
	}
	if m.HasEquivalentPLMNs {
		options = append(options, nasType.IE{ID: nasType.IEEquivalentPLMNs, Value: m.EquivalentPLMNs.Value})
	}
	if m.HasTAIList {
		options = append(options, nasType.IE{ID: nasType.IETAIList, Value: m.TAIList.Value})
	}
	if m.HasAllowedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IEAllowedNSSAI, Value: m.AllowedNSSAI.Value})
	}
	if m.HasRejectedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IERejectedNSSAI, Value: m.RejectedNSSAI.Value})
	}
	if m.HasConfiguredNSSAI {
		options = append(options, nasType.IE{ID: nasType.IEConfiguredNSSAI, Value: m.ConfiguredNSSAI.Value})
	}
	if m.HasNetworkFeatureSupport5GS {
		options = append(options, nasType.IE{ID: nasType.IENetworkFeatureSupport5GS, Value: m.NetworkFeatureSupport5GS.Value})
	}
	if m.HasPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionStatus, Value: m.PDUSessionStatus.Value})
	}
	if m.HasPDUSessionReactivationResult {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionReactivationResult, Value: m.PDUSessionReactivationResult.Value})
	}
	if m.HasPDUSessionReactivationResultErrorCause {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionReactivationResultErrorCause, Value: m.PDUSessionReactivationResultErrorCause.Value})
	}
	if m.HasT3512Value {
		options = append(options, nasType.IE{ID: nasType.IET3512Value, Value: []byte{m.T3512Value.Value}})
	}
	if m.HasT3502Value {
		options = append(options, nasType.IE{ID: nasType.IET3502Value, Value: []byte{m.T3502Value.Value}})
	}
	if m.HasNon3GPPDeregistrationTimerValue {
		options = append(options, nasType.IE{ID: nasType.IENon3GPPDeregistrationTimerValue, Value: []byte{m.Non3GPPDeregistrationTimerValue.Value}})
	}
	if m.HasServiceAreaList {
		options = append(options, nasType.IE{ID: nasType.IEServiceAreaList, Value: m.ServiceAreaList.Value})
	}
	if m.HasT3324Value {
		options = append(options, nasType.IE{ID: nasType.IET3324Value, Value: []byte{m.T3324Value.Value}})
	}
	if m.HasMICOIndication {
		options = append(options, nasType.IE{ID: nasType.IEMICOIndication, Value: []byte{m.MICOIndication.Value}})
	}
	if m.HasNetworkSlicingIndication {
		options = append(options, nasType.IE{ID: nasType.IENetworkSlicingIndication, Value: []byte{m.NetworkSlicingIndication.Value}})
	}
	w.Options(options, registrationAcceptOptions)
	return w.Result()
}
func DecodeRegistrationAccept(data []byte) (*RegistrationAccept, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgRegistrationAccept {
		return nil, fmt.Errorf("unexpected header for RegistrationAccept")
	}
	m := &RegistrationAccept{}
	r := nascodec.Reader{Data: data[3:]}
	m.RegistrationResult = r.LV(1, 1, 1)
	for _, ie := range r.Options(registrationAcceptOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEGUTI5GS:
			m.HasGUTI5GS = true
			m.GUTI5GS = nasType.GUTI5GS{Value: ie.Value}
		case nasType.IEEquivalentPLMNs:
			m.HasEquivalentPLMNs = true
			m.EquivalentPLMNs = nasType.EquivalentPLMNs{Value: ie.Value}
		case nasType.IETAIList:
			m.HasTAIList = true
			m.TAIList = nasType.TAIList{Value: ie.Value}
		case nasType.IEAllowedNSSAI:
			m.HasAllowedNSSAI = true
			m.AllowedNSSAI = nasType.AllowedNSSAI{Value: ie.Value}
		case nasType.IERejectedNSSAI:
			m.HasRejectedNSSAI = true
			m.RejectedNSSAI = nasType.RejectedNSSAI{Value: ie.Value}
		case nasType.IEConfiguredNSSAI:
			m.HasConfiguredNSSAI = true
			m.ConfiguredNSSAI = nasType.ConfiguredNSSAI{Value: ie.Value}
		case nasType.IENetworkFeatureSupport5GS:
			m.HasNetworkFeatureSupport5GS = true
			m.NetworkFeatureSupport5GS = nasType.NetworkFeatureSupport5GS{Value: ie.Value}
		case nasType.IEPDUSessionStatus:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = nasType.PDUSessionStatus{Value: ie.Value}
		case nasType.IEPDUSessionReactivationResult:
			m.HasPDUSessionReactivationResult = true
			m.PDUSessionReactivationResult = nasType.PDUSessionReactivationResult{Value: ie.Value}
		case nasType.IEPDUSessionReactivationResultErrorCause:
			m.HasPDUSessionReactivationResultErrorCause = true
			m.PDUSessionReactivationResultErrorCause = nasType.PDUSessionReactivationResultErrorCause{Value: ie.Value}
		case nasType.IET3512Value:
			m.HasT3512Value = true
			m.T3512Value = nasType.T3512Value{Value: ie.Value[0]}
		case nasType.IET3502Value:
			m.HasT3502Value = true
			m.T3502Value = nasType.T3502Value{Value: ie.Value[0]}
		case nasType.IENon3GPPDeregistrationTimerValue:
			m.HasNon3GPPDeregistrationTimerValue = true
			m.Non3GPPDeregistrationTimerValue = nasType.Non3GPPDeregistrationTimerValue{Value: ie.Value[0]}
		case nasType.IEServiceAreaList:
			m.HasServiceAreaList = true
			m.ServiceAreaList = nasType.ServiceAreaList{Value: ie.Value}
		case nasType.IET3324Value:
			m.HasT3324Value = true
			m.T3324Value = nasType.T3324Value{Value: ie.Value[0]}
		case nasType.IEMICOIndication:
			m.HasMICOIndication = true
			m.MICOIndication = nasType.MICOIndication{Value: ie.Value[0]}
		case nasType.IENetworkSlicingIndication:
			m.HasNetworkSlicingIndication = true
			m.NetworkSlicingIndication = nasType.NetworkSlicingIndication{Value: ie.Value[0]}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode RegistrationAccept: %w", r.Err)
	}
	return m, nil
}
