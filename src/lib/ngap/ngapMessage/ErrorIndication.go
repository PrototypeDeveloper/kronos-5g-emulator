// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// ErrorIndication represents ASN.1 ErrorIndication.
type ErrorIndication struct {
	HasAMFUENGAPID            bool
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:optional"`
	HasRANUENGAPID            bool
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:optional"`
	HasCause                  bool
	Cause                     ngapType.Cause `ngap:"id:15,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasFiveGSTMSI             bool
	FiveGSTMSI                ngapType.FiveGSTMSI `ngap:"id:26,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *ErrorIndication) Encode() ([]byte, error) { return ngapcodec.Encode("ErrorIndication", m) }

// DecodeErrorIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeErrorIndication(data []byte) (*ErrorIndication, error) {
	var m ErrorIndication
	err := ngapcodec.Decode("ErrorIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
