package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionReleaseReject carries a plain 5GSM message (TS 24.501 8.3.12–15).
type PDUSessionReleaseReject struct {
	PDUSessionID, ProcedureTransactionID    uint8
	Cause                                   uint8
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
}

var pDUSessionReleaseRejectOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionReleaseReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionReleaseReject")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionReleaseReject, 255)
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionReleaseRejectOptions)
	return w.Result()
}
func DecodePDUSessionReleaseReject(data []byte) (*PDUSessionReleaseReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionReleaseReject {
		return nil, fmt.Errorf("unexpected PDUSessionReleaseReject header")
	}
	m := &PDUSessionReleaseReject{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.Cause = r.U8(255)
	for _, ie := range r.Options(pDUSessionReleaseRejectOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionReleaseReject: %w", r.Err)
	}
	return m, nil
}
