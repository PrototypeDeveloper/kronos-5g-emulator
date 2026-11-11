// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIListForPaging represents ASN.1 TAIListForPaging.
type TAIListForPaging struct {
	List []TAIListForPagingItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForPaging", v)
}

// DecodeTAIListForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForPaging(data []byte) (TAIListForPaging, int, error) {
	var v TAIListForPaging
	err := ngapcodec.DecodeParameter("TAIListForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForPaging{}, 0, err
	}
	return v, len(data), err
}
