// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyItemModResExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyItemModResExtIEsEntryPresentNothing PDUSessionResourceModifyItemModResExtIEsEntryPresent = iota
	PDUSessionResourceModifyItemModResExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyItemModResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceModifyItemModResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyItemModResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModResExtIEsEntry(data []byte) (PDUSessionResourceModifyItemModResExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyItemModResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
