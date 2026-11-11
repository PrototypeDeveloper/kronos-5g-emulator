// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AIOTSessionReleaseComplete represents ASN.1 AIOTSessionReleaseComplete.
type AIOTSessionReleaseComplete struct {
	AIOTFIdentifier                    ngapType.AIOTFIdentifier                               `ngap:"id:467,criticality:ignore,presence:mandatory"`
	AIoTCorrelationIdentifier          ngapType.AIoTCorrelationIdentifier                     `ngap:"id:466,criticality:ignore,presence:mandatory"`
	AIOTSessionReleaseCompleteTransfer ngapType.AIOTSessionReleaseCompleteTransferOCTETSTRING `ngap:"id:464,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics          bool
	CriticalityDiagnostics             ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AIOTSessionReleaseComplete) Encode() ([]byte, error) {
	return ngapcodec.Encode("AIOTSessionReleaseComplete", m)
}

// DecodeAIOTSessionReleaseComplete decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAIOTSessionReleaseComplete(data []byte) (*AIOTSessionReleaseComplete, error) {
	var m AIOTSessionReleaseComplete
	err := ngapcodec.Decode("AIOTSessionReleaseComplete", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
