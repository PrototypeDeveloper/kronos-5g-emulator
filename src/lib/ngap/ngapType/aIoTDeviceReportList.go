// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTDeviceReportList represents ASN.1 AIoT-DeviceReportList.
type AIoTDeviceReportList struct {
	List []AIoTDeviceReportItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTDeviceReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-DeviceReportList", v)
}

// DecodeAIoTDeviceReportList decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTDeviceReportList(data []byte) (AIoTDeviceReportList, int, error) {
	var v AIoTDeviceReportList
	err := ngapcodec.DecodeParameter("AIoT-DeviceReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTDeviceReportList{}, 0, err
	}
	return v, len(data), err
}
