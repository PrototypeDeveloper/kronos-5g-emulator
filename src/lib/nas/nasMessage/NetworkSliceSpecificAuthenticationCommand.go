package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// NetworkSliceSpecificAuthenticationCommand carries an EAP message for slice authentication.
type NetworkSliceSpecificAuthenticationCommand struct {
	SNSSAI     []byte
	EAPMessage []byte
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var networkSliceSpecificAuthenticationCommandOptions = nascodec.Schema{}

func (m *NetworkSliceSpecificAuthenticationCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil NetworkSliceSpecificAuthenticationCommand")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgNetworkSliceSpecificAuthenticationCommand}
	w.LV(m.SNSSAI, 1, 1, 8)
	w.LV(m.EAPMessage, 2, 4, 1500)
	w.UnknownOptions(m.UnknownIEs, networkSliceSpecificAuthenticationCommandOptions)
	return w.Result()
}
func DecodeNetworkSliceSpecificAuthenticationCommand(data []byte) (*NetworkSliceSpecificAuthenticationCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgNetworkSliceSpecificAuthenticationCommand {
		return nil, fmt.Errorf("unexpected header for NetworkSliceSpecificAuthenticationCommand")
	}
	m := &NetworkSliceSpecificAuthenticationCommand{}
	r := nascodec.Reader{Data: data[3:]}
	m.SNSSAI = r.LV(1, 1, 8)
	m.EAPMessage = r.LV(2, 4, 1500)
	_, unknown := r.OptionsWithUnknown(networkSliceSpecificAuthenticationCommandOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode NetworkSliceSpecificAuthenticationCommand: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
