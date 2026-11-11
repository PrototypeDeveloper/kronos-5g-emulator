// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInEAINR represents ASN.1 CancelledCellsInEAI-NR.
type CancelledCellsInEAINR struct {
	List []CancelledCellsInEAINRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-NR", v)
}

// DecodeCancelledCellsInEAINR decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAINR(data []byte) (CancelledCellsInEAINR, int, error) {
	var v CancelledCellsInEAINR
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAINR{}, 0, err
	}
	return v, len(data), err
}
