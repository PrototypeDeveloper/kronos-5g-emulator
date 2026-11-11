// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAICancelledNRItemExtIEsEntryPresent uint32

const (
	TAICancelledNRItemExtIEsEntryPresentNothing TAICancelledNRItemExtIEsEntryPresent = iota
	TAICancelledNRItemExtIEsEntryPresentUnknown
)

type TAICancelledNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAICancelledNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAICancelledNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAICancelledNR-Item-ExtIEsEntry", v)
}

// DecodeTAICancelledNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAICancelledNRItemExtIEsEntry(data []byte) (TAICancelledNRItemExtIEsEntry, int, error) {
	var v TAICancelledNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAICancelledNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAICancelledNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
