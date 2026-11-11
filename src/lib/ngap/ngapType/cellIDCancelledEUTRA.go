// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDCancelledEUTRA represents ASN.1 CellIDCancelledEUTRA.
type CellIDCancelledEUTRA struct {
	List []CellIDCancelledEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledEUTRA", v)
}

// DecodeCellIDCancelledEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledEUTRA(data []byte) (CellIDCancelledEUTRA, int, error) {
	var v CellIDCancelledEUTRA
	err := ngapcodec.DecodeParameter("CellIDCancelledEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledEUTRA{}, 0, err
	}
	return v, len(data), err
}
