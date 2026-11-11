// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkRANStatusTransfer represents ASN.1 DownlinkRANStatusTransfer.
type DownlinkRANStatusTransfer struct {
	AMFUENGAPID                           ngapType.AMFUENGAPID                           `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                           ngapType.RANUENGAPID                           `ngap:"id:85,criticality:reject,presence:mandatory"`
	RANStatusTransferTransparentContainer ngapType.RANStatusTransferTransparentContainer `ngap:"id:84,criticality:reject,presence:mandatory"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkRANStatusTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkRANStatusTransfer", m)
}

// DecodeDownlinkRANStatusTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkRANStatusTransfer(data []byte) (*DownlinkRANStatusTransfer, error) {
	var m DownlinkRANStatusTransfer
	err := ngapcodec.Decode("DownlinkRANStatusTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
