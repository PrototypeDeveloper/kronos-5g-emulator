// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EquivalentSNPNsItemExtIEsEntryPresent uint32

const (
	EquivalentSNPNsItemExtIEsEntryPresentNothing EquivalentSNPNsItemExtIEsEntryPresent = iota
	EquivalentSNPNsItemExtIEsEntryPresentUnknown
)

type EquivalentSNPNsItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EquivalentSNPNsItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EquivalentSNPNsItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EquivalentSNPNsItem-ExtIEsEntry", v)
}

// DecodeEquivalentSNPNsItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEquivalentSNPNsItemExtIEsEntry(data []byte) (EquivalentSNPNsItemExtIEsEntry, int, error) {
	var v EquivalentSNPNsItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EquivalentSNPNsItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EquivalentSNPNsItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
