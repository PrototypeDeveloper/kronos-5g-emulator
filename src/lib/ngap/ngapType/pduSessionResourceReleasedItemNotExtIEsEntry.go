// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleasedItemNotExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleasedItemNotExtIEsEntryPresentNothing PDUSessionResourceReleasedItemNotExtIEsEntryPresent = iota
	PDUSessionResourceReleasedItemNotExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleasedItemNotExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceReleasedItemNotExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemNotExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemNot-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleasedItemNotExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemNotExtIEsEntry(data []byte) (PDUSessionResourceReleasedItemNotExtIEsEntry, int, error) {
	var v PDUSessionResourceReleasedItemNotExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemNot-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemNotExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
