// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// Paging represents ASN.1 Paging.
type Paging struct {
	UEPagingIdentity                ngapType.UEPagingIdentity `ngap:"id:115,criticality:ignore,presence:mandatory"`
	HasPagingDRX                    bool
	PagingDRX                       ngapType.PagingDRX        `ngap:"id:50,criticality:ignore,presence:optional"`
	TAIListForPaging                ngapType.TAIListForPaging `ngap:"id:103,criticality:ignore,presence:mandatory"`
	HasPagingPriority               bool
	PagingPriority                  ngapType.PagingPriority `ngap:"id:52,criticality:ignore,presence:optional"`
	HasUERadioCapabilityForPaging   bool
	UERadioCapabilityForPaging      ngapType.UERadioCapabilityForPaging `ngap:"id:118,criticality:ignore,presence:optional"`
	HasPagingOrigin                 bool
	PagingOrigin                    ngapType.PagingOrigin `ngap:"id:51,criticality:ignore,presence:optional"`
	HasAssistanceDataForPaging      bool
	AssistanceDataForPaging         ngapType.AssistanceDataForPaging `ngap:"id:11,criticality:ignore,presence:optional"`
	HasNBIoTPagingEDRXInfo          bool
	NBIoTPagingEDRXInfo             ngapType.NBIoTPagingEDRXInfo `ngap:"id:203,criticality:ignore,presence:optional"`
	HasNBIoTPagingDRX               bool
	NBIoTPagingDRX                  ngapType.NBIoTPagingDRX `ngap:"id:202,criticality:ignore,presence:optional"`
	HasEnhancedCoverageRestriction  bool
	EnhancedCoverageRestriction     ngapType.EnhancedCoverageRestriction `ngap:"id:205,criticality:ignore,presence:optional"`
	HasWUSAssistanceInformation     bool
	WUSAssistanceInformation        ngapType.WUSAssistanceInformation `ngap:"id:208,criticality:ignore,presence:optional"`
	HasEUTRAPagingeDRXInformation   bool
	EUTRAPagingeDRXInformation      ngapType.EUTRAPagingeDRXInformation `ngap:"id:223,criticality:ignore,presence:optional"`
	HasCEmodeBrestricted            bool
	CEmodeBrestricted               ngapType.CEmodeBrestricted `ngap:"id:222,criticality:ignore,presence:optional"`
	HasNRPagingeDRXInformation      bool
	NRPagingeDRXInformation         ngapType.NRPagingeDRXInformation `ngap:"id:332,criticality:ignore,presence:optional"`
	HasPagingCause                  bool
	PagingCause                     ngapType.PagingCause `ngap:"id:342,criticality:ignore,presence:optional"`
	HasPEIPSassistanceInformation   bool
	PEIPSassistanceInformation      ngapType.PEIPSassistanceInformation `ngap:"id:344,criticality:ignore,presence:optional"`
	HasLPWUSPSAssistanceInformation bool
	LPWUSPSAssistanceInformation    ngapType.LPWUSPSAssistanceInformation `ngap:"id:477,criticality:ignore,presence:optional"`
	HasLPWUSDisableIndication       bool
	LPWUSDisableIndication          ngapType.LPWUSDisableIndication `ngap:"id:495,criticality:ignore,presence:optional"`
	AdditionalIEs                   []ngapType.UnknownIE
	UnknownExtensions               []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *Paging) Encode() ([]byte, error) { return ngapcodec.Encode("Paging", m) }

// DecodePaging decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePaging(data []byte) (*Paging, error) {
	var m Paging
	err := ngapcodec.Decode("Paging", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
