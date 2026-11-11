// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceNotifyItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceNotifyItemExtIEsEntryPresentNothing PDUSessionResourceNotifyItemExtIEsEntryPresent = iota
	PDUSessionResourceNotifyItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceNotifyItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceNotifyItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceNotifyItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyItemExtIEsEntry(data []byte) (PDUSessionResourceNotifyItemExtIEsEntry, int, error) {
	var v PDUSessionResourceNotifyItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
