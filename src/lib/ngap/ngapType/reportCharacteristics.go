// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportCharacteristics represents ASN.1 ReportCharacteristics.
type ReportCharacteristics struct {
	Value BitString `aper:"sizeLB:32,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportCharacteristics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReportCharacteristics", v)
}

// DecodeReportCharacteristics decodes a standalone APER value and reports the consumed octets.
func DecodeReportCharacteristics(data []byte) (ReportCharacteristics, int, error) {
	var v ReportCharacteristics
	err := ngapcodec.DecodeParameter("ReportCharacteristics", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportCharacteristics{}, 0, err
	}
	return v, len(data), err
}
