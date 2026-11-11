// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetuporModifyRequestList represents ASN.1 MBSSessionSetuporModifyRequestList.
type MBSSessionSetuporModifyRequestList struct {
	List []MBSSessionSetuporModifyRequestItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetuporModifyRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetuporModifyRequestList", v)
}

// DecodeMBSSessionSetuporModifyRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetuporModifyRequestList(data []byte) (MBSSessionSetuporModifyRequestList, int, error) {
	var v MBSSessionSetuporModifyRequestList
	err := ngapcodec.DecodeParameter("MBSSessionSetuporModifyRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetuporModifyRequestList{}, 0, err
	}
	return v, len(data), err
}
