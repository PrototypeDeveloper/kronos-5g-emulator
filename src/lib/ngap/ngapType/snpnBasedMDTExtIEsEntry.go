// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNBasedMDTExtIEsEntryPresent uint32

const (
	SNPNBasedMDTExtIEsEntryPresentNothing SNPNBasedMDTExtIEsEntryPresent = iota
	SNPNBasedMDTExtIEsEntryPresentUnknown
)

type SNPNBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNPNBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-BasedMDT-ExtIEsEntry", v)
}

// DecodeSNPNBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNBasedMDTExtIEsEntry(data []byte) (SNPNBasedMDTExtIEsEntry, int, error) {
	var v SNPNBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPN-BasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
