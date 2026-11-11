// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TimeBasedHandoverInformationExtIEsEntryPresent uint32

const (
	TimeBasedHandoverInformationExtIEsEntryPresentNothing TimeBasedHandoverInformationExtIEsEntryPresent = iota
	TimeBasedHandoverInformationExtIEsEntryPresentUnknown
)

type TimeBasedHandoverInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TimeBasedHandoverInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeBasedHandoverInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeBasedHandoverInformation-ExtIEsEntry", v)
}

// DecodeTimeBasedHandoverInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTimeBasedHandoverInformationExtIEsEntry(data []byte) (TimeBasedHandoverInformationExtIEsEntry, int, error) {
	var v TimeBasedHandoverInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("TimeBasedHandoverInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeBasedHandoverInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
