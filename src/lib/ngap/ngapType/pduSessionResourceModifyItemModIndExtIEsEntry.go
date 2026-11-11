// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyItemModIndExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyItemModIndExtIEsEntryPresentNothing PDUSessionResourceModifyItemModIndExtIEsEntryPresent = iota
	PDUSessionResourceModifyItemModIndExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyItemModIndExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceModifyItemModIndExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModIndExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModInd-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyItemModIndExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModIndExtIEsEntry(data []byte) (PDUSessionResourceModifyItemModIndExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyItemModIndExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModInd-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModIndExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
