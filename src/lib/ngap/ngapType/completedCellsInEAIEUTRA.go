// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInEAIEUTRA represents ASN.1 CompletedCellsInEAI-EUTRA.
type CompletedCellsInEAIEUTRA struct {
	List []CompletedCellsInEAIEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInEAIEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInEAI-EUTRA", v)
}

// DecodeCompletedCellsInEAIEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInEAIEUTRA(data []byte) (CompletedCellsInEAIEUTRA, int, error) {
	var v CompletedCellsInEAIEUTRA
	err := ngapcodec.DecodeParameter("CompletedCellsInEAI-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInEAIEUTRA{}, 0, err
	}
	return v, len(data), err
}
