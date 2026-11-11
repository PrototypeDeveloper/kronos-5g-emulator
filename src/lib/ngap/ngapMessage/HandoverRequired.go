// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverRequired represents ASN.1 HandoverRequired.
type HandoverRequired struct {
	AMFUENGAPID                         ngapType.AMFUENGAPID  `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                         ngapType.RANUENGAPID  `ngap:"id:85,criticality:reject,presence:mandatory"`
	HandoverType                        ngapType.HandoverType `ngap:"id:29,criticality:reject,presence:mandatory"`
	Cause                               ngapType.Cause        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	TargetID                            ngapType.TargetID     `ngap:"id:105,criticality:reject,presence:mandatory"`
	HasDirectForwardingPathAvailability bool
	DirectForwardingPathAvailability    ngapType.DirectForwardingPathAvailability   `ngap:"id:22,criticality:ignore,presence:optional"`
	PDUSessionResourceListHORqd         ngapType.PDUSessionResourceListHORqd        `ngap:"id:61,criticality:reject,presence:mandatory"`
	SourceToTargetTransparentContainer  ngapType.SourceToTargetTransparentContainer `ngap:"id:101,criticality:reject,presence:mandatory"`
	AdditionalIEs                       []ngapType.UnknownIE
	UnknownExtensions                   []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverRequired) Encode() ([]byte, error) { return ngapcodec.Encode("HandoverRequired", m) }

// DecodeHandoverRequired decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverRequired(data []byte) (*HandoverRequired, error) {
	var m HandoverRequired
	err := ngapcodec.Decode("HandoverRequired", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
