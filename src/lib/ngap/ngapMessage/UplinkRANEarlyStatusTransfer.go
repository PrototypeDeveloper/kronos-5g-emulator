// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkRANEarlyStatusTransfer represents ASN.1 UplinkRANEarlyStatusTransfer.
type UplinkRANEarlyStatusTransfer struct {
	AMFUENGAPID                             ngapType.AMFUENGAPID                             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                             ngapType.RANUENGAPID                             `ngap:"id:85,criticality:reject,presence:mandatory"`
	EarlyStatusTransferTransparentContainer ngapType.EarlyStatusTransferTransparentContainer `ngap:"id:268,criticality:reject,presence:mandatory"`
	AdditionalIEs                           []ngapType.UnknownIE
	UnknownExtensions                       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkRANEarlyStatusTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkRANEarlyStatusTransfer", m)
}

// DecodeUplinkRANEarlyStatusTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkRANEarlyStatusTransfer(data []byte) (*UplinkRANEarlyStatusTransfer, error) {
	var m UplinkRANEarlyStatusTransfer
	err := ngapcodec.Decode("UplinkRANEarlyStatusTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
