// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSActiveSessionInformationSourcetoTargetList represents ASN.1 MBS-ActiveSessionInformation-SourcetoTargetList.
type MBSActiveSessionInformationSourcetoTargetList struct {
	List []MBSActiveSessionInformationSourcetoTargetItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationSourcetoTargetList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-SourcetoTargetList", v)
}

// DecodeMBSActiveSessionInformationSourcetoTargetList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationSourcetoTargetList(data []byte) (MBSActiveSessionInformationSourcetoTargetList, int, error) {
	var v MBSActiveSessionInformationSourcetoTargetList
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-SourcetoTargetList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationSourcetoTargetList{}, 0, err
	}
	return v, len(data), err
}
