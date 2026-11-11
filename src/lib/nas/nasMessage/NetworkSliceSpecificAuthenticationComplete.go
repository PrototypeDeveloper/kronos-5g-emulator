package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// NetworkSliceSpecificAuthenticationComplete carries the UE's EAP message for slice authentication.
type NetworkSliceSpecificAuthenticationComplete struct {
	SNSSAI     []byte
	EAPMessage []byte
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var networkSliceSpecificAuthenticationCompleteOptions = nascodec.Schema{}

func (m *NetworkSliceSpecificAuthenticationComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil NetworkSliceSpecificAuthenticationComplete")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgNetworkSliceSpecificAuthenticationComplete}
	w.LV(m.SNSSAI, 1, 1, 8)
	w.LV(m.EAPMessage, 2, 4, 1500)
	w.UnknownOptions(m.UnknownIEs, networkSliceSpecificAuthenticationCompleteOptions)
	return w.Result()
}
func DecodeNetworkSliceSpecificAuthenticationComplete(data []byte) (*NetworkSliceSpecificAuthenticationComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgNetworkSliceSpecificAuthenticationComplete {
		return nil, fmt.Errorf("unexpected header for NetworkSliceSpecificAuthenticationComplete")
	}
	m := &NetworkSliceSpecificAuthenticationComplete{}
	r := nascodec.Reader{Data: data[3:]}
	m.SNSSAI = r.LV(1, 1, 8)
	m.EAPMessage = r.LV(2, 4, 1500)
	_, unknown := r.OptionsWithUnknown(networkSliceSpecificAuthenticationCompleteOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode NetworkSliceSpecificAuthenticationComplete: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
