package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ULNASTransport contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type ULNASTransport struct {
	PayloadContainerType     uint8
	PayloadContainer         []byte
	HasPDUSessionID          bool
	PDUSessionID             nasType.PDUSessionID
	HasOldPDUSessionID       bool
	OldPDUSessionID          nasType.PDUSessionID
	HasRequestType           bool
	RequestType              nasType.RequestType
	HasSNSSAI                bool
	SNSSAI                   nasType.SNSSAI
	HasDNN                   bool
	DNN                      nasType.DNN
	HasAdditionalInformation bool
	AdditionalInformation    nasType.AdditionalInformation
}

var uLNASTransportOptions = nascodec.Schema{
	nasType.IEPDUSessionID:          {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEOldPDUSessionID:       {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IERequestType:           {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IESNSSAI:                {Format: nascodec.TLV, Min: 1, Max: 8},
	nasType.IEDNN:                   {Format: nascodec.TLV, Min: 1, Max: 100},
	nasType.IEAdditionalInformation: {Format: nascodec.TLV, Min: 1, Max: 255},
}

func (m *ULNASTransport) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ULNASTransport")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgULNASTransport}
	w.U8(m.PayloadContainerType, 15)
	w.LV(m.PayloadContainer, 2, 1, 65535)
	var options []nasType.IE
	if m.HasPDUSessionID {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionID, Value: []byte{m.PDUSessionID.Value}})
	}
	if m.HasOldPDUSessionID {
		options = append(options, nasType.IE{ID: nasType.IEOldPDUSessionID, Value: []byte{m.OldPDUSessionID.Value}})
	}
	if m.HasRequestType {
		options = append(options, nasType.IE{ID: nasType.IERequestType, Value: []byte{m.RequestType.Value}})
	}
	if m.HasSNSSAI {
		options = append(options, nasType.IE{ID: nasType.IESNSSAI, Value: m.SNSSAI.Value})
	}
	if m.HasDNN {
		options = append(options, nasType.IE{ID: nasType.IEDNN, Value: m.DNN.Value})
	}
	if m.HasAdditionalInformation {
		options = append(options, nasType.IE{ID: nasType.IEAdditionalInformation, Value: m.AdditionalInformation.Value})
	}
	w.Options(options, uLNASTransportOptions)
	return w.Result()
}
func DecodeULNASTransport(data []byte) (*ULNASTransport, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgULNASTransport {
		return nil, fmt.Errorf("unexpected header for ULNASTransport")
	}
	m := &ULNASTransport{}
	r := nascodec.Reader{Data: data[3:]}
	m.PayloadContainerType = r.U8(15)
	m.PayloadContainer = r.LV(2, 1, 65535)
	for _, ie := range r.Options(uLNASTransportOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEPDUSessionID:
			m.HasPDUSessionID = true
			m.PDUSessionID = nasType.PDUSessionID{Value: ie.Value[0]}
		case nasType.IEOldPDUSessionID:
			m.HasOldPDUSessionID = true
			m.OldPDUSessionID = nasType.PDUSessionID{Value: ie.Value[0]}
		case nasType.IERequestType:
			m.HasRequestType = true
			m.RequestType = nasType.RequestType{Value: ie.Value[0]}
		case nasType.IESNSSAI:
			m.HasSNSSAI = true
			m.SNSSAI = nasType.SNSSAI{Value: ie.Value}
		case nasType.IEDNN:
			m.HasDNN = true
			m.DNN = nasType.DNN{Value: ie.Value}
		case nasType.IEAdditionalInformation:
			m.HasAdditionalInformation = true
			m.AdditionalInformation = nasType.AdditionalInformation{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode ULNASTransport: %w", r.Err)
	}
	return m, nil
}
