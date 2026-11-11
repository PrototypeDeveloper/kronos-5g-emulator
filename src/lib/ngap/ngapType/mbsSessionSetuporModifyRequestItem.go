// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetuporModifyRequestItem represents ASN.1 MBSSessionSetuporModifyRequestItem.
type MBSSessionSetuporModifyRequestItem struct {
	MBSSessionID                                    MBSSessionID
	HasMBSAreaSessionID                             bool
	MBSAreaSessionID                                MBSAreaSessionID
	HasAssociatedMBSQosFlowSetuporModifyRequestList bool
	AssociatedMBSQosFlowSetuporModifyRequestList    AssociatedMBSQosFlowSetuporModifyRequestList
	HasMBSQosFlowToReleaseList                      bool
	MBSQosFlowToReleaseList                         QosFlowListWithCause
	HasIEExtensions                                 bool
	IEExtensions                                    []MBSSessionSetuporModifyRequestItemExtIEsEntry
	UnknownExtensions                               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetuporModifyRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetuporModifyRequestItem", v)
}

// DecodeMBSSessionSetuporModifyRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetuporModifyRequestItem(data []byte) (MBSSessionSetuporModifyRequestItem, int, error) {
	var v MBSSessionSetuporModifyRequestItem
	err := ngapcodec.DecodeParameter("MBSSessionSetuporModifyRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetuporModifyRequestItem{}, 0, err
	}
	return v, len(data), err
}
