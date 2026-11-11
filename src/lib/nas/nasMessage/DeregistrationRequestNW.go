package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// DeregistrationRequestNW is the UE-terminated (network-initiated) de-registration request.
type DeregistrationRequestNW struct {
	DeregistrationType    uint8 // low nibble; the high nibble is spare
	HasCause5GMM          bool
	Cause5GMM             nasType.Cause5GMM
	HasT3346Value         bool
	T3346Value            nasType.T3346Value
	HasRejectedNSSAI      bool
	RejectedNSSAI         nasType.RejectedNSSAI
	HasCAGInformationList bool
	CAGInformationList    nasType.CAGInformationList
	UnknownIEs            []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var deregistrationRequestNWOptions = nascodec.Schema{
	nasType.IECause5GMM:                   {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IET3346Value:                  {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEDeregistrationRejectedNSSAI: {Format: nascodec.TLV, Min: 2, Max: 40},
	nasType.IECAGInformationList:          {Format: nascodec.TLVE, Min: 0, Max: 65535},
}

func (m *DeregistrationRequestNW) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil DeregistrationRequestNW")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgDeregistrationRequestNW}
	w.U8(m.DeregistrationType, 15)
	var options []nasType.IE
	if m.HasCause5GMM {
		options = append(options, nasType.IE{ID: nasType.IECause5GMM, Value: []byte{m.Cause5GMM.Value}})
	}
	if m.HasT3346Value {
		options = append(options, nasType.IE{ID: nasType.IET3346Value, Value: []byte{m.T3346Value.Value}})
	}
	if m.HasRejectedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IEDeregistrationRejectedNSSAI, Value: m.RejectedNSSAI.Value})
	}
	if m.HasCAGInformationList {
		options = append(options, nasType.IE{ID: nasType.IECAGInformationList, Value: m.CAGInformationList.Value})
	}
	w.Options(options, deregistrationRequestNWOptions)
	w.UnknownOptions(m.UnknownIEs, deregistrationRequestNWOptions)
	return w.Result()
}
func DecodeDeregistrationRequestNW(data []byte) (*DeregistrationRequestNW, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgDeregistrationRequestNW {
		return nil, fmt.Errorf("unexpected header for DeregistrationRequestNW")
	}
	m := &DeregistrationRequestNW{}
	r := nascodec.Reader{Data: data[3:]}
	m.DeregistrationType = r.U8(15)
	known, unknown := r.OptionsWithUnknown(deregistrationRequestNWOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode DeregistrationRequestNW: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IECause5GMM:
			m.HasCause5GMM = true
			m.Cause5GMM = nasType.Cause5GMM{Value: ie.Value[0]}
		case nasType.IET3346Value:
			m.HasT3346Value = true
			m.T3346Value = nasType.T3346Value{Value: ie.Value[0]}
		case nasType.IEDeregistrationRejectedNSSAI:
			m.HasRejectedNSSAI = true
			m.RejectedNSSAI = nasType.RejectedNSSAI{Value: ie.Value}
		case nasType.IECAGInformationList:
			m.HasCAGInformationList = true
			m.CAGInformationList = nasType.CAGInformationList{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
