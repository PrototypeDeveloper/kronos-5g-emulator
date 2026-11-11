// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SensorMeasConfigNameList represents ASN.1 SensorMeasConfigNameList.
type SensorMeasConfigNameList struct {
	List []SensorMeasConfigNameItem `aper:"sizeLB:1,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SensorMeasConfigNameList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SensorMeasConfigNameList", v)
}

// DecodeSensorMeasConfigNameList decodes a standalone APER value and reports the consumed octets.
func DecodeSensorMeasConfigNameList(data []byte) (SensorMeasConfigNameList, int, error) {
	var v SensorMeasConfigNameList
	err := ngapcodec.DecodeParameter("SensorMeasConfigNameList", data, &v)
	if ngapcodec.IsFatal(err) {
		return SensorMeasConfigNameList{}, 0, err
	}
	return v, len(data), err
}
