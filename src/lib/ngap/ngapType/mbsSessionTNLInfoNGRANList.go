// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfoNGRANList represents ASN.1 MBS-SessionTNLInfoNGRANList.
type MBSSessionTNLInfoNGRANList struct {
	List []MBSSessionTNLInfoNGRANItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfoNGRANList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfoNGRANList", v)
}

// DecodeMBSSessionTNLInfoNGRANList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfoNGRANList(data []byte) (MBSSessionTNLInfoNGRANList, int, error) {
	var v MBSSessionTNLInfoNGRANList
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfoNGRANList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfoNGRANList{}, 0, err
	}
	return v, len(data), err
}
