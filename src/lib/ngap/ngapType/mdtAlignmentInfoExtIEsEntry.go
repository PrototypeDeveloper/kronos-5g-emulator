// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTAlignmentInfoExtIEsEntryPresent uint32

const (
	MDTAlignmentInfoExtIEsEntryPresentNothing MDTAlignmentInfoExtIEsEntryPresent = iota
	MDTAlignmentInfoExtIEsEntryPresentUnknown
)

type MDTAlignmentInfoExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MDTAlignmentInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTAlignmentInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-AlignmentInfo-ExtIEsEntry", v)
}

// DecodeMDTAlignmentInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTAlignmentInfoExtIEsEntry(data []byte) (MDTAlignmentInfoExtIEsEntry, int, error) {
	var v MDTAlignmentInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-AlignmentInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTAlignmentInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
