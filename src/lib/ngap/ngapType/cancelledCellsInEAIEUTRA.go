// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelledCellsInEAIEUTRA represents ASN.1 CancelledCellsInEAI-EUTRA.
type CancelledCellsInEAIEUTRA struct {
	List []CancelledCellsInEAIEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelledCellsInEAIEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelledCellsInEAI-EUTRA", v)
}

// DecodeCancelledCellsInEAIEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeCancelledCellsInEAIEUTRA(data []byte) (CancelledCellsInEAIEUTRA, int, error) {
	var v CancelledCellsInEAIEUTRA
	err := ngapcodec.DecodeParameter("CancelledCellsInEAI-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelledCellsInEAIEUTRA{}, 0, err
	}
	return v, len(data), err
}
