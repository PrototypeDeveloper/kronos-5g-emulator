// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TSCTrafficCharacteristicsFeedbackExtIEsEntryPresent uint32

const (
	TSCTrafficCharacteristicsFeedbackExtIEsEntryPresentNothing TSCTrafficCharacteristicsFeedbackExtIEsEntryPresent = iota
	TSCTrafficCharacteristicsFeedbackExtIEsEntryPresentUnknown
)

type TSCTrafficCharacteristicsFeedbackExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TSCTrafficCharacteristicsFeedbackExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCTrafficCharacteristicsFeedbackExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCTrafficCharacteristicsFeedback-ExtIEsEntry", v)
}

// DecodeTSCTrafficCharacteristicsFeedbackExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTSCTrafficCharacteristicsFeedbackExtIEsEntry(data []byte) (TSCTrafficCharacteristicsFeedbackExtIEsEntry, int, error) {
	var v TSCTrafficCharacteristicsFeedbackExtIEsEntry
	err := ngapcodec.DecodeParameter("TSCTrafficCharacteristicsFeedback-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCTrafficCharacteristicsFeedbackExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
