// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEInformationTransfer represents ASN.1 UEInformationTransfer.
type UEInformationTransfer struct {
	FiveGSTMSI               ngapType.FiveGSTMSI `ngap:"id:26,criticality:reject,presence:mandatory"`
	HasNBIoTUEPriority       bool
	NBIoTUEPriority          ngapType.NBIoTUEPriority `ngap:"id:210,criticality:ignore,presence:optional"`
	HasUERadioCapability     bool
	UERadioCapability        ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:optional"`
	HasSNSSAI                bool
	SNSSAI                   ngapType.SNSSAI `ngap:"id:148,criticality:ignore,presence:optional"`
	HasAllowedNSSAI          bool
	AllowedNSSAI             ngapType.AllowedNSSAI `ngap:"id:0,criticality:ignore,presence:optional"`
	HasUEDifferentiationInfo bool
	UEDifferentiationInfo    ngapType.UEDifferentiationInfo `ngap:"id:209,criticality:ignore,presence:optional"`
	HasMaskedIMEISV          bool
	MaskedIMEISV             ngapType.MaskedIMEISV `ngap:"id:34,criticality:ignore,presence:optional"`
	HasPartiallyAllowedNSSAI bool
	PartiallyAllowedNSSAI    ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	AdditionalIEs            []ngapType.UnknownIE
	UnknownExtensions        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEInformationTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEInformationTransfer", m)
}

// DecodeUEInformationTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEInformationTransfer(data []byte) (*UEInformationTransfer, error) {
	var m UEInformationTransfer
	err := ngapcodec.Decode("UEInformationTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
