// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceModifyResponse represents ASN.1 PDUSessionResourceModifyResponse.
type PDUSessionResourceModifyResponse struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceModifyListModRes         bool
	PDUSessionResourceModifyListModRes            ngapType.PDUSessionResourceModifyListModRes `ngap:"id:65,criticality:ignore,presence:optional"`
	HasPDUSessionResourceFailedToModifyListModRes bool
	PDUSessionResourceFailedToModifyListModRes    ngapType.PDUSessionResourceFailedToModifyListModRes `ngap:"id:54,criticality:ignore,presence:optional"`
	HasUserLocationInformation                    bool
	UserLocationInformation                       ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics                     bool
	CriticalityDiagnostics                        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceModifyResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceModifyResponse", m)
}

// DecodePDUSessionResourceModifyResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceModifyResponse(data []byte) (*PDUSessionResourceModifyResponse, error) {
	var m PDUSessionResourceModifyResponse
	err := ngapcodec.Decode("PDUSessionResourceModifyResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
