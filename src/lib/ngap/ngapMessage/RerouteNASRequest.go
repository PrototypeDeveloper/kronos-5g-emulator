// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RerouteNASRequest represents ASN.1 RerouteNASRequest.
type RerouteNASRequest struct {
	RANUENGAPID                            ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasAMFUENGAPID                         bool
	AMFUENGAPID                            ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:optional"`
	NGAPMessage                            ngapType.NGAPMessage `ngap:"id:42,criticality:reject,presence:mandatory"`
	AMFSetID                               ngapType.AMFSetID    `ngap:"id:3,criticality:reject,presence:mandatory"`
	HasAllowedNSSAI                        bool
	AllowedNSSAI                           ngapType.AllowedNSSAI `ngap:"id:0,criticality:reject,presence:optional"`
	HasSourceToTargetAMFInformationReroute bool
	SourceToTargetAMFInformationReroute    ngapType.SourceToTargetAMFInformationReroute `ngap:"id:171,criticality:ignore,presence:optional"`
	HasPartiallyAllowedNSSAI               bool
	PartiallyAllowedNSSAI                  ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	AdditionalIEs                          []ngapType.UnknownIE
	UnknownExtensions                      []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RerouteNASRequest) Encode() ([]byte, error) { return ngapcodec.Encode("RerouteNASRequest", m) }

// DecodeRerouteNASRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRerouteNASRequest(data []byte) (*RerouteNASRequest, error) {
	var m RerouteNASRequest
	err := ngapcodec.Decode("RerouteNASRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
