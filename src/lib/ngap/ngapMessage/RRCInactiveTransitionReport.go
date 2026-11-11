// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RRCInactiveTransitionReport represents ASN.1 RRCInactiveTransitionReport.
type RRCInactiveTransitionReport struct {
	AMFUENGAPID             ngapType.AMFUENGAPID             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID             ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	RRCState                ngapType.RRCState                `ngap:"id:92,criticality:ignore,presence:mandatory"`
	UserLocationInformation ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:mandatory"`
	AdditionalIEs           []ngapType.UnknownIE
	UnknownExtensions       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RRCInactiveTransitionReport) Encode() ([]byte, error) {
	return ngapcodec.Encode("RRCInactiveTransitionReport", m)
}

// DecodeRRCInactiveTransitionReport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRRCInactiveTransitionReport(data []byte) (*RRCInactiveTransitionReport, error) {
	var m RRCInactiveTransitionReport
	err := ngapcodec.Decode("RRCInactiveTransitionReport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
