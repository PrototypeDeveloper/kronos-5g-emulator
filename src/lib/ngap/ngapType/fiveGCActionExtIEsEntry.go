// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGCActionExtIEsEntryPresent uint32

const (
	FiveGCActionExtIEsEntryPresentNothing FiveGCActionExtIEsEntryPresent = iota
	FiveGCActionExtIEsEntryPresentUnknown
)

type FiveGCActionExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      FiveGCActionExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGCActionExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGCAction-ExtIEsEntry", v)
}

// DecodeFiveGCActionExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGCActionExtIEsEntry(data []byte) (FiveGCActionExtIEsEntry, int, error) {
	var v FiveGCActionExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveGCAction-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGCActionExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
