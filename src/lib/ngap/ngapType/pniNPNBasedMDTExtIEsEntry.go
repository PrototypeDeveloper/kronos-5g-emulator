// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PNINPNBasedMDTExtIEsEntryPresent uint32

const (
	PNINPNBasedMDTExtIEsEntryPresentNothing PNINPNBasedMDTExtIEsEntryPresent = iota
	PNINPNBasedMDTExtIEsEntryPresentUnknown
)

type PNINPNBasedMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PNINPNBasedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNBasedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPNBasedMDT-ExtIEsEntry", v)
}

// DecodePNINPNBasedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNBasedMDTExtIEsEntry(data []byte) (PNINPNBasedMDTExtIEsEntry, int, error) {
	var v PNINPNBasedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("PNI-NPNBasedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNBasedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
