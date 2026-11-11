// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AMFTNLAssociationToAddItemExtIEsEntryPresent uint32

const (
	AMFTNLAssociationToAddItemExtIEsEntryPresentNothing AMFTNLAssociationToAddItemExtIEsEntryPresent = iota
	AMFTNLAssociationToAddItemExtIEsEntryPresentUnknown
)

type AMFTNLAssociationToAddItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AMFTNLAssociationToAddItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToAddItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToAddItem-ExtIEsEntry", v)
}

// DecodeAMFTNLAssociationToAddItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToAddItemExtIEsEntry(data []byte) (AMFTNLAssociationToAddItemExtIEsEntry, int, error) {
	var v AMFTNLAssociationToAddItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToAddItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToAddItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
