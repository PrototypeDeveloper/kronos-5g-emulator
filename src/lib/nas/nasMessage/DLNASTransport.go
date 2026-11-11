package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// DLNASTransport contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type DLNASTransport struct {
	PayloadContainerType     uint8
	PayloadContainer         []byte
	HasPDUSessionID          bool
	PDUSessionID             nasType.PDUSessionID
	HasAdditionalInformation bool
	AdditionalInformation    nasType.AdditionalInformation
	HasCause5GMM             bool
	Cause5GMM                nasType.Cause5GMM
	HasBackOffTimerValue     bool
	BackOffTimerValue        nasType.BackOffTimerValue
}

var dLNASTransportOptions = nascodec.Schema{
	nasType.IEPDUSessionID:          {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEAdditionalInformation: {Format: nascodec.TLV, Min: 1, Max: 255},
	nasType.IECause5GMM:             {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEBackOffTimerValue:     {Format: nascodec.TLV, Min: 1, Max: 1},
}

func (m *DLNASTransport) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil DLNASTransport")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgDLNASTransport}
	w.U8(m.PayloadContainerType, 15)
	w.LV(m.PayloadContainer, 2, 1, 65535)
	var options []nasType.IE
	if m.HasPDUSessionID {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionID, Value: []byte{m.PDUSessionID.Value}})
	}
	if m.HasAdditionalInformation {
		options = append(options, nasType.IE{ID: nasType.IEAdditionalInformation, Value: m.AdditionalInformation.Value})
	}
	if m.HasCause5GMM {
		options = append(options, nasType.IE{ID: nasType.IECause5GMM, Value: []byte{m.Cause5GMM.Value}})
	}
	if m.HasBackOffTimerValue {
		options = append(options, nasType.IE{ID: nasType.IEBackOffTimerValue, Value: []byte{m.BackOffTimerValue.Value}})
	}
	w.Options(options, dLNASTransportOptions)
	return w.Result()
}
func DecodeDLNASTransport(data []byte) (*DLNASTransport, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgDLNASTransport {
		return nil, fmt.Errorf("unexpected header for DLNASTransport")
	}
	m := &DLNASTransport{}
	r := nascodec.Reader{Data: data[3:]}
	m.PayloadContainerType = r.U8(15)
	m.PayloadContainer = r.LV(2, 1, 65535)
	for _, ie := range r.Options(dLNASTransportOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEPDUSessionID:
			m.HasPDUSessionID = true
			m.PDUSessionID = nasType.PDUSessionID{Value: ie.Value[0]}
		case nasType.IEAdditionalInformation:
			m.HasAdditionalInformation = true
			m.AdditionalInformation = nasType.AdditionalInformation{Value: ie.Value}
		case nasType.IECause5GMM:
			m.HasCause5GMM = true
			m.Cause5GMM = nasType.Cause5GMM{Value: ie.Value[0]}
		case nasType.IEBackOffTimerValue:
			m.HasBackOffTimerValue = true
			m.BackOffTimerValue = nasType.BackOffTimerValue{Value: ie.Value[0]}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode DLNASTransport: %w", r.Err)
	}
	return m, nil
}
