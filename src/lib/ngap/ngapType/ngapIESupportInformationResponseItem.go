// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationResponseItem represents ASN.1 NGAPIESupportInformationResponseItem.
type NGAPIESupportInformationResponseItem struct {
	NgapProtocolIEId           ProtocolIEID
	NgapProtocolIESupportInfo  NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo  `aper:"valueExt,valueLB:0,valueUB:1"`
	NgapProtocolIEPresenceInfo NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions            bool
	IEExtensions               []NGAPIESupportInformationResponseItemExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationResponseItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationResponseItem", v)
}

// DecodeNGAPIESupportInformationResponseItem decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationResponseItem(data []byte) (NGAPIESupportInformationResponseItem, int, error) {
	var v NGAPIESupportInformationResponseItem
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationResponseItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationResponseItem{}, 0, err
	}
	return v, len(data), err
}
