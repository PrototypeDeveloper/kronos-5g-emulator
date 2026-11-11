// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UERadioCapabilityInfoIndication represents ASN.1 UERadioCapabilityInfoIndication.
type UERadioCapabilityInfoIndication struct {
	AMFUENGAPID                     ngapType.AMFUENGAPID       `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                     ngapType.RANUENGAPID       `ngap:"id:85,criticality:reject,presence:mandatory"`
	UERadioCapability               ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:mandatory"`
	HasUERadioCapabilityForPaging   bool
	UERadioCapabilityForPaging      ngapType.UERadioCapabilityForPaging `ngap:"id:118,criticality:ignore,presence:optional"`
	HasUERadioCapabilityEUTRAFormat bool
	UERadioCapabilityEUTRAFormat    ngapType.UERadioCapability `ngap:"id:265,criticality:ignore,presence:optional"`
	HasXrDeviceWith2Rx              bool
	XrDeviceWith2Rx                 ngapType.XrDeviceWith2Rx `ngap:"id:428,criticality:ignore,presence:optional"`
	AdditionalIEs                   []ngapType.UnknownIE
	UnknownExtensions               []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UERadioCapabilityInfoIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("UERadioCapabilityInfoIndication", m)
}

// DecodeUERadioCapabilityInfoIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUERadioCapabilityInfoIndication(data []byte) (*UERadioCapabilityInfoIndication, error) {
	var m UERadioCapabilityInfoIndication
	err := ngapcodec.Decode("UERadioCapabilityInfoIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
