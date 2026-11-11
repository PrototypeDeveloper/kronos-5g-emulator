// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ERABInformationList represents ASN.1 E-RABInformationList.
type ERABInformationList struct {
	List []ERABInformationItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ERABInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("E-RABInformationList", v)
}

// DecodeERABInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeERABInformationList(data []byte) (ERABInformationList, int, error) {
	var v ERABInformationList
	err := ngapcodec.DecodeParameter("E-RABInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return ERABInformationList{}, 0, err
	}
	return v, len(data), err
}
