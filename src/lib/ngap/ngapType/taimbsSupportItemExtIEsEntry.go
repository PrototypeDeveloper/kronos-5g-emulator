// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIMBSSupportItemExtIEsEntryPresent uint32

const (
	TAIMBSSupportItemExtIEsEntryPresentNothing TAIMBSSupportItemExtIEsEntryPresent = iota
	TAIMBSSupportItemExtIEsEntryPresentUnknown
)

type TAIMBSSupportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIMBSSupportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIMBSSupportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIMBSSupportItem-ExtIEsEntry", v)
}

// DecodeTAIMBSSupportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIMBSSupportItemExtIEsEntry(data []byte) (TAIMBSSupportItemExtIEsEntry, int, error) {
	var v TAIMBSSupportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIMBSSupportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIMBSSupportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
