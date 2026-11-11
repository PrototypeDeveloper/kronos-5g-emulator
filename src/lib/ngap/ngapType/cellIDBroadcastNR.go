// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CellIDBroadcastNR represents ASN.1 CellIDBroadcastNR.
type CellIDBroadcastNR struct {
	List []CellIDBroadcastNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDBroadcastNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDBroadcastNR", v)
}

// DecodeCellIDBroadcastNR decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDBroadcastNR(data []byte) (CellIDBroadcastNR, int, error) {
	var v CellIDBroadcastNR
	err := ngapcodec.DecodeParameter("CellIDBroadcastNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDBroadcastNR{}, 0, err
	}
	return v, len(data), err
}
