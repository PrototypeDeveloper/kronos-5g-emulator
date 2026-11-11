// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAINSAGSupportItemExtIEsEntryPresent uint32

const (
	TAINSAGSupportItemExtIEsEntryPresentNothing TAINSAGSupportItemExtIEsEntryPresent = iota
	TAINSAGSupportItemExtIEsEntryPresentUnknown
)

type TAINSAGSupportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAINSAGSupportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAINSAGSupportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAINSAGSupportItem-ExtIEsEntry", v)
}

// DecodeTAINSAGSupportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAINSAGSupportItemExtIEsEntry(data []byte) (TAINSAGSupportItemExtIEsEntry, int, error) {
	var v TAINSAGSupportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAINSAGSupportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAINSAGSupportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
