// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInTAINR represents ASN.1 CompletedCellsInTAI-NR.
type CompletedCellsInTAINR struct {
	List []CompletedCellsInTAINRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAINR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-NR", v)
}

// DecodeCompletedCellsInTAINR decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAINR(data []byte) (CompletedCellsInTAINR, int, error) {
	var v CompletedCellsInTAINR
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAINR{}, 0, err
	}
	return v, len(data), err
}
