// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRARFCN represents ASN.1 NRARFCN.
type NRARFCN struct {
	Value int64 `aper:"valueLB:0,valueUB:3279165"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRARFCN) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NRARFCN", v) }

// DecodeNRARFCN decodes a standalone APER value and reports the consumed octets.
func DecodeNRARFCN(data []byte) (NRARFCN, int, error) {
	var v NRARFCN
	err := ngapcodec.DecodeParameter("NRARFCN", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRARFCN{}, 0, err
	}
	return v, len(data), err
}
