// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceAreaInformationItem represents ASN.1 MBS-ServiceAreaInformationItem.
type MBSServiceAreaInformationItem struct {
	MBSAreaSessionID          MBSAreaSessionID
	MBSServiceAreaInformation MBSServiceAreaInformation
	HasIEExtensions           bool
	IEExtensions              []MBSServiceAreaInformationItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaInformationItem", v)
}

// DecodeMBSServiceAreaInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaInformationItem(data []byte) (MBSServiceAreaInformationItem, int, error) {
	var v MBSServiceAreaInformationItem
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaInformationItem{}, 0, err
	}
	return v, len(data), err
}
