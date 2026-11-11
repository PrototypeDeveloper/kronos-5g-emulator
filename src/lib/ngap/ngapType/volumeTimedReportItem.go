// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// VolumeTimedReportItem represents ASN.1 VolumeTimedReport-Item.
type VolumeTimedReportItem struct {
	StartTimeStamp    []byte `aper:"sizeLB:4,sizeUB:4"`
	EndTimeStamp      []byte `aper:"sizeLB:4,sizeUB:4"`
	UsageCountUL      uint64 `aper:"valueLB:0,valueUB:18446744073709551615"`
	UsageCountDL      uint64 `aper:"valueLB:0,valueUB:18446744073709551615"`
	HasIEExtensions   bool
	IEExtensions      []VolumeTimedReportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v VolumeTimedReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("VolumeTimedReport-Item", v)
}

// DecodeVolumeTimedReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeVolumeTimedReportItem(data []byte) (VolumeTimedReportItem, int, error) {
	var v VolumeTimedReportItem
	err := ngapcodec.DecodeParameter("VolumeTimedReport-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return VolumeTimedReportItem{}, 0, err
	}
	return v, len(data), err
}
