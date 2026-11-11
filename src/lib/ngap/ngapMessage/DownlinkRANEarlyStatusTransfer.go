// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkRANEarlyStatusTransfer represents ASN.1 DownlinkRANEarlyStatusTransfer.
type DownlinkRANEarlyStatusTransfer struct {
	AMFUENGAPID                             ngapType.AMFUENGAPID                             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                             ngapType.RANUENGAPID                             `ngap:"id:85,criticality:reject,presence:mandatory"`
	EarlyStatusTransferTransparentContainer ngapType.EarlyStatusTransferTransparentContainer `ngap:"id:268,criticality:reject,presence:mandatory"`
	AdditionalIEs                           []ngapType.UnknownIE
	UnknownExtensions                       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkRANEarlyStatusTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkRANEarlyStatusTransfer", m)
}

// DecodeDownlinkRANEarlyStatusTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkRANEarlyStatusTransfer(data []byte) (*DownlinkRANEarlyStatusTransfer, error) {
	var m DownlinkRANEarlyStatusTransfer
	err := ngapcodec.Decode("DownlinkRANEarlyStatusTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
