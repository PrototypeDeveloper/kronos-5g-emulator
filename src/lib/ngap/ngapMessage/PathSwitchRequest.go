// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PathSwitchRequest represents ASN.1 PathSwitchRequest.
type PathSwitchRequest struct {
	RANUENGAPID                                 ngapType.RANUENGAPID                          `ngap:"id:85,criticality:reject,presence:mandatory"`
	SourceAMFUENGAPID                           ngapType.AMFUENGAPID                          `ngap:"id:100,criticality:reject,presence:mandatory"`
	UserLocationInformation                     ngapType.UserLocationInformation              `ngap:"id:121,criticality:ignore,presence:mandatory"`
	UESecurityCapabilities                      ngapType.UESecurityCapabilities               `ngap:"id:119,criticality:ignore,presence:mandatory"`
	PDUSessionResourceToBeSwitchedDLList        ngapType.PDUSessionResourceToBeSwitchedDLList `ngap:"id:76,criticality:reject,presence:mandatory"`
	HasPDUSessionResourceFailedToSetupListPSReq bool
	PDUSessionResourceFailedToSetupListPSReq    ngapType.PDUSessionResourceFailedToSetupListPSReq `ngap:"id:57,criticality:ignore,presence:optional"`
	HasRRCResumeCause                           bool
	RRCResumeCause                              ngapType.RRCEstablishmentCause `ngap:"id:237,criticality:ignore,presence:optional"`
	HasRedCapIndication                         bool
	RedCapIndication                            ngapType.RedCapIndication `ngap:"id:333,criticality:ignore,presence:optional"`
	HasERedCapIndication                        bool
	ERedCapIndication                           ngapType.ERedCapIndication `ngap:"id:427,criticality:ignore,presence:optional"`
	HasSourceAMFGUAMI                           bool
	SourceAMFGUAMI                              ngapType.GUAMI `ngap:"id:453,criticality:ignore,presence:optional"`
	AdditionalIEs                               []ngapType.UnknownIE
	UnknownExtensions                           []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PathSwitchRequest) Encode() ([]byte, error) { return ngapcodec.Encode("PathSwitchRequest", m) }

// DecodePathSwitchRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePathSwitchRequest(data []byte) (*PathSwitchRequest, error) {
	var m PathSwitchRequest
	err := ngapcodec.Decode("PathSwitchRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
