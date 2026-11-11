// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsToQosFlowsMappingList represents ASN.1 DRBsToQosFlowsMappingList.
type DRBsToQosFlowsMappingList struct {
	List []DRBsToQosFlowsMappingItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsToQosFlowsMappingList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsToQosFlowsMappingList", v)
}

// DecodeDRBsToQosFlowsMappingList decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsToQosFlowsMappingList(data []byte) (DRBsToQosFlowsMappingList, int, error) {
	var v DRBsToQosFlowsMappingList
	err := ngapcodec.DecodeParameter("DRBsToQosFlowsMappingList", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsToQosFlowsMappingList{}, 0, err
	}
	return v, len(data), err
}
