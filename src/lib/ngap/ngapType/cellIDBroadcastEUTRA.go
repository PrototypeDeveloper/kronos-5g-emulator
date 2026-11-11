// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDBroadcastEUTRA represents ASN.1 CellIDBroadcastEUTRA.
type CellIDBroadcastEUTRA struct {
	List []CellIDBroadcastEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastEUTRA", v)
}

// DecodeCellIDBroadcastEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastEUTRA(data []byte) (CellIDBroadcastEUTRA, int, error) {
	var v CellIDBroadcastEUTRA
	err := ngapcodec.DecodeParameter("CellIDBroadcastEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastEUTRA{}, 0, err
	}
	return v, len(data), err
}
