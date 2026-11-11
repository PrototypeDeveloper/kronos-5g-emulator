// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M1ThresholdTypeExtIEsEntryPresent uint32

const (
	M1ThresholdTypeExtIEsEntryPresentNothing M1ThresholdTypeExtIEsEntryPresent = iota
	M1ThresholdTypeExtIEsEntryPresentUnknown
)

type M1ThresholdTypeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      M1ThresholdTypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ThresholdTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1ThresholdType-ExtIEsEntry", v)
}

// DecodeM1ThresholdTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM1ThresholdTypeExtIEsEntry(data []byte) (M1ThresholdTypeExtIEsEntry, int, error) {
	var v M1ThresholdTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("M1ThresholdType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ThresholdTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
