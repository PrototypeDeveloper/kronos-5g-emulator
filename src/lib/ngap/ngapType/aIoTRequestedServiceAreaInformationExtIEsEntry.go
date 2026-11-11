// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTRequestedServiceAreaInformationExtIEsEntryPresent uint32

const (
	AIoTRequestedServiceAreaInformationExtIEsEntryPresentNothing AIoTRequestedServiceAreaInformationExtIEsEntryPresent = iota
	AIoTRequestedServiceAreaInformationExtIEsEntryPresentUnknown
)

type AIoTRequestedServiceAreaInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTRequestedServiceAreaInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedServiceAreaInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedServiceAreaInformation-ExtIEsEntry", v)
}

// DecodeAIoTRequestedServiceAreaInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedServiceAreaInformationExtIEsEntry(data []byte) (AIoTRequestedServiceAreaInformationExtIEsEntry, int, error) {
	var v AIoTRequestedServiceAreaInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-RequestedServiceAreaInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedServiceAreaInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
