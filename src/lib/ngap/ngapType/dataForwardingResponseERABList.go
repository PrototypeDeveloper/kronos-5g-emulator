// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingResponseERABList represents ASN.1 DataForwardingResponseERABList.
type DataForwardingResponseERABList struct {
	List []DataForwardingResponseERABListItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseERABList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseERABList", v)
}

// DecodeDataForwardingResponseERABList decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseERABList(data []byte) (DataForwardingResponseERABList, int, error) {
	var v DataForwardingResponseERABList
	err := ngapcodec.DecodeParameter("DataForwardingResponseERABList", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseERABList{}, 0, err
	}
	return v, len(data), err
}
