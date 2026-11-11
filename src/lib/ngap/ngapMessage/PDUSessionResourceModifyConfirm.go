// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceModifyConfirm represents ASN.1 PDUSessionResourceModifyConfirm.
type PDUSessionResourceModifyConfirm struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceModifyListModCfm         bool
	PDUSessionResourceModifyListModCfm            ngapType.PDUSessionResourceModifyListModCfm `ngap:"id:62,criticality:ignore,presence:optional"`
	HasPDUSessionResourceFailedToModifyListModCfm bool
	PDUSessionResourceFailedToModifyListModCfm    ngapType.PDUSessionResourceFailedToModifyListModCfm `ngap:"id:131,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics                     bool
	CriticalityDiagnostics                        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceModifyConfirm) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceModifyConfirm", m)
}

// DecodePDUSessionResourceModifyConfirm decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceModifyConfirm(data []byte) (*PDUSessionResourceModifyConfirm, error) {
	var m PDUSessionResourceModifyConfirm
	err := ngapcodec.Decode("PDUSessionResourceModifyConfirm", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
