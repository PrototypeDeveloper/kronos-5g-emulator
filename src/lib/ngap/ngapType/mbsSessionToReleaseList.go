// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionToReleaseList represents ASN.1 MBSSessionToReleaseList.
type MBSSessionToReleaseList struct {
	List []MBSSessionToReleaseItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionToReleaseList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionToReleaseList", v)
}

// DecodeMBSSessionToReleaseList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionToReleaseList(data []byte) (MBSSessionToReleaseList, int, error) {
	var v MBSSessionToReleaseList
	err := ngapcodec.DecodeParameter("MBSSessionToReleaseList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionToReleaseList{}, 0, err
	}
	return v, len(data), err
}
