// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionFSAIDList represents ASN.1 MBS-SessionFSAIDList.
type MBSSessionFSAIDList struct {
	List []MBSSessionFSAID `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionFSAIDList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionFSAIDList", v)
}

// DecodeMBSSessionFSAIDList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionFSAIDList(data []byte) (MBSSessionFSAIDList, int, error) {
	var v MBSSessionFSAIDList
	err := ngapcodec.DecodeParameter("MBS-SessionFSAIDList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionFSAIDList{}, 0, err
	}
	return v, len(data), err
}
