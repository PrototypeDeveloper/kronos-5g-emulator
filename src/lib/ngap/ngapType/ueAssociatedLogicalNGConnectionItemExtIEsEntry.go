// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEAssociatedLogicalNGConnectionItemExtIEsEntryPresent uint32

const (
	UEAssociatedLogicalNGConnectionItemExtIEsEntryPresentNothing UEAssociatedLogicalNGConnectionItemExtIEsEntryPresent = iota
	UEAssociatedLogicalNGConnectionItemExtIEsEntryPresentUnknown
)

type UEAssociatedLogicalNGConnectionItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEAssociatedLogicalNGConnectionItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAssociatedLogicalNGConnectionItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-associatedLogicalNG-connectionItem-ExtIEsEntry", v)
}

// DecodeUEAssociatedLogicalNGConnectionItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEAssociatedLogicalNGConnectionItemExtIEsEntry(data []byte) (UEAssociatedLogicalNGConnectionItemExtIEsEntry, int, error) {
	var v UEAssociatedLogicalNGConnectionItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UE-associatedLogicalNG-connectionItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAssociatedLogicalNGConnectionItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
