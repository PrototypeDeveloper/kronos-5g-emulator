// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReportArea represents ASN.1 ReportArea.
const (
	ReportAreaPresentCell Enumerated = 0
)

type ReportArea struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReportArea) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ReportArea", v) }

// DecodeReportArea decodes a standalone APER value and reports the consumed octets.
func DecodeReportArea(data []byte) (ReportArea, int, error) {
	var v ReportArea
	err := ngapcodec.DecodeParameter("ReportArea", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReportArea{}, 0, err
	}
	return v, len(data), err
}
