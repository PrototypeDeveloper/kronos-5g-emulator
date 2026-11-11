// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToSetupItemSUResExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToSetupItemSUResExtIEsEntryPresentNothing PDUSessionResourceFailedToSetupItemSUResExtIEsEntryPresent = iota
	PDUSessionResourceFailedToSetupItemSUResExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToSetupItemSUResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToSetupItemSUResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemSUResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemSURes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToSetupItemSUResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemSUResExtIEsEntry(data []byte) (PDUSessionResourceFailedToSetupItemSUResExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToSetupItemSUResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemSURes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemSUResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
