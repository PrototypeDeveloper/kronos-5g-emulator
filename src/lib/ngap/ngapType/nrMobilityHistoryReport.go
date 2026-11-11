// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRMobilityHistoryReport represents ASN.1 NRMobilityHistoryReport.
type NRMobilityHistoryReport struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRMobilityHistoryReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRMobilityHistoryReport", v)
}

// DecodeNRMobilityHistoryReport decodes a standalone APER value and reports the consumed octets.
func DecodeNRMobilityHistoryReport(data []byte) (NRMobilityHistoryReport, int, error) {
	var v NRMobilityHistoryReport
	err := ngapcodec.DecodeParameter("NRMobilityHistoryReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRMobilityHistoryReport{}, 0, err
	}
	return v, len(data), err
}
