package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ServiceRequest is TS 24.501 8.2.16. Service type is the high nibble.
type ServiceRequest struct {
	ServiceTypeAndNgKSI        uint8
	FiveGSTMSI                 []byte
	HasUplinkDataStatus        bool
	UplinkDataStatus           []byte
	HasPDUSessionStatus        bool
	PDUSessionStatus           []byte
	HasAllowedPDUSessionStatus bool
	AllowedPDUSessionStatus    []byte
	HasNASMessageContainer     bool
	NASMessageContainer        []byte
}

var serviceRequestOptions = nascodec.Schema{
	0x40: {Format: nascodec.TLV, Min: 2, Max: 32},
	0x50: {Format: nascodec.TLV, Min: 2, Max: 32},
	0x25: {Format: nascodec.TLV, Min: 2, Max: 32},
	0x71: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *ServiceRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ServiceRequest")
	}
	w := nascodec.Writer{Data: []byte{nasType.EPD5GMM, 0, nasType.MsgServiceRequest}}
	w.U8(m.ServiceTypeAndNgKSI, 255)
	w.LV(m.FiveGSTMSI, 2, 7, 7)
	var opts []nasType.IE
	if m.HasUplinkDataStatus {
		opts = append(opts, nasType.IE{ID: 0x40, Value: m.UplinkDataStatus})
	}
	if m.HasPDUSessionStatus {
		opts = append(opts, nasType.IE{ID: 0x50, Value: m.PDUSessionStatus})
	}
	if m.HasAllowedPDUSessionStatus {
		opts = append(opts, nasType.IE{ID: 0x25, Value: m.AllowedPDUSessionStatus})
	}
	if m.HasNASMessageContainer {
		opts = append(opts, nasType.IE{ID: 0x71, Value: m.NASMessageContainer})
	}
	w.Options(opts, serviceRequestOptions)
	return w.Result()
}
func DecodeServiceRequest(data []byte) (*ServiceRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgServiceRequest {
		return nil, fmt.Errorf("unexpected ServiceRequest header")
	}
	r := nascodec.Reader{Data: data[3:]}
	m := &ServiceRequest{ServiceTypeAndNgKSI: r.U8(255), FiveGSTMSI: r.LV(2, 7, 7)}
	for _, ie := range r.Options(serviceRequestOptions) {
		switch ie.ID {
		case 0x40:
			m.HasUplinkDataStatus = true
			m.UplinkDataStatus = ie.Value
		case 0x50:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = ie.Value
		case 0x25:
			m.HasAllowedPDUSessionStatus = true
			m.AllowedPDUSessionStatus = ie.Value
		case 0x71:
			m.HasNASMessageContainer = true
			m.NASMessageContainer = ie.Value
		}
	}
	return m, r.Err
}
