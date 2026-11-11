// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionFailedtoSetupList represents ASN.1 MBSSessionFailedtoSetupList.
type MBSSessionFailedtoSetupList struct {
	List []MBSSessionFailedtoSetupItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionFailedtoSetupList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionFailedtoSetupList", v)
}

// DecodeMBSSessionFailedtoSetupList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionFailedtoSetupList(data []byte) (MBSSessionFailedtoSetupList, int, error) {
	var v MBSSessionFailedtoSetupList
	err := ngapcodec.DecodeParameter("MBSSessionFailedtoSetupList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionFailedtoSetupList{}, 0, err
	}
	return v, len(data), err
}
