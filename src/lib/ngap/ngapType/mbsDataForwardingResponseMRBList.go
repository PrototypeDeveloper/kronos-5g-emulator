// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDataForwardingResponseMRBList represents ASN.1 MBS-DataForwardingResponseMRBList.
type MBSDataForwardingResponseMRBList struct {
	List []MBSDataForwardingResponseMRBItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDataForwardingResponseMRBList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DataForwardingResponseMRBList", v)
}

// DecodeMBSDataForwardingResponseMRBList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDataForwardingResponseMRBList(data []byte) (MBSDataForwardingResponseMRBList, int, error) {
	var v MBSDataForwardingResponseMRBList
	err := ngapcodec.DecodeParameter("MBS-DataForwardingResponseMRBList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDataForwardingResponseMRBList{}, 0, err
	}
	return v, len(data), err
}
