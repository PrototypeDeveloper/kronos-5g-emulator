// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkNASTransport represents ASN.1 UplinkNASTransport.
type UplinkNASTransport struct {
	AMFUENGAPID                ngapType.AMFUENGAPID             `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	NASPDU                     ngapType.NASPDU                  `ngap:"id:38,criticality:reject,presence:mandatory"`
	UserLocationInformation    ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:mandatory"`
	HasWAGFIdentityInformation bool
	WAGFIdentityInformation    ngapType.WAGFIdentityInformation `ngap:"id:239,criticality:reject,presence:optional"`
	HasTNGFIdentityInformation bool
	TNGFIdentityInformation    ngapType.TNGFIdentityInformation `ngap:"id:246,criticality:reject,presence:optional"`
	HasTWIFIdentityInformation bool
	TWIFIdentityInformation    ngapType.TWIFIdentityInformation `ngap:"id:247,criticality:reject,presence:optional"`
	AdditionalIEs              []ngapType.UnknownIE
	UnknownExtensions          []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkNASTransport) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkNASTransport", m)
}

// DecodeUplinkNASTransport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkNASTransport(data []byte) (*UplinkNASTransport, error) {
	var m UplinkNASTransport
	err := ngapcodec.Decode("UplinkNASTransport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
