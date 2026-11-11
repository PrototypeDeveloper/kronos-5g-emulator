// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MulticastGroupPagingAreaExtIEsEntryPresent uint32

const (
	MulticastGroupPagingAreaExtIEsEntryPresentNothing MulticastGroupPagingAreaExtIEsEntryPresent = iota
	MulticastGroupPagingAreaExtIEsEntryPresentUnknown
)

type MulticastGroupPagingAreaExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MulticastGroupPagingAreaExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastGroupPagingAreaExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastGroupPagingArea-ExtIEsEntry", v)
}

// DecodeMulticastGroupPagingAreaExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastGroupPagingAreaExtIEsEntry(data []byte) (MulticastGroupPagingAreaExtIEsEntry, int, error) {
	var v MulticastGroupPagingAreaExtIEsEntry
	err := ngapcodec.DecodeParameter("MulticastGroupPagingArea-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastGroupPagingAreaExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
