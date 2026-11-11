// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// WriteReplaceWarningResponse represents ASN.1 WriteReplaceWarningResponse.
type WriteReplaceWarningResponse struct {
	MessageIdentifier             ngapType.MessageIdentifier `ngap:"id:35,criticality:reject,presence:mandatory"`
	SerialNumber                  ngapType.SerialNumber      `ngap:"id:95,criticality:reject,presence:mandatory"`
	HasBroadcastCompletedAreaList bool
	BroadcastCompletedAreaList    ngapType.BroadcastCompletedAreaList `ngap:"id:13,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics     bool
	CriticalityDiagnostics        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                 []ngapType.UnknownIE
	UnknownExtensions             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *WriteReplaceWarningResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("WriteReplaceWarningResponse", m)
}

// DecodeWriteReplaceWarningResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeWriteReplaceWarningResponse(data []byte) (*WriteReplaceWarningResponse, error) {
	var m WriteReplaceWarningResponse
	err := ngapcodec.Decode("WriteReplaceWarningResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
