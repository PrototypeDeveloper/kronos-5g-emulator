// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNCellBasedMDTExtIEsEntryPresent uint32

const (
	SNPNCellBasedMDTExtIEsEntryPresentNothing SNPNCellBasedMDTExtIEsEntryPresent = iota
	SNPNCellBasedMDTExtIEsEntryPresentUnknown
)

type SNPNCellBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNCellBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNCellBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-CellBasedMDT-ExtIEsEntry", v)
}

// DecodeSNPNCellBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNCellBasedMDTExtIEsEntry(data []byte) (SNPNCellBasedMDTExtIEsEntry, int, error) {
	var v SNPNCellBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPN-CellBasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNCellBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
