// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSMappingandDataForwardingRequestList represents ASN.1 MBS-MappingandDataForwardingRequestList.
type MBSMappingandDataForwardingRequestList struct {
	List []MBSMappingandDataForwardingRequestItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSMappingandDataForwardingRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-MappingandDataForwardingRequestList", v)
}

// DecodeMBSMappingandDataForwardingRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSMappingandDataForwardingRequestList(data []byte) (MBSMappingandDataForwardingRequestList, int, error) {
	var v MBSMappingandDataForwardingRequestList
	err := ngapcodec.DecodeParameter("MBS-MappingandDataForwardingRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSMappingandDataForwardingRequestList{}, 0, err
	}
	return v, len(data), err
}
