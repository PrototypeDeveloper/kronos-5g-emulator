// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestTAIList represents ASN.1 AreaOfInterestTAIList.
type AreaOfInterestTAIList struct {
	List []AreaOfInterestTAIItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestTAIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestTAIList", v)
}

// DecodeAreaOfInterestTAIList decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestTAIList(data []byte) (AreaOfInterestTAIList, int, error) {
	var v AreaOfInterestTAIList
	err := ngapcodec.DecodeParameter("AreaOfInterestTAIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestTAIList{}, 0, err
	}
	return v, len(data), err
}
