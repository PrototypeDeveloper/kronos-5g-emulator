// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TNLAssociationItemExtIEsEntryPresent uint32

const (
	TNLAssociationItemExtIEsEntryPresentNothing TNLAssociationItemExtIEsEntryPresent = iota
	TNLAssociationItemExtIEsEntryPresentUnknown
)

type TNLAssociationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TNLAssociationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNLAssociationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNLAssociationItem-ExtIEsEntry", v)
}

// DecodeTNLAssociationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTNLAssociationItemExtIEsEntry(data []byte) (TNLAssociationItemExtIEsEntry, int, error) {
	var v TNLAssociationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TNLAssociationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNLAssociationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
