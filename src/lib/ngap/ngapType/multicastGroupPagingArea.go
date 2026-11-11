// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastGroupPagingArea represents ASN.1 MulticastGroupPagingArea.
type MulticastGroupPagingArea struct {
	MBSAreaTAIList    MBSAreaTAIList
	HasIEExtensions   bool
	IEExtensions      []MulticastGroupPagingAreaExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastGroupPagingArea) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastGroupPagingArea", v)
}

// DecodeMulticastGroupPagingArea decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastGroupPagingArea(data []byte) (MulticastGroupPagingArea, int, error) {
	var v MulticastGroupPagingArea
	err := ngapcodec.DecodeParameter("MulticastGroupPagingArea", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastGroupPagingArea{}, 0, err
	}
	return v, len(data), err
}
