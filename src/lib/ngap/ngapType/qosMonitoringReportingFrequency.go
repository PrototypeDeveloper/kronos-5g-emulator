// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosMonitoringReportingFrequency represents ASN.1 QosMonitoringReportingFrequency.
type QosMonitoringReportingFrequency struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:1800"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosMonitoringReportingFrequency) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosMonitoringReportingFrequency", v)
}

// DecodeQosMonitoringReportingFrequency decodes a standalone APER value and reports the consumed octets.
func DecodeQosMonitoringReportingFrequency(data []byte) (QosMonitoringReportingFrequency, int, error) {
	var v QosMonitoringReportingFrequency
	err := ngapcodec.DecodeParameter("QosMonitoringReportingFrequency", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosMonitoringReportingFrequency{}, 0, err
	}
	return v, len(data), err
}
