// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDCancelledNR represents ASN.1 CellIDCancelledNR.
type CellIDCancelledNR struct {
	List []CellIDCancelledNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledNR", v)
}

// DecodeCellIDCancelledNR decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledNR(data []byte) (CellIDCancelledNR, int, error) {
	var v CellIDCancelledNR
	err := ngapcodec.DecodeParameter("CellIDCancelledNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledNR{}, 0, err
	}
	return v, len(data), err
}
