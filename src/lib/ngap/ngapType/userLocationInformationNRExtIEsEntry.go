// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationNRExtIEsEntryPresent uint32

const (
	UserLocationInformationNRExtIEsEntryPresentNothing UserLocationInformationNRExtIEsEntryPresent = iota
	UserLocationInformationNRExtIEsEntryPresentPSCellInformation
	UserLocationInformationNRExtIEsEntryPresentNID
	UserLocationInformationNRExtIEsEntryPresentNRNTNTAIInformation
	UserLocationInformationNRExtIEsEntryPresentMobileIABMTUserLocationInformation
	UserLocationInformationNRExtIEsEntryPresentAdditionalULI
	UserLocationInformationNRExtIEsEntryPresentAerialUEFlightInformationReporting
	UserLocationInformationNRExtIEsEntryPresentUnknown
)

type UserLocationInformationNRExtIEsEntry struct {
	ID                                 ProtocolExtensionID
	Criticality                        Criticality
	Present                            UserLocationInformationNRExtIEsEntryPresent
	PSCellInformation                  NGRANCGI                           `ngap:"id:149,criticality:ignore,presence:optional"`
	NID                                NID                                `ngap:"id:263,criticality:reject,presence:optional"`
	NRNTNTAIInformation                NRNTNTAIInformation                `ngap:"id:287,criticality:ignore,presence:optional"`
	MobileIABMTUserLocationInformation MobileIABMTUserLocationInformation `ngap:"id:401,criticality:ignore,presence:optional"`
	AdditionalULI                      AdditionalULI                      `ngap:"id:483,criticality:ignore,presence:optional"`
	AerialUEFlightInformationReporting AerialUEFlightInformationReporting `ngap:"id:493,criticality:ignore,presence:optional"`
	UnknownValue                       []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationNRExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationNR-ExtIEsEntry", v)
}

// DecodeUserLocationInformationNRExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationNRExtIEsEntry(data []byte) (UserLocationInformationNRExtIEsEntry, int, error) {
	var v UserLocationInformationNRExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationNR-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationNRExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
