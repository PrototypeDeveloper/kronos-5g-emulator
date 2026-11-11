package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// NetworkSliceSpecificAuthenticationResult carries the EAP result of slice authentication.
type NetworkSliceSpecificAuthenticationResult struct {
	SNSSAI     []byte
	EAPMessage []byte
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var networkSliceSpecificAuthenticationResultOptions = nascodec.Schema{}

func (m *NetworkSliceSpecificAuthenticationResult) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil NetworkSliceSpecificAuthenticationResult")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgNetworkSliceSpecificAuthenticationResult}
	w.LV(m.SNSSAI, 1, 1, 8)
	w.LV(m.EAPMessage, 2, 4, 1500)
	w.UnknownOptions(m.UnknownIEs, networkSliceSpecificAuthenticationResultOptions)
	return w.Result()
}
func DecodeNetworkSliceSpecificAuthenticationResult(data []byte) (*NetworkSliceSpecificAuthenticationResult, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgNetworkSliceSpecificAuthenticationResult {
		return nil, fmt.Errorf("unexpected header for NetworkSliceSpecificAuthenticationResult")
	}
	m := &NetworkSliceSpecificAuthenticationResult{}
	r := nascodec.Reader{Data: data[3:]}
	m.SNSSAI = r.LV(1, 1, 8)
	m.EAPMessage = r.LV(2, 4, 1500)
	_, unknown := r.OptionsWithUnknown(networkSliceSpecificAuthenticationResultOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode NetworkSliceSpecificAuthenticationResult: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
