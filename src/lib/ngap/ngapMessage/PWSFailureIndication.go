// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PWSFailureIndication represents ASN.1 PWSFailureIndication.
type PWSFailureIndication struct {
	PWSFailedCellIDList ngapType.PWSFailedCellIDList `ngap:"id:81,criticality:reject,presence:mandatory"`
	GlobalRANNodeID     ngapType.GlobalRANNodeID     `ngap:"id:27,criticality:reject,presence:mandatory"`
	AdditionalIEs       []ngapType.UnknownIE
	UnknownExtensions   []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PWSFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("PWSFailureIndication", m)
}

// DecodePWSFailureIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePWSFailureIndication(data []byte) (*PWSFailureIndication, error) {
	var m PWSFailureIndication
	err := ngapcodec.Decode("PWSFailureIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
