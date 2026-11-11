// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AMFTNLAssociationSetupItemExtIEsEntryPresent uint32

const (
	AMFTNLAssociationSetupItemExtIEsEntryPresentNothing AMFTNLAssociationSetupItemExtIEsEntryPresent = iota
	AMFTNLAssociationSetupItemExtIEsEntryPresentUnknown
)

type AMFTNLAssociationSetupItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AMFTNLAssociationSetupItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationSetupItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationSetupItem-ExtIEsEntry", v)
}

// DecodeAMFTNLAssociationSetupItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationSetupItemExtIEsEntry(data []byte) (AMFTNLAssociationSetupItemExtIEsEntry, int, error) {
	var v AMFTNLAssociationSetupItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationSetupItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationSetupItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
