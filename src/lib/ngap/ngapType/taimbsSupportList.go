// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIMBSSupportList represents ASN.1 TAIMBSSupportList.
type TAIMBSSupportList struct {
	List []TAIMBSSupportItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIMBSSupportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIMBSSupportList", v)
}

// DecodeTAIMBSSupportList decodes a standalone APER value and reports the consumed octets.
func DecodeTAIMBSSupportList(data []byte) (TAIMBSSupportList, int, error) {
	var v TAIMBSSupportList
	err := ngapcodec.DecodeParameter("TAIMBSSupportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIMBSSupportList{}, 0, err
	}
	return v, len(data), err
}
