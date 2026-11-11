// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSNGUFailureIndicationExtIEsEntryPresent uint32

const (
	MBSNGUFailureIndicationExtIEsEntryPresentNothing MBSNGUFailureIndicationExtIEsEntryPresent = iota
	MBSNGUFailureIndicationExtIEsEntryPresentUnknown
)

type MBSNGUFailureIndicationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MBSNGUFailureIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSNGUFailureIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-NGUFailureIndication-ExtIEsEntry", v)
}

// DecodeMBSNGUFailureIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSNGUFailureIndicationExtIEsEntry(data []byte) (MBSNGUFailureIndicationExtIEsEntry, int, error) {
	var v MBSNGUFailureIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-NGUFailureIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSNGUFailureIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
