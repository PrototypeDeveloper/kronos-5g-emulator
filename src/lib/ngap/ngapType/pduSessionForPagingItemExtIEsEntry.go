// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionForPagingItemExtIEsEntryPresent uint32

const (
	PDUSessionForPagingItemExtIEsEntryPresentNothing PDUSessionForPagingItemExtIEsEntryPresent = iota
	PDUSessionForPagingItemExtIEsEntryPresentUnknown
)

type PDUSessionForPagingItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionForPagingItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionForPagingItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionForPagingItem-ExtIEsEntry", v)
}

// DecodePDUSessionForPagingItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionForPagingItemExtIEsEntry(data []byte) (PDUSessionForPagingItemExtIEsEntry, int, error) {
	var v PDUSessionForPagingItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionForPagingItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionForPagingItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
