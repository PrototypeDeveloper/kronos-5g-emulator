// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSActiveSessionInformationTargettoSourceList represents ASN.1 MBS-ActiveSessionInformation-TargettoSourceList.
type MBSActiveSessionInformationTargettoSourceList struct {
	List []MBSActiveSessionInformationTargettoSourceItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationTargettoSourceList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-TargettoSourceList", v)
}

// DecodeMBSActiveSessionInformationTargettoSourceList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationTargettoSourceList(data []byte) (MBSActiveSessionInformationTargettoSourceList, int, error) {
	var v MBSActiveSessionInformationTargettoSourceList
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-TargettoSourceList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationTargettoSourceList{}, 0, err
	}
	return v, len(data), err
}
