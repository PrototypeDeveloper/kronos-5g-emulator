// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNTAIListforMDTItemExtIEsEntryPresent uint32

const (
	SNPNTAIListforMDTItemExtIEsEntryPresentNothing SNPNTAIListforMDTItemExtIEsEntryPresent = iota
	SNPNTAIListforMDTItemExtIEsEntryPresentUnknown
)

type SNPNTAIListforMDTItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNTAIListforMDTItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNTAIListforMDTItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNTAIListforMDTItem-ExtIEsEntry", v)
}

// DecodeSNPNTAIListforMDTItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNTAIListforMDTItemExtIEsEntry(data []byte) (SNPNTAIListforMDTItemExtIEsEntry, int, error) {
	var v SNPNTAIListforMDTItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPNTAIListforMDTItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNTAIListforMDTItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
