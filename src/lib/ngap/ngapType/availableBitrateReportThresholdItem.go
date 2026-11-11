// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AvailableBitrateReportThresholdItem represents ASN.1 AvailableBitrateReportThresholdItem.
type AvailableBitrateReportThresholdItem struct {
	ReportingThreshold ReportingThreshold
	HasIEExtensions    bool
	IEExtensions       []AvailableBitrateReportThresholdItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AvailableBitrateReportThresholdItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AvailableBitrateReportThresholdItem", v)
}

// DecodeAvailableBitrateReportThresholdItem decodes a standalone APER value and reports the consumed octets.
func DecodeAvailableBitrateReportThresholdItem(data []byte) (AvailableBitrateReportThresholdItem, int, error) {
	var v AvailableBitrateReportThresholdItem
	err := ngapcodec.DecodeParameter("AvailableBitrateReportThresholdItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableBitrateReportThresholdItem{}, 0, err
	}
	return v, len(data), err
}
