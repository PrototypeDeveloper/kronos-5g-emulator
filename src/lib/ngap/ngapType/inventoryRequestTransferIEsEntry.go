// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InventoryRequestTransferIEsEntryPresent uint32

const (
	InventoryRequestTransferIEsEntryPresentNothing InventoryRequestTransferIEsEntryPresent = iota
	InventoryRequestTransferIEsEntryPresentAIoTCorrelationIdentifier
	InventoryRequestTransferIEsEntryPresentAIoTDeviceIdentificationRequested
	InventoryRequestTransferIEsEntryPresentAIoTRequestedServiceAreaInformation
	InventoryRequestTransferIEsEntryPresentAIoTInventoryAssistanceInformation
	InventoryRequestTransferIEsEntryPresentAIoTFollowonCommandIndication
	InventoryRequestTransferIEsEntryPresentAIoTSecurityInfo
	InventoryRequestTransferIEsEntryPresentUnknown
)

type InventoryRequestTransferIEsEntry struct {
	ID                                  ProtocolIEID
	Criticality                         Criticality
	Present                             InventoryRequestTransferIEsEntryPresent
	AIoTCorrelationIdentifier           AIoTCorrelationIdentifier           `ngap:"id:466,criticality:reject,presence:mandatory"`
	AIoTDeviceIdentificationRequested   AIoTDeviceIdentificationRequested   `ngap:"id:468,criticality:reject,presence:mandatory"`
	AIoTRequestedServiceAreaInformation AIoTRequestedServiceAreaInformation `ngap:"id:469,criticality:reject,presence:mandatory"`
	AIoTInventoryAssistanceInformation  AIoTInventoryAssistanceInformation  `ngap:"id:470,criticality:reject,presence:mandatory"`
	AIoTFollowonCommandIndication       AIoTFollowonCommandIndication       `ngap:"id:471,criticality:reject,presence:optional"`
	AIoTSecurityInfo                    AIoTSecurityInfo                    `ngap:"id:496,criticality:reject,presence:mandatory"`
	UnknownValue                        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryRequestTransferIEsEntry", v)
}

// DecodeInventoryRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryRequestTransferIEsEntry(data []byte) (InventoryRequestTransferIEsEntry, int, error) {
	var v InventoryRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("InventoryRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
