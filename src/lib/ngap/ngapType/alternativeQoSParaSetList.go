// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AlternativeQoSParaSetList represents ASN.1 AlternativeQoSParaSetList.
type AlternativeQoSParaSetList struct {
	List []AlternativeQoSParaSetItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AlternativeQoSParaSetList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AlternativeQoSParaSetList", v)
}

// DecodeAlternativeQoSParaSetList decodes a standalone APER value and reports the consumed octets.
func DecodeAlternativeQoSParaSetList(data []byte) (AlternativeQoSParaSetList, int, error) {
	var v AlternativeQoSParaSetList
	err := ngapcodec.DecodeParameter("AlternativeQoSParaSetList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AlternativeQoSParaSetList{}, 0, err
	}
	return v, len(data), err
}
