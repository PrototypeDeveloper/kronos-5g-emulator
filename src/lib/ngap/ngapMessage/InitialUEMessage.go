// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InitialUEMessage represents ASN.1 InitialUEMessage.
type InitialUEMessage struct {
	RANUENGAPID                            ngapType.RANUENGAPID             `ngap:"id:85,criticality:reject,presence:mandatory"`
	NASPDU                                 ngapType.NASPDU                  `ngap:"id:38,criticality:reject,presence:mandatory"`
	UserLocationInformation                ngapType.UserLocationInformation `ngap:"id:121,criticality:reject,presence:mandatory"`
	RRCEstablishmentCause                  ngapType.RRCEstablishmentCause   `ngap:"id:90,criticality:ignore,presence:mandatory"`
	HasFiveGSTMSI                          bool
	FiveGSTMSI                             ngapType.FiveGSTMSI `ngap:"id:26,criticality:reject,presence:optional"`
	HasAMFSetID                            bool
	AMFSetID                               ngapType.AMFSetID `ngap:"id:3,criticality:ignore,presence:optional"`
	HasUEContextRequest                    bool
	UEContextRequest                       ngapType.UEContextRequest `ngap:"id:112,criticality:ignore,presence:optional"`
	HasAllowedNSSAI                        bool
	AllowedNSSAI                           ngapType.AllowedNSSAI `ngap:"id:0,criticality:reject,presence:optional"`
	HasSourceToTargetAMFInformationReroute bool
	SourceToTargetAMFInformationReroute    ngapType.SourceToTargetAMFInformationReroute `ngap:"id:171,criticality:ignore,presence:optional"`
	HasSelectedPLMNIdentity                bool
	SelectedPLMNIdentity                   ngapType.PLMNIdentity `ngap:"id:174,criticality:ignore,presence:optional"`
	HasIABNodeIndication                   bool
	IABNodeIndication                      ngapType.IABNodeIndication `ngap:"id:201,criticality:reject,presence:optional"`
	HasCEmodeBSupportIndicator             bool
	CEmodeBSupportIndicator                ngapType.CEmodeBSupportIndicator `ngap:"id:224,criticality:reject,presence:optional"`
	HasLTEMIndication                      bool
	LTEMIndication                         ngapType.LTEMIndication `ngap:"id:225,criticality:ignore,presence:optional"`
	HasEDTSession                          bool
	EDTSession                             ngapType.EDTSession `ngap:"id:227,criticality:ignore,presence:optional"`
	HasAuthenticatedIndication             bool
	AuthenticatedIndication                ngapType.AuthenticatedIndication `ngap:"id:245,criticality:ignore,presence:optional"`
	HasNPNAccessInformation                bool
	NPNAccessInformation                   ngapType.NPNAccessInformation `ngap:"id:259,criticality:reject,presence:optional"`
	HasRedCapIndication                    bool
	RedCapIndication                       ngapType.RedCapIndication `ngap:"id:333,criticality:ignore,presence:optional"`
	HasSelectedNID                         bool
	SelectedNID                            ngapType.NID `ngap:"id:371,criticality:ignore,presence:optional"`
	HasMobileIABNodeIndication             bool
	MobileIABNodeIndication                ngapType.MobileIABNodeIndication `ngap:"id:402,criticality:reject,presence:optional"`
	HasPartiallyAllowedNSSAI               bool
	PartiallyAllowedNSSAI                  ngapType.PartiallyAllowedNSSAI `ngap:"id:414,criticality:ignore,presence:optional"`
	HasERedCapIndication                   bool
	ERedCapIndication                      ngapType.ERedCapIndication `ngap:"id:427,criticality:ignore,presence:optional"`
	HasAUN3DeviceAccessInfo                bool
	AUN3DeviceAccessInfo                   ngapType.AUN3DeviceAccessInfo `ngap:"id:440,criticality:ignore,presence:optional"`
	HasGUAMI                               bool
	GUAMI                                  ngapType.GUAMI `ngap:"id:28,criticality:ignore,presence:optional"`
	HasGUAMIType                           bool
	GUAMIType                              ngapType.GUAMIType `ngap:"id:176,criticality:ignore,presence:optional"`
	HasRequestedNSSAI                      bool
	RequestedNSSAI                         ngapType.RequestedNSSAI `ngap:"id:454,criticality:ignore,presence:optional"`
	AdditionalIEs                          []ngapType.UnknownIE
	UnknownExtensions                      []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InitialUEMessage) Encode() ([]byte, error) { return ngapcodec.Encode("InitialUEMessage", m) }

// DecodeInitialUEMessage decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInitialUEMessage(data []byte) (*InitialUEMessage, error) {
	var m InitialUEMessage
	err := ngapcodec.Decode("InitialUEMessage", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
