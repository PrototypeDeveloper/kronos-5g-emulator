// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSIntendedServiceAreaList represents ASN.1 MBS-IntendedServiceAreaList.
type MBSIntendedServiceAreaList struct {
	List []IntendedServiceAreaCoordinates `aper:"sizeLB:1,sizeUB:65536"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSIntendedServiceAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-IntendedServiceAreaList", v)
}

// DecodeMBSIntendedServiceAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSIntendedServiceAreaList(data []byte) (MBSIntendedServiceAreaList, int, error) {
	var v MBSIntendedServiceAreaList
	err := ngapcodec.DecodeParameter("MBS-IntendedServiceAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSIntendedServiceAreaList{}, 0, err
	}
	return v, len(data), err
}
