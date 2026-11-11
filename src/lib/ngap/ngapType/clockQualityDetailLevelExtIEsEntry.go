// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ClockQualityDetailLevelExtIEsEntryPresent uint32

const (
	ClockQualityDetailLevelExtIEsEntryPresentNothing ClockQualityDetailLevelExtIEsEntryPresent = iota
	ClockQualityDetailLevelExtIEsEntryPresentUnknown
)

type ClockQualityDetailLevelExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ClockQualityDetailLevelExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityDetailLevelExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityDetailLevel-ExtIEsEntry", v)
}

// DecodeClockQualityDetailLevelExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityDetailLevelExtIEsEntry(data []byte) (ClockQualityDetailLevelExtIEsEntry, int, error) {
	var v ClockQualityDetailLevelExtIEsEntry
	err := ngapcodec.DecodeParameter("ClockQualityDetailLevel-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityDetailLevelExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
