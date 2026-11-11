// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// OverloadStop represents ASN.1 OverloadStop.
type OverloadStop struct {
	HasGUAMIList      bool
	GUAMIList         ngapType.GUAMIList `ngap:"id:451,criticality:ignore,presence:optional"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *OverloadStop) Encode() ([]byte, error) { return ngapcodec.Encode("OverloadStop", m) }

// DecodeOverloadStop decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeOverloadStop(data []byte) (*OverloadStop, error) {
	var m OverloadStop
	err := ngapcodec.Decode("OverloadStop", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
