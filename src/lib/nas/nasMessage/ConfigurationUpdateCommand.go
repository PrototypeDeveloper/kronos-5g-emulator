package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// ConfigurationUpdateCommand updates the UE configuration.
type ConfigurationUpdateCommand struct {
	HasConfigurationUpdateIndication            bool
	ConfigurationUpdateIndication               nasType.ConfigurationUpdateIndication
	HasGUTI5GS                                  bool
	GUTI5GS                                     nasType.GUTI5GS
	HasTAIList                                  bool
	TAIList                                     nasType.TAIList
	HasAllowedNSSAI                             bool
	AllowedNSSAI                                nasType.AllowedNSSAI
	HasServiceAreaList                          bool
	ServiceAreaList                             nasType.ServiceAreaList
	HasFullNameForNetwork                       bool
	FullNameForNetwork                          nasType.NetworkName
	HasShortNameForNetwork                      bool
	ShortNameForNetwork                         nasType.NetworkName
	HasLocalTimeZone                            bool
	LocalTimeZone                               nasType.TimeZone
	HasUniversalTimeAndLocalTimeZone            bool
	UniversalTimeAndLocalTimeZone               nasType.TimeZoneAndTime
	HasNetworkDaylightSavingTime                bool
	NetworkDaylightSavingTime                   nasType.DaylightSavingTime
	HasLADNInformation                          bool
	LADNInformation                             nasType.LADNInformation
	HasMICOIndication                           bool
	MICOIndication                              nasType.MICOIndication
	HasNetworkSlicingIndication                 bool
	NetworkSlicingIndication                    nasType.NetworkSlicingIndication
	HasConfiguredNSSAI                          bool
	ConfiguredNSSAI                             nasType.ConfiguredNSSAI
	HasRejectedNSSAI                            bool
	RejectedNSSAI                               nasType.RejectedNSSAI
	HasOperatorDefinedAccessCategoryDefinitions bool
	OperatorDefinedAccessCategoryDefinitions    nasType.OperatorDefinedAccessCategoryDefinitions
	HasSMSIndication                            bool
	SMSIndication                               nasType.SMSIndication
	HasT3447Value                               bool
	T3447Value                                  nasType.T3447Value
	HasCAGInformationList                       bool
	CAGInformationList                          nasType.CAGInformationList
	HasUERadioCapabilityID                      bool
	UERadioCapabilityID                         nasType.UERadioCapabilityID
	HasUERadioCapabilityIDDeletionIndication    bool
	UERadioCapabilityIDDeletionIndication       nasType.UERadioCapabilityIDDeletionIndication
	HasRegistrationResult5GS                    bool
	RegistrationResult5GS                       nasType.RegistrationResult5GS
	HasTruncated5GSTMSIConfiguration            bool
	Truncated5GSTMSIConfiguration               nasType.Truncated5GSTMSIConfiguration
	HasAdditionalConfigurationIndication        bool
	AdditionalConfigurationIndication           nasType.AdditionalConfigurationIndication
	UnknownIEs                                  []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var configurationUpdateCommandOptions = nascodec.Schema{
	nasType.IEConfigurationUpdateIndication:            {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEGUTI5GS:                                  {Format: nascodec.TLVE, Min: 11, Max: 11},
	nasType.IETAIList:                                  {Format: nascodec.TLV, Min: 7, Max: 112},
	nasType.IEAllowedNSSAI:                             {Format: nascodec.TLV, Min: 2, Max: 72},
	nasType.IEServiceAreaList:                          {Format: nascodec.TLV, Min: 4, Max: 112},
	nasType.IEFullNameForNetwork:                       {Format: nascodec.TLV, Min: 1, Max: 255},
	nasType.IEShortNameForNetwork:                      {Format: nascodec.TLV, Min: 1, Max: 255},
	nasType.IELocalTimeZone:                            {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEUniversalTimeAndLocalTimeZone:            {Format: nascodec.TV, Min: 7, Max: 7},
	nasType.IENetworkDaylightSavingTime:                {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IELADNInformation:                          {Format: nascodec.TLVE, Min: 0, Max: 1712},
	nasType.IEMICOIndication:                           {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IENetworkSlicingIndication:                 {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IEConfiguredNSSAI:                          {Format: nascodec.TLV, Min: 2, Max: 144},
	nasType.IERejectedNSSAI:                            {Format: nascodec.TLV, Min: 2, Max: 40},
	nasType.IEOperatorDefinedAccessCategoryDefinitions: {Format: nascodec.TLVE, Min: 0, Max: 65535},
	nasType.IESMSIndication:                            {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IET3447Value:                               {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IECAGInformationList:                       {Format: nascodec.TLVE, Min: 0, Max: 65535},
	nasType.IEUERadioCapabilityID:                      {Format: nascodec.TLV, Min: 1, Max: 255},
	nasType.IEUERadioCapabilityIDDeletionIndication:    {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IERegistrationResult5GS:                    {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IETruncated5GSTMSIConfiguration:            {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEAdditionalConfigurationIndication:        {Format: nascodec.HalfTV, Min: 1, Max: 1},
}

func (m *ConfigurationUpdateCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil ConfigurationUpdateCommand")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgConfigurationUpdateCommand}
	var options []nasType.IE
	if m.HasConfigurationUpdateIndication {
		options = append(options, nasType.IE{ID: nasType.IEConfigurationUpdateIndication, Value: []byte{m.ConfigurationUpdateIndication.Value}})
	}
	if m.HasGUTI5GS {
		options = append(options, nasType.IE{ID: nasType.IEGUTI5GS, Value: m.GUTI5GS.Value})
	}
	if m.HasTAIList {
		options = append(options, nasType.IE{ID: nasType.IETAIList, Value: m.TAIList.Value})
	}
	if m.HasAllowedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IEAllowedNSSAI, Value: m.AllowedNSSAI.Value})
	}
	if m.HasServiceAreaList {
		options = append(options, nasType.IE{ID: nasType.IEServiceAreaList, Value: m.ServiceAreaList.Value})
	}
	if m.HasFullNameForNetwork {
		options = append(options, nasType.IE{ID: nasType.IEFullNameForNetwork, Value: m.FullNameForNetwork.Value})
	}
	if m.HasShortNameForNetwork {
		options = append(options, nasType.IE{ID: nasType.IEShortNameForNetwork, Value: m.ShortNameForNetwork.Value})
	}
	if m.HasLocalTimeZone {
		options = append(options, nasType.IE{ID: nasType.IELocalTimeZone, Value: []byte{m.LocalTimeZone.Value}})
	}
	if m.HasUniversalTimeAndLocalTimeZone {
		options = append(options, nasType.IE{ID: nasType.IEUniversalTimeAndLocalTimeZone, Value: m.UniversalTimeAndLocalTimeZone.Value})
	}
	if m.HasNetworkDaylightSavingTime {
		options = append(options, nasType.IE{ID: nasType.IENetworkDaylightSavingTime, Value: []byte{m.NetworkDaylightSavingTime.Value}})
	}
	if m.HasLADNInformation {
		options = append(options, nasType.IE{ID: nasType.IELADNInformation, Value: m.LADNInformation.Value})
	}
	if m.HasMICOIndication {
		options = append(options, nasType.IE{ID: nasType.IEMICOIndication, Value: []byte{m.MICOIndication.Value}})
	}
	if m.HasNetworkSlicingIndication {
		options = append(options, nasType.IE{ID: nasType.IENetworkSlicingIndication, Value: []byte{m.NetworkSlicingIndication.Value}})
	}
	if m.HasConfiguredNSSAI {
		options = append(options, nasType.IE{ID: nasType.IEConfiguredNSSAI, Value: m.ConfiguredNSSAI.Value})
	}
	if m.HasRejectedNSSAI {
		options = append(options, nasType.IE{ID: nasType.IERejectedNSSAI, Value: m.RejectedNSSAI.Value})
	}
	if m.HasOperatorDefinedAccessCategoryDefinitions {
		options = append(options, nasType.IE{ID: nasType.IEOperatorDefinedAccessCategoryDefinitions, Value: m.OperatorDefinedAccessCategoryDefinitions.Value})
	}
	if m.HasSMSIndication {
		options = append(options, nasType.IE{ID: nasType.IESMSIndication, Value: []byte{m.SMSIndication.Value}})
	}
	if m.HasT3447Value {
		options = append(options, nasType.IE{ID: nasType.IET3447Value, Value: []byte{m.T3447Value.Value}})
	}
	if m.HasCAGInformationList {
		options = append(options, nasType.IE{ID: nasType.IECAGInformationList, Value: m.CAGInformationList.Value})
	}
	if m.HasUERadioCapabilityID {
		options = append(options, nasType.IE{ID: nasType.IEUERadioCapabilityID, Value: m.UERadioCapabilityID.Value})
	}
	if m.HasUERadioCapabilityIDDeletionIndication {
		options = append(options, nasType.IE{ID: nasType.IEUERadioCapabilityIDDeletionIndication, Value: []byte{m.UERadioCapabilityIDDeletionIndication.Value}})
	}
	if m.HasRegistrationResult5GS {
		options = append(options, nasType.IE{ID: nasType.IERegistrationResult5GS, Value: []byte{m.RegistrationResult5GS.Value}})
	}
	if m.HasTruncated5GSTMSIConfiguration {
		options = append(options, nasType.IE{ID: nasType.IETruncated5GSTMSIConfiguration, Value: []byte{m.Truncated5GSTMSIConfiguration.Value}})
	}
	if m.HasAdditionalConfigurationIndication {
		options = append(options, nasType.IE{ID: nasType.IEAdditionalConfigurationIndication, Value: []byte{m.AdditionalConfigurationIndication.Value}})
	}
	w.Options(options, configurationUpdateCommandOptions)
	w.UnknownOptions(m.UnknownIEs, configurationUpdateCommandOptions)
	return w.Result()
}
func DecodeConfigurationUpdateCommand(data []byte) (*ConfigurationUpdateCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgConfigurationUpdateCommand {
		return nil, fmt.Errorf("unexpected header for ConfigurationUpdateCommand")
	}
	m := &ConfigurationUpdateCommand{}
	r := nascodec.Reader{Data: data[3:]}
	known, unknown := r.OptionsWithUnknown(configurationUpdateCommandOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode ConfigurationUpdateCommand: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEConfigurationUpdateIndication:
			m.HasConfigurationUpdateIndication = true
			m.ConfigurationUpdateIndication = nasType.ConfigurationUpdateIndication{Value: ie.Value[0]}
		case nasType.IEGUTI5GS:
			m.HasGUTI5GS = true
			m.GUTI5GS = nasType.GUTI5GS{Value: ie.Value}
		case nasType.IETAIList:
			m.HasTAIList = true
			m.TAIList = nasType.TAIList{Value: ie.Value}
		case nasType.IEAllowedNSSAI:
			m.HasAllowedNSSAI = true
			m.AllowedNSSAI = nasType.AllowedNSSAI{Value: ie.Value}
		case nasType.IEServiceAreaList:
			m.HasServiceAreaList = true
			m.ServiceAreaList = nasType.ServiceAreaList{Value: ie.Value}
		case nasType.IEFullNameForNetwork:
			m.HasFullNameForNetwork = true
			m.FullNameForNetwork = nasType.NetworkName{Value: ie.Value}
		case nasType.IEShortNameForNetwork:
			m.HasShortNameForNetwork = true
			m.ShortNameForNetwork = nasType.NetworkName{Value: ie.Value}
		case nasType.IELocalTimeZone:
			m.HasLocalTimeZone = true
			m.LocalTimeZone = nasType.TimeZone{Value: ie.Value[0]}
		case nasType.IEUniversalTimeAndLocalTimeZone:
			m.HasUniversalTimeAndLocalTimeZone = true
			m.UniversalTimeAndLocalTimeZone = nasType.TimeZoneAndTime{Value: ie.Value}
		case nasType.IENetworkDaylightSavingTime:
			m.HasNetworkDaylightSavingTime = true
			m.NetworkDaylightSavingTime = nasType.DaylightSavingTime{Value: ie.Value[0]}
		case nasType.IELADNInformation:
			m.HasLADNInformation = true
			m.LADNInformation = nasType.LADNInformation{Value: ie.Value}
		case nasType.IEMICOIndication:
			m.HasMICOIndication = true
			m.MICOIndication = nasType.MICOIndication{Value: ie.Value[0]}
		case nasType.IENetworkSlicingIndication:
			m.HasNetworkSlicingIndication = true
			m.NetworkSlicingIndication = nasType.NetworkSlicingIndication{Value: ie.Value[0]}
		case nasType.IEConfiguredNSSAI:
			m.HasConfiguredNSSAI = true
			m.ConfiguredNSSAI = nasType.ConfiguredNSSAI{Value: ie.Value}
		case nasType.IERejectedNSSAI:
			m.HasRejectedNSSAI = true
			m.RejectedNSSAI = nasType.RejectedNSSAI{Value: ie.Value}
		case nasType.IEOperatorDefinedAccessCategoryDefinitions:
			m.HasOperatorDefinedAccessCategoryDefinitions = true
			m.OperatorDefinedAccessCategoryDefinitions = nasType.OperatorDefinedAccessCategoryDefinitions{Value: ie.Value}
		case nasType.IESMSIndication:
			m.HasSMSIndication = true
			m.SMSIndication = nasType.SMSIndication{Value: ie.Value[0]}
		case nasType.IET3447Value:
			m.HasT3447Value = true
			m.T3447Value = nasType.T3447Value{Value: ie.Value[0]}
		case nasType.IECAGInformationList:
			m.HasCAGInformationList = true
			m.CAGInformationList = nasType.CAGInformationList{Value: ie.Value}
		case nasType.IEUERadioCapabilityID:
			m.HasUERadioCapabilityID = true
			m.UERadioCapabilityID = nasType.UERadioCapabilityID{Value: ie.Value}
		case nasType.IEUERadioCapabilityIDDeletionIndication:
			m.HasUERadioCapabilityIDDeletionIndication = true
			m.UERadioCapabilityIDDeletionIndication = nasType.UERadioCapabilityIDDeletionIndication{Value: ie.Value[0]}
		case nasType.IERegistrationResult5GS:
			m.HasRegistrationResult5GS = true
			m.RegistrationResult5GS = nasType.RegistrationResult5GS{Value: ie.Value[0]}
		case nasType.IETruncated5GSTMSIConfiguration:
			m.HasTruncated5GSTMSIConfiguration = true
			m.Truncated5GSTMSIConfiguration = nasType.Truncated5GSTMSIConfiguration{Value: ie.Value[0]}
		case nasType.IEAdditionalConfigurationIndication:
			m.HasAdditionalConfigurationIndication = true
			m.AdditionalConfigurationIndication = nasType.AdditionalConfigurationIndication{Value: ie.Value[0]}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
