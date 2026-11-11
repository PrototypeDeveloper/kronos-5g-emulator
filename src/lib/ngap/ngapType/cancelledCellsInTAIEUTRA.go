// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInTAIEUTRA represents ASN.1 CancelledCellsInTAI-EUTRA.
type CancelledCellsInTAIEUTRA struct {
	List []CancelledCellsInTAIEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInTAIEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInTAI-EUTRA", v)
}

// DecodeCancelledCellsInTAIEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInTAIEUTRA(data []byte) (CancelledCellsInTAIEUTRA, int, error) {
	var v CancelledCellsInTAIEUTRA
	err := ngapcodec.DecodeParameter("CancelledCellsInTAI-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInTAIEUTRA{}, 0, err
	}
	return v, len(data), err
}
