// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInEAINR represents ASN.1 CompletedCellsInEAI-NR.
type CompletedCellsInEAINR struct {
	List []CompletedCellsInEAINRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-NR", v)
}

// DecodeCompletedCellsInEAINR decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAINR(data []byte) (CompletedCellsInEAINR, int, error) {
	var v CompletedCellsInEAINR
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAINR{}, 0, err
	}
	return v, len(data), err
}
