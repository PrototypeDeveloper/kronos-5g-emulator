// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NASNonDeliveryIndication represents ASN.1 NASNonDeliveryIndication.
type NASNonDeliveryIndication struct {
	AMFUENGAPID       ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	NASPDU            ngapType.NASPDU      `ngap:"id:38,criticality:ignore,presence:mandatory"`
	Cause             ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NASNonDeliveryIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("NASNonDeliveryIndication", m)
}

// DecodeNASNonDeliveryIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNASNonDeliveryIndication(data []byte) (*NASNonDeliveryIndication, error) {
	var m NASNonDeliveryIndication
	err := ngapcodec.Decode("NASNonDeliveryIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
