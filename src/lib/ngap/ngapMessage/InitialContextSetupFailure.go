// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InitialContextSetupFailure represents ASN.1 InitialContextSetupFailure.
type InitialContextSetupFailure struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceFailedToSetupListCxtFail bool
	PDUSessionResourceFailedToSetupListCxtFail    ngapType.PDUSessionResourceFailedToSetupListCxtFail `ngap:"id:132,criticality:ignore,presence:optional"`
	Cause                                         ngapType.Cause                                      `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics                     bool
	CriticalityDiagnostics                        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InitialContextSetupFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("InitialContextSetupFailure", m)
}

// DecodeInitialContextSetupFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInitialContextSetupFailure(data []byte) (*InitialContextSetupFailure, error) {
	var m InitialContextSetupFailure
	err := ngapcodec.Decode("InitialContextSetupFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
