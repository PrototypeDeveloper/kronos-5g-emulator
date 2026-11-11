// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleasedItemPSFailExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleasedItemPSFailExtIEsEntryPresentNothing PDUSessionResourceReleasedItemPSFailExtIEsEntryPresent = iota
	PDUSessionResourceReleasedItemPSFailExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleasedItemPSFailExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceReleasedItemPSFailExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemPSFailExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemPSFail-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleasedItemPSFailExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemPSFailExtIEsEntry(data []byte) (PDUSessionResourceReleasedItemPSFailExtIEsEntry, int, error) {
	var v PDUSessionResourceReleasedItemPSFailExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemPSFail-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemPSFailExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
