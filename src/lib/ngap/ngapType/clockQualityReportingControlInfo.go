// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityReportingControlInfo represents ASN.1 ClockQualityReportingControlInfo.
type ClockQualityReportingControlInfo struct {
	ClockQualityDetailLevel ClockQualityDetailLevel
	HasIEExtensions         bool
	IEExtensions            []ClockQualityReportingControlInfoExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityReportingControlInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityReportingControlInfo", v)
}

// DecodeClockQualityReportingControlInfo decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityReportingControlInfo(data []byte) (ClockQualityReportingControlInfo, int, error) {
	var v ClockQualityReportingControlInfo
	err := ngapcodec.DecodeParameter("ClockQualityReportingControlInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityReportingControlInfo{}, 0, err
	}
	return v, len(data), err
}
