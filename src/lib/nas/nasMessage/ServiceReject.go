package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ServiceReject rejects a service request.
type ServiceReject struct {
	Cause                 uint8 // 5GMM cause
	HasPDUSessionStatus   bool
	PDUSessionStatus      nasType.PDUSessionStatus
	HasT3346Value         bool
	T3346Value            nasType.T3346Value
	HasEAPMessage         bool
	EAPMessage            nasType.EAPMessage
	HasT3448Value         bool
	T3448Value            nasType.T3448Value
	HasCAGInformationList bool
	CAGInformationList    nasType.CAGInformationList
	UnknownIEs            []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var serviceRejectOptions = nascodec.Schema{
	nasType.IEPDUSessionStatus:   {Format: nascodec.TLV, Min: 2, Max: 32},
	nasType.IET3346Value:         {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEEAPMessage:         {Format: nascodec.TLVE, Min: 4, Max: 1500},
	nasType.IET3448Value:         {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IECAGInformationList: {Format: nascodec.TLVE, Min: 0, Max: 65535},
}

func (m *ServiceReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ServiceReject")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgServiceReject}
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionStatus, Value: m.PDUSessionStatus.Value})
	}
	if m.HasT3346Value {
		options = append(options, nasType.IE{ID: nasType.IET3346Value, Value: []byte{m.T3346Value.Value}})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	if m.HasT3448Value {
		options = append(options, nasType.IE{ID: nasType.IET3448Value, Value: []byte{m.T3448Value.Value}})
	}
	if m.HasCAGInformationList {
		options = append(options, nasType.IE{ID: nasType.IECAGInformationList, Value: m.CAGInformationList.Value})
	}
	w.Options(options, serviceRejectOptions)
	w.UnknownOptions(m.UnknownIEs, serviceRejectOptions)
	return w.Result()
}
func DecodeServiceReject(data []byte) (*ServiceReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgServiceReject {
		return nil, fmt.Errorf("unexpected header for ServiceReject")
	}
	m := &ServiceReject{}
	r := nascodec.Reader{Data: data[3:]}
	m.Cause = r.U8(255)
	known, unknown := r.OptionsWithUnknown(serviceRejectOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode ServiceReject: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEPDUSessionStatus:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = nasType.PDUSessionStatus{Value: ie.Value}
		case nasType.IET3346Value:
			m.HasT3346Value = true
			m.T3346Value = nasType.T3346Value{Value: ie.Value[0]}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		case nasType.IET3448Value:
			m.HasT3448Value = true
			m.T3448Value = nasType.T3448Value{Value: ie.Value[0]}
		case nasType.IECAGInformationList:
			m.HasCAGInformationList = true
			m.CAGInformationList = nasType.CAGInformationList{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
