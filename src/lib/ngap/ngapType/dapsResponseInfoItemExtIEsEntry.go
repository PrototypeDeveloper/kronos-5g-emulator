// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DAPSResponseInfoItemExtIEsEntryPresent uint32

const (
	DAPSResponseInfoItemExtIEsEntryPresentNothing DAPSResponseInfoItemExtIEsEntryPresent = iota
	DAPSResponseInfoItemExtIEsEntryPresentUnknown
)

type DAPSResponseInfoItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DAPSResponseInfoItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSResponseInfoItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSResponseInfoItem-ExtIEsEntry", v)
}

// DecodeDAPSResponseInfoItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSResponseInfoItemExtIEsEntry(data []byte) (DAPSResponseInfoItemExtIEsEntry, int, error) {
	var v DAPSResponseInfoItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DAPSResponseInfoItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfoItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
