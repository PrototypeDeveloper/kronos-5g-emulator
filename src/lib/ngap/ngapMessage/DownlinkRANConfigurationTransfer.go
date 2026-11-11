// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkRANConfigurationTransfer represents ASN.1 DownlinkRANConfigurationTransfer.
type DownlinkRANConfigurationTransfer struct {
	HasSONConfigurationTransferDL            bool
	SONConfigurationTransferDL               ngapType.SONConfigurationTransfer `ngap:"id:98,criticality:ignore,presence:optional"`
	HasENDCSONConfigurationTransferDL        bool
	ENDCSONConfigurationTransferDL           ngapType.ENDCSONConfigurationTransfer `ngap:"id:157,criticality:ignore,presence:optional"`
	HasIntersystemSONConfigurationTransferDL bool
	IntersystemSONConfigurationTransferDL    ngapType.IntersystemSONConfigurationTransfer `ngap:"id:250,criticality:ignore,presence:optional"`
	AdditionalIEs                            []ngapType.UnknownIE
	UnknownExtensions                        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkRANConfigurationTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkRANConfigurationTransfer", m)
}

// DecodeDownlinkRANConfigurationTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkRANConfigurationTransfer(data []byte) (*DownlinkRANConfigurationTransfer, error) {
	var m DownlinkRANConfigurationTransfer
	err := ngapcodec.Decode("DownlinkRANConfigurationTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
