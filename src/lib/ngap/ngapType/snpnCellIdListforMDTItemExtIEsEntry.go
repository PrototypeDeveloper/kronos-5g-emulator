// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNCellIdListforMDTItemExtIEsEntryPresent uint32

const (
	SNPNCellIdListforMDTItemExtIEsEntryPresentNothing SNPNCellIdListforMDTItemExtIEsEntryPresent = iota
	SNPNCellIdListforMDTItemExtIEsEntryPresentUnknown
)

type SNPNCellIdListforMDTItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNCellIdListforMDTItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNCellIdListforMDTItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNCellIdListforMDTItem-ExtIEsEntry", v)
}

// DecodeSNPNCellIdListforMDTItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNCellIdListforMDTItemExtIEsEntry(data []byte) (SNPNCellIdListforMDTItemExtIEsEntry, int, error) {
	var v SNPNCellIdListforMDTItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPNCellIdListforMDTItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNCellIdListforMDTItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
