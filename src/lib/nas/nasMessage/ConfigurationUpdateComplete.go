package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ConfigurationUpdateComplete acknowledges a configuration update command.
type ConfigurationUpdateComplete struct {
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var configurationUpdateCompleteOptions = nascodec.Schema{}

func (m *ConfigurationUpdateComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ConfigurationUpdateComplete")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgConfigurationUpdateComplete}
	w.UnknownOptions(m.UnknownIEs, configurationUpdateCompleteOptions)
	return w.Result()
}
func DecodeConfigurationUpdateComplete(data []byte) (*ConfigurationUpdateComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgConfigurationUpdateComplete {
		return nil, fmt.Errorf("unexpected header for ConfigurationUpdateComplete")
	}
	m := &ConfigurationUpdateComplete{}
	r := nascodec.Reader{Data: data[3:]}
	_, unknown := r.OptionsWithUnknown(configurationUpdateCompleteOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode ConfigurationUpdateComplete: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
