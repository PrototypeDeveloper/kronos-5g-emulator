// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSQoSFlowsToBeSetupItem represents ASN.1 MBS-QoSFlowsToBeSetupItem.
type MBSQoSFlowsToBeSetupItem struct {
	MBSqosFlowIdentifier         QosFlowIdentifier
	MBSqosFlowLevelQosParameters QosFlowLevelQosParameters
	HasIEExtensions              bool
	IEExtensions                 []MBSQoSFlowsToBeSetupItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSQoSFlowsToBeSetupItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-QoSFlowsToBeSetupItem", v)
}

// DecodeMBSQoSFlowsToBeSetupItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSQoSFlowsToBeSetupItem(data []byte) (MBSQoSFlowsToBeSetupItem, int, error) {
	var v MBSQoSFlowsToBeSetupItem
	err := ngapcodec.DecodeParameter("MBS-QoSFlowsToBeSetupItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSQoSFlowsToBeSetupItem{}, 0, err
	}
	return v, len(data), err
}
