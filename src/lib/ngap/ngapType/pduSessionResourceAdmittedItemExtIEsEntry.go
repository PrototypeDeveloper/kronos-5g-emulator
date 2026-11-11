// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceAdmittedItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceAdmittedItemExtIEsEntryPresentNothing PDUSessionResourceAdmittedItemExtIEsEntryPresent = iota
	PDUSessionResourceAdmittedItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceAdmittedItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceAdmittedItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceAdmittedItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceAdmittedItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceAdmittedItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceAdmittedItemExtIEsEntry(data []byte) (PDUSessionResourceAdmittedItemExtIEsEntry, int, error) {
	var v PDUSessionResourceAdmittedItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceAdmittedItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceAdmittedItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
