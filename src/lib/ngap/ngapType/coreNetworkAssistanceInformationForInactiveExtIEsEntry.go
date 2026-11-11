// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresent uint32

const (
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentNothing CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresent = iota
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentEUTRAPagingeDRXInformation
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentExtendedUEIdentityIndexValue
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentUERadioCapabilityForPaging
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentMicoAllPLMN
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentNRPagingeDRXInformation
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentPagingCauseIndicationForVoiceService
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentPEIPSassistanceInformation
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentHashedUEIdentityIndexValue
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentCNMTCommunicationHandling
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentLPWUSPSAssistanceInformation
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentFurtherExtendedUEIdentityIndexValue
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentLPWUSDisableIndication
	CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresentUnknown
)

type CoreNetworkAssistanceInformationForInactiveExtIEsEntry struct {
	ID                                   ProtocolExtensionID
	Criticality                          Criticality
	Present                              CoreNetworkAssistanceInformationForInactiveExtIEsEntryPresent
	EUTRAPagingeDRXInformation           EUTRAPagingeDRXInformation           `ngap:"id:223,criticality:ignore,presence:optional"`
	ExtendedUEIdentityIndexValue         ExtendedUEIdentityIndexValue         `ngap:"id:280,criticality:ignore,presence:optional"`
	UERadioCapabilityForPaging           UERadioCapabilityForPaging           `ngap:"id:118,criticality:ignore,presence:optional"`
	MicoAllPLMN                          MicoAllPLMN                          `ngap:"id:282,criticality:ignore,presence:optional"`
	NRPagingeDRXInformation              NRPagingeDRXInformation              `ngap:"id:332,criticality:ignore,presence:optional"`
	PagingCauseIndicationForVoiceService PagingCauseIndicationForVoiceService `ngap:"id:343,criticality:ignore,presence:optional"`
	PEIPSassistanceInformation           PEIPSassistanceInformation           `ngap:"id:344,criticality:ignore,presence:optional"`
	HashedUEIdentityIndexValue           HashedUEIdentityIndexValue           `ngap:"id:365,criticality:ignore,presence:optional"`
	CNMTCommunicationHandling            CNMTCommunicationHandling            `ngap:"id:405,criticality:ignore,presence:optional"`
	LPWUSPSAssistanceInformation         LPWUSPSAssistanceInformation         `ngap:"id:477,criticality:ignore,presence:optional"`
	FurtherExtendedUEIdentityIndexValue  FurtherExtendedUEIdentityIndexValue  `ngap:"id:478,criticality:ignore,presence:optional"`
	LPWUSDisableIndication               LPWUSDisableIndication               `ngap:"id:495,criticality:ignore,presence:optional"`
	UnknownValue                         []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CoreNetworkAssistanceInformationForInactiveExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CoreNetworkAssistanceInformationForInactive-ExtIEsEntry", v)
}

// DecodeCoreNetworkAssistanceInformationForInactiveExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCoreNetworkAssistanceInformationForInactiveExtIEsEntry(data []byte) (CoreNetworkAssistanceInformationForInactiveExtIEsEntry, int, error) {
	var v CoreNetworkAssistanceInformationForInactiveExtIEsEntry
	err := ngapcodec.DecodeParameter("CoreNetworkAssistanceInformationForInactive-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CoreNetworkAssistanceInformationForInactiveExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
