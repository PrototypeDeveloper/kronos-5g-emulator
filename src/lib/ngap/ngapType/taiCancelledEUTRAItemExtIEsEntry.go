// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAICancelledEUTRAItemExtIEsEntryPresent uint32

const (
	TAICancelledEUTRAItemExtIEsEntryPresentNothing TAICancelledEUTRAItemExtIEsEntryPresent = iota
	TAICancelledEUTRAItemExtIEsEntryPresentUnknown
)

type TAICancelledEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAICancelledEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledEUTRA-Item-ExtIEsEntry", v)
}

// DecodeTAICancelledEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledEUTRAItemExtIEsEntry(data []byte) (TAICancelledEUTRAItemExtIEsEntry, int, error) {
	var v TAICancelledEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAICancelledEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
