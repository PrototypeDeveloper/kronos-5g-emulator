// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleasedItemRelResExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleasedItemRelResExtIEsEntryPresentNothing PDUSessionResourceReleasedItemRelResExtIEsEntryPresent = iota
	PDUSessionResourceReleasedItemRelResExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleasedItemRelResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceReleasedItemRelResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemRelResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemRelRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleasedItemRelResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemRelResExtIEsEntry(data []byte) (PDUSessionResourceReleasedItemRelResExtIEsEntry, int, error) {
	var v PDUSessionResourceReleasedItemRelResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemRelRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemRelResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
