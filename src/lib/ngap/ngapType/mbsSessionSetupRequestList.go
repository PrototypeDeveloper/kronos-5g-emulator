// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupRequestList represents ASN.1 MBSSessionSetupRequestList.
type MBSSessionSetupRequestList struct {
	List []MBSSessionSetupRequestItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupRequestList", v)
}

// DecodeMBSSessionSetupRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupRequestList(data []byte) (MBSSessionSetupRequestList, int, error) {
	var v MBSSessionSetupRequestList
	err := ngapcodec.DecodeParameter("MBSSessionSetupRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupRequestList{}, 0, err
	}
	return v, len(data), err
}
