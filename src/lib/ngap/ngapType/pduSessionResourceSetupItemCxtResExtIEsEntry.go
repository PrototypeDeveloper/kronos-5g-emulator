// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupItemCxtResExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupItemCxtResExtIEsEntryPresentNothing PDUSessionResourceSetupItemCxtResExtIEsEntryPresent = iota
	PDUSessionResourceSetupItemCxtResExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupItemCxtResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceSetupItemCxtResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemCxtResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemCxtRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupItemCxtResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemCxtResExtIEsEntry(data []byte) (PDUSessionResourceSetupItemCxtResExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupItemCxtResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemCxtRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemCxtResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
