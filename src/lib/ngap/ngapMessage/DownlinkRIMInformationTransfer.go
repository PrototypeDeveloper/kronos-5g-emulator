// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkRIMInformationTransfer represents ASN.1 DownlinkRIMInformationTransfer.
type DownlinkRIMInformationTransfer struct {
	HasRIMInformationTransfer bool
	RIMInformationTransfer    ngapType.RIMInformationTransfer `ngap:"id:175,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkRIMInformationTransfer) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkRIMInformationTransfer", m)
}

// DecodeDownlinkRIMInformationTransfer decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkRIMInformationTransfer(data []byte) (*DownlinkRIMInformationTransfer, error) {
	var m DownlinkRIMInformationTransfer
	err := ngapcodec.Decode("DownlinkRIMInformationTransfer", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
