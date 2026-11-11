// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSUPFailureIndicationItemExtIEsEntryPresent uint32

const (
	MBSUPFailureIndicationItemExtIEsEntryPresentNothing MBSUPFailureIndicationItemExtIEsEntryPresent = iota
	MBSUPFailureIndicationItemExtIEsEntryPresentUnknown
)

type MBSUPFailureIndicationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSUPFailureIndicationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSUPFailureIndicationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-UP-FailureIndicationItem-ExtIEsEntry", v)
}

// DecodeMBSUPFailureIndicationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSUPFailureIndicationItemExtIEsEntry(data []byte) (MBSUPFailureIndicationItemExtIEsEntry, int, error) {
	var v MBSUPFailureIndicationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-UP-FailureIndicationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSUPFailureIndicationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
