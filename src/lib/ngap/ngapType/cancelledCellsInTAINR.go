// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInTAINR represents ASN.1 CancelledCellsInTAI-NR.
type CancelledCellsInTAINR struct {
	List []CancelledCellsInTAINRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-NR", v)
}

// DecodeCancelledCellsInTAINR decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAINR(data []byte) (CancelledCellsInTAINR, int, error) {
	var v CancelledCellsInTAINR
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAINR{}, 0, err
	}
	return v, len(data), err
}
