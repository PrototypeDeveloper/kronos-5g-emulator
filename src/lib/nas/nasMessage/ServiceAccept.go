package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

type ServiceAccept struct {
	HasPDUSessionStatus   bool
	PDUSessionStatus      []byte
	HasReactivationResult bool
	ReactivationResult    []byte
}

var serviceAcceptOptions = nascodec.Schema{0x50: {Format: nascodec.TLV, Min: 2, Max: 32}, 0x26: {Format: nascodec.TLV, Min: 2, Max: 32}}

func (m *ServiceAccept) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ServiceAccept")
	}
	w := nascodec.Writer{Data: []byte{nasType.EPD5GMM, 0, nasType.MsgServiceAccept}}
	var opts []nasType.IE
	if m.HasPDUSessionStatus {
		opts = append(opts, nasType.IE{ID: 0x50, Value: m.PDUSessionStatus})
	}
	if m.HasReactivationResult {
		opts = append(opts, nasType.IE{ID: 0x26, Value: m.ReactivationResult})
	}
	w.Options(opts, serviceAcceptOptions)
	return w.Result()
}
func DecodeServiceAccept(data []byte) (*ServiceAccept, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgServiceAccept {
		return nil, fmt.Errorf("unexpected ServiceAccept header")
	}
	r := nascodec.Reader{Data: data[3:]}
	m := &ServiceAccept{}
	for _, ie := range r.Options(serviceAcceptOptions) {
		switch ie.ID {
		case 0x50:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = ie.Value
		case 0x26:
			m.HasReactivationResult = true
			m.ReactivationResult = ie.Value
		}
	}
	return m, r.Err
}
