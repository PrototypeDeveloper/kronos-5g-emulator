// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosCharacteristicsExtIEsEntryPresent uint32

const (
	QosCharacteristicsExtIEsEntryPresentNothing QosCharacteristicsExtIEsEntryPresent = iota
	QosCharacteristicsExtIEsEntryPresentUnknown
)

type QosCharacteristicsExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      QosCharacteristicsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosCharacteristicsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosCharacteristics-ExtIEsEntry", v)
}

// DecodeQosCharacteristicsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosCharacteristicsExtIEsEntry(data []byte) (QosCharacteristicsExtIEsEntry, int, error) {
	var v QosCharacteristicsExtIEsEntry
	err := ngapcodec.DecodeParameter("QosCharacteristics-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosCharacteristicsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
