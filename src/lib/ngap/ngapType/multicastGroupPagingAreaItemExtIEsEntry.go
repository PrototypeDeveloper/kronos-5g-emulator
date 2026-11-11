// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MulticastGroupPagingAreaItemExtIEsEntryPresent uint32

const (
	MulticastGroupPagingAreaItemExtIEsEntryPresentNothing MulticastGroupPagingAreaItemExtIEsEntryPresent = iota
	MulticastGroupPagingAreaItemExtIEsEntryPresentUnknown
)

type MulticastGroupPagingAreaItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MulticastGroupPagingAreaItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastGroupPagingAreaItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastGroupPagingAreaItem-ExtIEsEntry", v)
}

// DecodeMulticastGroupPagingAreaItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastGroupPagingAreaItemExtIEsEntry(data []byte) (MulticastGroupPagingAreaItemExtIEsEntry, int, error) {
	var v MulticastGroupPagingAreaItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MulticastGroupPagingAreaItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastGroupPagingAreaItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
