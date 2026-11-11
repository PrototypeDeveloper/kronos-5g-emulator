// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTCircleExtIEsEntryPresent uint32

const (
	MDTCircleExtIEsEntryPresentNothing MDTCircleExtIEsEntryPresent = iota
	MDTCircleExtIEsEntryPresentUnknown
)

type MDTCircleExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MDTCircleExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTCircleExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-circle-ExtIEsEntry", v)
}

// DecodeMDTCircleExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTCircleExtIEsEntry(data []byte) (MDTCircleExtIEsEntry, int, error) {
	var v MDTCircleExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-circle-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTCircleExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
