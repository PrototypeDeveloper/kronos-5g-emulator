// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupItemSUResExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupItemSUResExtIEsEntryPresentNothing PDUSessionResourceSetupItemSUResExtIEsEntryPresent = iota
	PDUSessionResourceSetupItemSUResExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupItemSUResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceSetupItemSUResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemSUResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemSURes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupItemSUResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemSUResExtIEsEntry(data []byte) (PDUSessionResourceSetupItemSUResExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupItemSUResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemSURes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemSUResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
