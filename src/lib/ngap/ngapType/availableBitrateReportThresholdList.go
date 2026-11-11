// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AvailableBitrateReportThresholdList represents ASN.1 AvailableBitrateReportThresholdList.
type AvailableBitrateReportThresholdList struct {
	List []AvailableBitrateReportThresholdItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AvailableBitrateReportThresholdList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AvailableBitrateReportThresholdList", v)
}

// DecodeAvailableBitrateReportThresholdList decodes a standalone APER value and reports the consumed octets.
func DecodeAvailableBitrateReportThresholdList(data []byte) (AvailableBitrateReportThresholdList, int, error) {
	var v AvailableBitrateReportThresholdList
	err := ngapcodec.DecodeParameter("AvailableBitrateReportThresholdList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableBitrateReportThresholdList{}, 0, err
	}
	return v, len(data), err
}
