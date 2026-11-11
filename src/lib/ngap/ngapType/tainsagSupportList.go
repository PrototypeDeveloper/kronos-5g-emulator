// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAINSAGSupportList represents ASN.1 TAINSAGSupportList.
type TAINSAGSupportList struct {
	List []TAINSAGSupportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAINSAGSupportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAINSAGSupportList", v)
}

// DecodeTAINSAGSupportList decodes a standalone APER value and reports the consumed octets.
func DecodeTAINSAGSupportList(data []byte) (TAINSAGSupportList, int, error) {
	var v TAINSAGSupportList
	err := ngapcodec.DecodeParameter("TAINSAGSupportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAINSAGSupportList{}, 0, err
	}
	return v, len(data), err
}
