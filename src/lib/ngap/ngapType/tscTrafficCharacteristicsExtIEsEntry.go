// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TSCTrafficCharacteristicsExtIEsEntryPresent uint32

const (
	TSCTrafficCharacteristicsExtIEsEntryPresentNothing TSCTrafficCharacteristicsExtIEsEntryPresent = iota
	TSCTrafficCharacteristicsExtIEsEntryPresentUnknown
)

type TSCTrafficCharacteristicsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TSCTrafficCharacteristicsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCTrafficCharacteristicsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCTrafficCharacteristics-ExtIEsEntry", v)
}

// DecodeTSCTrafficCharacteristicsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTSCTrafficCharacteristicsExtIEsEntry(data []byte) (TSCTrafficCharacteristicsExtIEsEntry, int, error) {
	var v TSCTrafficCharacteristicsExtIEsEntry
	err := ngapcodec.DecodeParameter("TSCTrafficCharacteristics-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCTrafficCharacteristicsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
