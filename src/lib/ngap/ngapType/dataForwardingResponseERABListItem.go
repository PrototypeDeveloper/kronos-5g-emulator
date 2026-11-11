// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingResponseERABListItem represents ASN.1 DataForwardingResponseERABListItem.
type DataForwardingResponseERABListItem struct {
	ERABID                       ERABID
	DLForwardingUPTNLInformation UPTransportLayerInformation
	HasIEExtensions              bool
	IEExtensions                 []DataForwardingResponseERABListItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseERABListItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseERABListItem", v)
}

// DecodeDataForwardingResponseERABListItem decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseERABListItem(data []byte) (DataForwardingResponseERABListItem, int, error) {
	var v DataForwardingResponseERABListItem
	err := ngapcodec.DecodeParameter("DataForwardingResponseERABListItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseERABListItem{}, 0, err
	}
	return v, len(data), err
}
