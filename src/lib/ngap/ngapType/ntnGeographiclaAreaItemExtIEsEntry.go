// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NTNGeographiclaAreaItemExtIEsEntryPresent uint32

const (
	NTNGeographiclaAreaItemExtIEsEntryPresentNothing NTNGeographiclaAreaItemExtIEsEntryPresent = iota
	NTNGeographiclaAreaItemExtIEsEntryPresentUnknown
)

type NTNGeographiclaAreaItemExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NTNGeographiclaAreaItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NTNGeographiclaAreaItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NTN-Geographicla-Area-Item-ExtIEsEntry", v)
}

// DecodeNTNGeographiclaAreaItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNTNGeographiclaAreaItemExtIEsEntry(data []byte) (NTNGeographiclaAreaItemExtIEsEntry, int, error) {
	var v NTNGeographiclaAreaItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NTN-Geographicla-Area-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NTNGeographiclaAreaItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
