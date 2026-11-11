// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationRequestItem represents ASN.1 NGAPIESupportInformationRequestItem.
type NGAPIESupportInformationRequestItem struct {
	NgapProtocolIEId  ProtocolIEID
	HasIEExtensions   bool
	IEExtensions      []NGAPIESupportInformationRequestItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationRequestItem", v)
}

// DecodeNGAPIESupportInformationRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationRequestItem(data []byte) (NGAPIESupportInformationRequestItem, int, error) {
	var v NGAPIESupportInformationRequestItem
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationRequestItem{}, 0, err
	}
	return v, len(data), err
}
