// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNListforMDTItemExtIEsEntryPresent uint32

const (
	SNPNListforMDTItemExtIEsEntryPresentNothing SNPNListforMDTItemExtIEsEntryPresent = iota
	SNPNListforMDTItemExtIEsEntryPresentUnknown
)

type SNPNListforMDTItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNListforMDTItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNListforMDTItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNListforMDTItem-ExtIEsEntry", v)
}

// DecodeSNPNListforMDTItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNListforMDTItemExtIEsEntry(data []byte) (SNPNListforMDTItemExtIEsEntry, int, error) {
	var v SNPNListforMDTItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPNListforMDTItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNListforMDTItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
