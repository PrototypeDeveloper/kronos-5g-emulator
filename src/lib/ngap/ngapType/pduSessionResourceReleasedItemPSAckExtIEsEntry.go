// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleasedItemPSAckExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleasedItemPSAckExtIEsEntryPresentNothing PDUSessionResourceReleasedItemPSAckExtIEsEntryPresent = iota
	PDUSessionResourceReleasedItemPSAckExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleasedItemPSAckExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceReleasedItemPSAckExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemPSAckExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemPSAck-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleasedItemPSAckExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemPSAckExtIEsEntry(data []byte) (PDUSessionResourceReleasedItemPSAckExtIEsEntry, int, error) {
	var v PDUSessionResourceReleasedItemPSAckExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemPSAck-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemPSAckExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
