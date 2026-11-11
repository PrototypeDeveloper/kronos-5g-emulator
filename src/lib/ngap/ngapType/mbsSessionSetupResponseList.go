// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupResponseList represents ASN.1 MBSSessionSetupResponseList.
type MBSSessionSetupResponseList struct {
	List []MBSSessionSetupResponseItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupResponseList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupResponseList", v)
}

// DecodeMBSSessionSetupResponseList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupResponseList(data []byte) (MBSSessionSetupResponseList, int, error) {
	var v MBSSessionSetupResponseList
	err := ngapcodec.DecodeParameter("MBSSessionSetupResponseList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupResponseList{}, 0, err
	}
	return v, len(data), err
}
