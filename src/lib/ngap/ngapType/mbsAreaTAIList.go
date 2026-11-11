// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSAreaTAIList represents ASN.1 MBS-AreaTAIList.
type MBSAreaTAIList struct {
	List []TAI `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSAreaTAIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-AreaTAIList", v)
}

// DecodeMBSAreaTAIList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSAreaTAIList(data []byte) (MBSAreaTAIList, int, error) {
	var v MBSAreaTAIList
	err := ngapcodec.DecodeParameter("MBS-AreaTAIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSAreaTAIList{}, 0, err
	}
	return v, len(data), err
}
