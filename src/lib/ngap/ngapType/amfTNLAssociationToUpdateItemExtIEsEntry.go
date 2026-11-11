// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AMFTNLAssociationToUpdateItemExtIEsEntryPresent uint32

const (
	AMFTNLAssociationToUpdateItemExtIEsEntryPresentNothing AMFTNLAssociationToUpdateItemExtIEsEntryPresent = iota
	AMFTNLAssociationToUpdateItemExtIEsEntryPresentUnknown
)

type AMFTNLAssociationToUpdateItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AMFTNLAssociationToUpdateItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToUpdateItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToUpdateItem-ExtIEsEntry", v)
}

// DecodeAMFTNLAssociationToUpdateItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToUpdateItemExtIEsEntry(data []byte) (AMFTNLAssociationToUpdateItemExtIEsEntry, int, error) {
	var v AMFTNLAssociationToUpdateItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToUpdateItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToUpdateItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
