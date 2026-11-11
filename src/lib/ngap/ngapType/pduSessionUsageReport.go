// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionUsageReport represents ASN.1 PDUSessionUsageReport.
type PDUSessionUsageReport struct {
	RATType                   PDUSessionUsageReportRATType `aper:"valueExt,valueLB:0,valueUB:1"`
	PDUSessionTimedReportList VolumeTimedReportList
	HasIEExtensions           bool
	IEExtensions              []PDUSessionUsageReportExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionUsageReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionUsageReport", v)
}

// DecodePDUSessionUsageReport decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionUsageReport(data []byte) (PDUSessionUsageReport, int, error) {
	var v PDUSessionUsageReport
	err := ngapcodec.DecodeParameter("PDUSessionUsageReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionUsageReport{}, 0, err
	}
	return v, len(data), err
}
