// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverNotify represents ASN.1 HandoverNotify.
type HandoverNotify struct {
	AMFUENGAPID              ngapType.AMFUENGAPID             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID              ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	UserLocationInformation  ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:mandatory"`
	HasNotifySourceNGRANNode bool
	NotifySourceNGRANNode    ngapType.NotifySourceNGRANNode `ngap:"id:269,criticality:ignore,presence:optional"`
	AdditionalIEs            []ngapType.UnknownIE
	UnknownExtensions        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverNotify) Encode() ([]byte, error) { return ngapcodec.Encode("HandoverNotify", m) }

// DecodeHandoverNotify decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverNotify(data []byte) (*HandoverNotify, error) {
	var m HandoverNotify
	err := ngapcodec.Decode("HandoverNotify", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
