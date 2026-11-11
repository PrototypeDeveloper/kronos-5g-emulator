// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkRANStatusTransfer represents ASN.1 UplinkRANStatusTransfer.
type UplinkRANStatusTransfer struct {
	AMFUENGAPID                           ngapType.AMFUENGAPID                           `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                           ngapType.RANUENGAPID                           `ngap:"id:85,criticality:reject,presence:mandatory"`
	RANStatusTransferTransparentContainer ngapType.RANStatusTransferTransparentContainer `ngap:"id:84,criticality:reject,presence:mandatory"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkRANStatusTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkRANStatusTransfer", m)
}

// DecodeUplinkRANStatusTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkRANStatusTransfer(data []byte) (*UplinkRANStatusTransfer, error) {
	var m UplinkRANStatusTransfer
	err := ngapcodec.Decode("UplinkRANStatusTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
