// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CompletedCellsInTAIEUTRA represents ASN.1 CompletedCellsInTAI-EUTRA.
type CompletedCellsInTAIEUTRA struct {
	List []CompletedCellsInTAIEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CompletedCellsInTAIEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CompletedCellsInTAI-EUTRA", v)
}

// DecodeCompletedCellsInTAIEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCompletedCellsInTAIEUTRA(data []byte) (CompletedCellsInTAIEUTRA, int, error) {
	var v CompletedCellsInTAIEUTRA
	err := ngapcodec.DecodeParameter("CompletedCellsInTAI-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CompletedCellsInTAIEUTRA{}, 0, err
	}
	return v, len(data), err
}
