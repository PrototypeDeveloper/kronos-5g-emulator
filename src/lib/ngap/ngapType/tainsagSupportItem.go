// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAINSAGSupportItem represents ASN.1 TAINSAGSupportItem.
type TAINSAGSupportItem struct {
	NSAGID               NSAGID
	NSAGSliceSupportList ExtendedSliceSupportList
	HasIEExtensions      bool
	IEExtensions         []TAINSAGSupportItemExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAINSAGSupportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAINSAGSupportItem", v)
}

// DecodeTAINSAGSupportItem decodes a standalone APER value and reports the consumed octets.
func DecodeTAINSAGSupportItem(data []byte) (TAINSAGSupportItem, int, error) {
	var v TAINSAGSupportItem
	err := ngapcodec.DecodeParameter("TAINSAGSupportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAINSAGSupportItem{}, 0, err
	}
	return v, len(data), err
}
