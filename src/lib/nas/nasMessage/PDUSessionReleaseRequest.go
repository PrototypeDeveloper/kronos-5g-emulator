package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionReleaseRequest carries a plain 5GSM message (TS 24.501 8.3.12–15).
type PDUSessionReleaseRequest struct {
	PDUSessionID, ProcedureTransactionID    uint8
	HasCause5GSM                            bool
	Cause5GSM                               nasType.Cause5GSM
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
}

var pDUSessionReleaseRequestOptions = nascodec.Schema{
	nasType.IECause5GSM:                            {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionReleaseRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionReleaseRequest")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionReleaseRequest, 255)

	var options []nasType.IE
	if m.HasCause5GSM {
		options = append(options, nasType.IE{ID: nasType.IECause5GSM, Value: []byte{m.Cause5GSM.Value}})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionReleaseRequestOptions)
	return w.Result()
}
func DecodePDUSessionReleaseRequest(data []byte) (*PDUSessionReleaseRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionReleaseRequest {
		return nil, fmt.Errorf("unexpected PDUSessionReleaseRequest header")
	}
	m := &PDUSessionReleaseRequest{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}

	for _, ie := range r.Options(pDUSessionReleaseRequestOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IECause5GSM:
			m.HasCause5GSM = true
			m.Cause5GSM = nasType.Cause5GSM{Value: ie.Value[0]}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionReleaseRequest: %w", r.Err)
	}
	return m, nil
}
