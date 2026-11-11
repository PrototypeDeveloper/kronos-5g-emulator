// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANTSSCellItemExtIEsEntryPresent uint32

const (
	RANTSSCellItemExtIEsEntryPresentNothing RANTSSCellItemExtIEsEntryPresent = iota
	RANTSSCellItemExtIEsEntryPresentUnknown
)

type RANTSSCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RANTSSCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANTSSCellItem-ExtIEsEntry", v)
}

// DecodeRANTSSCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSCellItemExtIEsEntry(data []byte) (RANTSSCellItemExtIEsEntry, int, error) {
	var v RANTSSCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RANTSSCellItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
