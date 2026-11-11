// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSServiceAreaExtIEsEntryPresent uint32

const (
	MBSServiceAreaExtIEsEntryPresentNothing MBSServiceAreaExtIEsEntryPresent = iota
	MBSServiceAreaExtIEsEntryPresentUnknown
)

type MBSServiceAreaExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MBSServiceAreaExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceArea-ExtIEsEntry", v)
}

// DecodeMBSServiceAreaExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaExtIEsEntry(data []byte) (MBSServiceAreaExtIEsEntry, int, error) {
	var v MBSServiceAreaExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-ServiceArea-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
