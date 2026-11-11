// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfo5GCList represents ASN.1 MBS-SessionTNLInfo5GCList.
type MBSSessionTNLInfo5GCList struct {
	List []MBSSessionTNLInfo5GCItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfo5GCList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfo5GCList", v)
}

// DecodeMBSSessionTNLInfo5GCList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfo5GCList(data []byte) (MBSSessionTNLInfo5GCList, int, error) {
	var v MBSSessionTNLInfo5GCList
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfo5GCList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfo5GCList{}, 0, err
	}
	return v, len(data), err
}
