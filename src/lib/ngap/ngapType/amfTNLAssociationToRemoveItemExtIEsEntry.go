// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AMFTNLAssociationToRemoveItemExtIEsEntryPresent uint32

const (
	AMFTNLAssociationToRemoveItemExtIEsEntryPresentNothing AMFTNLAssociationToRemoveItemExtIEsEntryPresent = iota
	AMFTNLAssociationToRemoveItemExtIEsEntryPresentTNLAssociationTransportLayerAddressNGRAN
	AMFTNLAssociationToRemoveItemExtIEsEntryPresentUnknown
)

type AMFTNLAssociationToRemoveItemExtIEsEntry struct {
	ID                                       ProtocolExtensionID
	Criticality                              Criticality
	Present                                  AMFTNLAssociationToRemoveItemExtIEsEntryPresent
	TNLAssociationTransportLayerAddressNGRAN CPTransportLayerInformation `ngap:"id:168,criticality:reject,presence:optional"`
	UnknownValue                             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToRemoveItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToRemoveItem-ExtIEsEntry", v)
}

// DecodeAMFTNLAssociationToRemoveItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToRemoveItemExtIEsEntry(data []byte) (AMFTNLAssociationToRemoveItemExtIEsEntry, int, error) {
	var v AMFTNLAssociationToRemoveItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToRemoveItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToRemoveItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
