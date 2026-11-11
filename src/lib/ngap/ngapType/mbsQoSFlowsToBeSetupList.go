// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSQoSFlowsToBeSetupList represents ASN.1 MBS-QoSFlowsToBeSetupList.
type MBSQoSFlowsToBeSetupList struct {
	List []MBSQoSFlowsToBeSetupItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSQoSFlowsToBeSetupList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-QoSFlowsToBeSetupList", v)
}

// DecodeMBSQoSFlowsToBeSetupList decodes a standalone APER value and reports the consumed octets.
func DecodeMBSQoSFlowsToBeSetupList(data []byte) (MBSQoSFlowsToBeSetupList, int, error) {
	var v MBSQoSFlowsToBeSetupList
	err := ngapcodec.DecodeParameter("MBS-QoSFlowsToBeSetupList", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSQoSFlowsToBeSetupList{}, 0, err
	}
	return v, len(data), err
}
