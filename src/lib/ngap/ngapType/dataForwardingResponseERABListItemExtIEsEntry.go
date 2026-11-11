// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DataForwardingResponseERABListItemExtIEsEntryPresent uint32

const (
	DataForwardingResponseERABListItemExtIEsEntryPresentNothing DataForwardingResponseERABListItemExtIEsEntryPresent = iota
	DataForwardingResponseERABListItemExtIEsEntryPresentUnknown
)

type DataForwardingResponseERABListItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DataForwardingResponseERABListItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseERABListItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseERABListItem-ExtIEsEntry", v)
}

// DecodeDataForwardingResponseERABListItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseERABListItemExtIEsEntry(data []byte) (DataForwardingResponseERABListItemExtIEsEntry, int, error) {
	var v DataForwardingResponseERABListItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DataForwardingResponseERABListItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseERABListItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
