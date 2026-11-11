// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEPagingItemExtIEsEntryPresent uint32

const (
	UEPagingItemExtIEsEntryPresentNothing UEPagingItemExtIEsEntryPresent = iota
	UEPagingItemExtIEsEntryPresentUnknown
)

type UEPagingItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEPagingItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPagingItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-PagingItem-ExtIEsEntry", v)
}

// DecodeUEPagingItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEPagingItemExtIEsEntry(data []byte) (UEPagingItemExtIEsEntry, int, error) {
	var v UEPagingItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UE-PagingItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPagingItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
