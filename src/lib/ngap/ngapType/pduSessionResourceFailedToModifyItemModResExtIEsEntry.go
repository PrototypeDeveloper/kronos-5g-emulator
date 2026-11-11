// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToModifyItemModResExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToModifyItemModResExtIEsEntryPresentNothing PDUSessionResourceFailedToModifyItemModResExtIEsEntryPresent = iota
	PDUSessionResourceFailedToModifyItemModResExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToModifyItemModResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToModifyItemModResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyItemModResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyItemModRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToModifyItemModResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyItemModResExtIEsEntry(data []byte) (PDUSessionResourceFailedToModifyItemModResExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToModifyItemModResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyItemModRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyItemModResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
