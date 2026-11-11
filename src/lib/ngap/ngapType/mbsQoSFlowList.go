// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSQoSFlowList represents ASN.1 MBS-QoSFlowList.
type MBSQoSFlowList struct {
	List []QosFlowIdentifier `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSQoSFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-QoSFlowList", v)
}

// DecodeMBSQoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSQoSFlowList(data []byte) (MBSQoSFlowList, int, error) {
	var v MBSQoSFlowList
	err := ngapcodec.DecodeParameter("MBS-QoSFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSQoSFlowList{}, 0, err
	}
	return v, len(data), err
}
