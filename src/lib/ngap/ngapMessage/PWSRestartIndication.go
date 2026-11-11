// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PWSRestartIndication represents ASN.1 PWSRestartIndication.
type PWSRestartIndication struct {
	CellIDListForRestart             ngapType.CellIDListForRestart `ngap:"id:16,criticality:reject,presence:mandatory"`
	GlobalRANNodeID                  ngapType.GlobalRANNodeID      `ngap:"id:27,criticality:reject,presence:mandatory"`
	TAIListForRestart                ngapType.TAIListForRestart    `ngap:"id:104,criticality:reject,presence:mandatory"`
	HasEmergencyAreaIDListForRestart bool
	EmergencyAreaIDListForRestart    ngapType.EmergencyAreaIDListForRestart `ngap:"id:23,criticality:reject,presence:optional"`
	AdditionalIEs                    []ngapType.UnknownIE
	UnknownExtensions                []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PWSRestartIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("PWSRestartIndication", m)
}

// DecodePWSRestartIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePWSRestartIndication(data []byte) (*PWSRestartIndication, error) {
	var m PWSRestartIndication
	err := ngapcodec.Decode("PWSRestartIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
