// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkRANConfigurationTransfer represents ASN.1 UplinkRANConfigurationTransfer.
type UplinkRANConfigurationTransfer struct {
	HasSONConfigurationTransferUL            bool
	SONConfigurationTransferUL               ngapType.SONConfigurationTransfer `ngap:"id:99,criticality:ignore,presence:optional"`
	HasENDCSONConfigurationTransferUL        bool
	ENDCSONConfigurationTransferUL           ngapType.ENDCSONConfigurationTransfer `ngap:"id:158,criticality:ignore,presence:optional"`
	HasIntersystemSONConfigurationTransferUL bool
	IntersystemSONConfigurationTransferUL    ngapType.IntersystemSONConfigurationTransfer `ngap:"id:251,criticality:ignore,presence:optional"`
	AdditionalIEs                            []ngapType.UnknownIE
	UnknownExtensions                        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkRANConfigurationTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkRANConfigurationTransfer", m)
}

// DecodeUplinkRANConfigurationTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkRANConfigurationTransfer(data []byte) (*UplinkRANConfigurationTransfer, error) {
	var m UplinkRANConfigurationTransfer
	err := ngapcodec.Decode("UplinkRANConfigurationTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
