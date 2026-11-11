// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasConfigNameList represents ASN.1 WLANMeasConfigNameList.
type WLANMeasConfigNameList struct {
	List []WLANMeasConfigNameItem `aper:"sizeLB:1,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasConfigNameList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasConfigNameList", v)
}

// DecodeWLANMeasConfigNameList decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasConfigNameList(data []byte) (WLANMeasConfigNameList, int, error) {
	var v WLANMeasConfigNameList
	err := ngapcodec.DecodeParameter("WLANMeasConfigNameList", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasConfigNameList{}, 0, err
	}
	return v, len(data), err
}
