// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSResponseInfoItem represents ASN.1 DAPSResponseInfoItem.
type DAPSResponseInfoItem struct {
	DRBID             DRBID
	DAPSResponseInfo  DAPSResponseInfo
	HasIEExtension    bool
	IEExtension       []DAPSResponseInfoItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSResponseInfoItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSResponseInfoItem", v)
}

// DecodeDAPSResponseInfoItem decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSResponseInfoItem(data []byte) (DAPSResponseInfoItem, int, error) {
	var v DAPSResponseInfoItem
	err := ngapcodec.DecodeParameter("DAPSResponseInfoItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfoItem{}, 0, err
	}
	return v, len(data), err
}
