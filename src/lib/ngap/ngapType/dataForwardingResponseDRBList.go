// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingResponseDRBList represents ASN.1 DataForwardingResponseDRBList.
type DataForwardingResponseDRBList struct {
	List []DataForwardingResponseDRBItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseDRBList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseDRBList", v)
}

// DecodeDataForwardingResponseDRBList decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseDRBList(data []byte) (DataForwardingResponseDRBList, int, error) {
	var v DataForwardingResponseDRBList
	err := ngapcodec.DecodeParameter("DataForwardingResponseDRBList", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseDRBList{}, 0, err
	}
	return v, len(data), err
}
