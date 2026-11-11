// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceAreaInformationList represents ASN.1 MBS-ServiceAreaInformationList.
type MBSServiceAreaInformationList struct {
	List []MBSServiceAreaInformationItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaInformationList", v)
}

// DecodeMBSServiceAreaInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaInformationList(data []byte) (MBSServiceAreaInformationList, int, error) {
	var v MBSServiceAreaInformationList
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaInformationList{}, 0, err
	}
	return v, len(data), err
}
