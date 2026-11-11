// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceReleaseResponse represents ASN.1 PDUSessionResourceReleaseResponse.
type PDUSessionResourceReleaseResponse struct {
	AMFUENGAPID                          ngapType.AMFUENGAPID                          `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                          ngapType.RANUENGAPID                          `ngap:"id:85,criticality:ignore,presence:mandatory"`
	PDUSessionResourceReleasedListRelRes ngapType.PDUSessionResourceReleasedListRelRes `ngap:"id:70,criticality:ignore,presence:mandatory"`
	HasUserLocationInformation           bool
	UserLocationInformation              ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics            bool
	CriticalityDiagnostics               ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceReleaseResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceReleaseResponse", m)
}

// DecodePDUSessionResourceReleaseResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceReleaseResponse(data []byte) (*PDUSessionResourceReleaseResponse, error) {
	var m PDUSessionResourceReleaseResponse
	err := ngapcodec.Decode("PDUSessionResourceReleaseResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
