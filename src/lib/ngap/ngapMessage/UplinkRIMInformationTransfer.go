// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkRIMInformationTransfer represents ASN.1 UplinkRIMInformationTransfer.
type UplinkRIMInformationTransfer struct {
	HasRIMInformationTransfer bool
	RIMInformationTransfer    ngapType.RIMInformationTransfer `ngap:"id:175,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkRIMInformationTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkRIMInformationTransfer", m)
}

// DecodeUplinkRIMInformationTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkRIMInformationTransfer(data []byte) (*UplinkRIMInformationTransfer, error) {
	var m UplinkRIMInformationTransfer
	err := ngapcodec.Decode("UplinkRIMInformationTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
