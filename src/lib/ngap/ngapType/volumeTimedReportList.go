// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// VolumeTimedReportList represents ASN.1 VolumeTimedReportList.
type VolumeTimedReportList struct {
	List []VolumeTimedReportItem `aper:"sizeLB:1,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v VolumeTimedReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("VolumeTimedReportList", v)
}

// DecodeVolumeTimedReportList decodes a standalone APER value and reports the consumed octets.
func DecodeVolumeTimedReportList(data []byte) (VolumeTimedReportList, int, error) {
	var v VolumeTimedReportList
	err := ngapcodec.DecodeParameter("VolumeTimedReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return VolumeTimedReportList{}, 0, err
	}
	return v, len(data), err
}
