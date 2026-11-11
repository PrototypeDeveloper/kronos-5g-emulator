// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSUPFailureIndicationList represents ASN.1 MBS-UP-FailureIndicationList.
type MBSUPFailureIndicationList struct {
	List []MBSUPFailureIndicationItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSUPFailureIndicationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-UP-FailureIndicationList", v)
}

// DecodeMBSUPFailureIndicationList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSUPFailureIndicationList(data []byte) (MBSUPFailureIndicationList, int, error) {
	var v MBSUPFailureIndicationList
	err := ngapcodec.DecodeParameter("MBS-UP-FailureIndicationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSUPFailureIndicationList{}, 0, err
	}
	return v, len(data), err
}
