// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTInventoryAssistanceInformationExtIEsEntryPresent uint32

const (
	AIoTInventoryAssistanceInformationExtIEsEntryPresentNothing AIoTInventoryAssistanceInformationExtIEsEntryPresent = iota
	AIoTInventoryAssistanceInformationExtIEsEntryPresentUnknown
)

type AIoTInventoryAssistanceInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTInventoryAssistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTInventoryAssistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-InventoryAssistanceInformation-ExtIEsEntry", v)
}

// DecodeAIoTInventoryAssistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTInventoryAssistanceInformationExtIEsEntry(data []byte) (AIoTInventoryAssistanceInformationExtIEsEntry, int, error) {
	var v AIoTInventoryAssistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-InventoryAssistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTInventoryAssistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
