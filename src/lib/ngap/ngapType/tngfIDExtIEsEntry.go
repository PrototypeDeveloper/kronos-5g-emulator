// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TNGFIDExtIEsEntryPresent uint32

const (
	TNGFIDExtIEsEntryPresentNothing TNGFIDExtIEsEntryPresent = iota
	TNGFIDExtIEsEntryPresentUnknown
)

type TNGFIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      TNGFIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNGFIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNGF-ID-ExtIEsEntry", v)
}

// DecodeTNGFIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTNGFIDExtIEsEntry(data []byte) (TNGFIDExtIEsEntry, int, error) {
	var v TNGFIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TNGF-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNGFIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
