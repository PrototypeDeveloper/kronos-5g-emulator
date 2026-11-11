// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MonitoringRequestonAvailableBitrate represents ASN.1 MonitoringRequestonAvailableBitrate.
type MonitoringRequestonAvailableBitrate struct {
	MonitoringRequest                     MonitoringRequest
	HasDlAvailableBitrateReportThresholds bool
	DlAvailableBitrateReportThresholds    AvailableBitrateReportThresholdList
	HasUlAvailableBitrateReportThresholds bool
	UlAvailableBitrateReportThresholds    AvailableBitrateReportThresholdList
	HasIEExtensions                       bool
	IEExtensions                          []MonitoringRequestonAvailableBitrateExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MonitoringRequestonAvailableBitrate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MonitoringRequestonAvailableBitrate", v)
}

// DecodeMonitoringRequestonAvailableBitrate decodes a standalone APER value and reports the consumed octets.
func DecodeMonitoringRequestonAvailableBitrate(data []byte) (MonitoringRequestonAvailableBitrate, int, error) {
	var v MonitoringRequestonAvailableBitrate
	err := ngapcodec.DecodeParameter("MonitoringRequestonAvailableBitrate", data, &v)
	if ngapcodec.IsFatal(err) {
		return MonitoringRequestonAvailableBitrate{}, 0, err
	}
	return v, len(data), err
}
