// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TooearlyIntersystemHOExtIEsEntryPresent uint32

const (
	TooearlyIntersystemHOExtIEsEntryPresentNothing TooearlyIntersystemHOExtIEsEntryPresent = iota
	TooearlyIntersystemHOExtIEsEntryPresentUnknown
)

type TooearlyIntersystemHOExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TooearlyIntersystemHOExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TooearlyIntersystemHOExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TooearlyIntersystemHO-ExtIEsEntry", v)
}

// DecodeTooearlyIntersystemHOExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTooearlyIntersystemHOExtIEsEntry(data []byte) (TooearlyIntersystemHOExtIEsEntry, int, error) {
	var v TooearlyIntersystemHOExtIEsEntry
	err := ngapcodec.DecodeParameter("TooearlyIntersystemHO-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TooearlyIntersystemHOExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
