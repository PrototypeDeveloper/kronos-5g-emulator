// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNTAIBasedMDTExtIEsEntryPresent uint32

const (
	SNPNTAIBasedMDTExtIEsEntryPresentNothing SNPNTAIBasedMDTExtIEsEntryPresent = iota
	SNPNTAIBasedMDTExtIEsEntryPresentUnknown
)

type SNPNTAIBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNTAIBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNTAIBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-TAIBasedMDT-ExtIEsEntry", v)
}

// DecodeSNPNTAIBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNTAIBasedMDTExtIEsEntry(data []byte) (SNPNTAIBasedMDTExtIEsEntry, int, error) {
	var v SNPNTAIBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPN-TAIBasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNTAIBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
