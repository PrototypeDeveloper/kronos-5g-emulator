// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToSetupItemCxtResExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToSetupItemCxtResExtIEsEntryPresentNothing PDUSessionResourceFailedToSetupItemCxtResExtIEsEntryPresent = iota
	PDUSessionResourceFailedToSetupItemCxtResExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToSetupItemCxtResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemCxtRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToSetupItemCxtResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemCxtResExtIEsEntry(data []byte) (PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemCxtRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemCxtResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
