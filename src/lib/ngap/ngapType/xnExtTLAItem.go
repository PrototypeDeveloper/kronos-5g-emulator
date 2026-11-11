// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XnExtTLAItem represents ASN.1 XnExtTLA-Item.
type XnExtTLAItem struct {
	HasIPsecTLA       bool
	IPsecTLA          TransportLayerAddress
	HasGTPTLAs        bool
	GTPTLAs           XnGTPTLAs
	HasIEExtensions   bool
	IEExtensions      []XnExtTLAItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnExtTLAItem) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("XnExtTLA-Item", v) }

// DecodeXnExtTLAItem decodes a standalone APER value and reports the consumed octets.
func DecodeXnExtTLAItem(data []byte) (XnExtTLAItem, int, error) {
	var v XnExtTLAItem
	err := ngapcodec.DecodeParameter("XnExtTLA-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnExtTLAItem{}, 0, err
	}
	return v, len(data), err
}
