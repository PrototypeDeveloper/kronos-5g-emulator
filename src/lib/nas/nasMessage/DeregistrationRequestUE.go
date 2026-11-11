package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// DeregistrationRequestUE is the UE-originating message (TS 24.501 8.2.12).
// The high nibble is ngKSI; the low nibble is the de-registration type.
type DeregistrationRequestUE struct {
	DeregistrationTypeAndNgKSI   uint8
	MobileIdentity               []byte
	HasUnavailabilityInformation bool
	UnavailabilityInformation    []byte
	HasNASMessageContainer       bool
	NASMessageContainer          nasType.NASMessageContainer
}

var deregistrationUEOptions = nascodec.Schema{
	0x3c:                          {Format: nascodec.TLV, Min: 1, Max: 7},
	nasType.IENASMessageContainer: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *DeregistrationRequestUE) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil DeregistrationRequestUE")
	}
	w := nascodec.Writer{Data: []byte{nasType.EPD5GMM, 0, nasType.MsgDeregistrationRequestUE}}
	w.U8(m.DeregistrationTypeAndNgKSI, 255)
	w.LV(m.MobileIdentity, 2, 4, 65535)
	var options []nasType.IE
	if m.HasUnavailabilityInformation {
		options = append(options, nasType.IE{ID: 0x3c, Value: m.UnavailabilityInformation})
	}
	if m.HasNASMessageContainer {
		options = append(options, nasType.IE{ID: nasType.IENASMessageContainer, Value: m.NASMessageContainer.Value})
	}
	w.Options(options, deregistrationUEOptions)
	return w.Result()
}
func DecodeDeregistrationRequestUE(data []byte) (*DeregistrationRequestUE, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgDeregistrationRequestUE {
		return nil, fmt.Errorf("unexpected DeregistrationRequestUE header")
	}
	r := nascodec.Reader{Data: data[3:]}
	m := &DeregistrationRequestUE{DeregistrationTypeAndNgKSI: r.U8(255), MobileIdentity: r.LV(2, 4, 65535)}
	for _, ie := range r.Options(deregistrationUEOptions) {
		switch ie.ID {
		case 0x3c:
			m.HasUnavailabilityInformation = true
			m.UnavailabilityInformation = ie.Value
		case nasType.IENASMessageContainer:
			m.HasNASMessageContainer = true
			m.NASMessageContainer = nasType.NASMessageContainer{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode DeregistrationRequestUE: %w", r.Err)
	}
	return m, nil
}
