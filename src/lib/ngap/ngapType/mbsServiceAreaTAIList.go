// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceAreaTAIList represents ASN.1 MBS-ServiceAreaTAIList.
type MBSServiceAreaTAIList struct {
	List []TAI `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaTAIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaTAIList", v)
}

// DecodeMBSServiceAreaTAIList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaTAIList(data []byte) (MBSServiceAreaTAIList, int, error) {
	var v MBSServiceAreaTAIList
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaTAIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaTAIList{}, 0, err
	}
	return v, len(data), err
}
