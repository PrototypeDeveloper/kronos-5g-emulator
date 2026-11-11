// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupRequestItem represents ASN.1 MBSSessionSetupRequestItem.
type MBSSessionSetupRequestItem struct {
	MBSSessionID                            MBSSessionID
	HasMBSAreaSessionID                     bool
	MBSAreaSessionID                        MBSAreaSessionID
	HasAssociatedMBSQosFlowSetupRequestList bool
	AssociatedMBSQosFlowSetupRequestList    AssociatedMBSQosFlowSetupRequestList
	HasIEExtensions                         bool
	IEExtensions                            []MBSSessionSetupRequestItemExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupRequestItem", v)
}

// DecodeMBSSessionSetupRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupRequestItem(data []byte) (MBSSessionSetupRequestItem, int, error) {
	var v MBSSessionSetupRequestItem
	err := ngapcodec.DecodeParameter("MBSSessionSetupRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupRequestItem{}, 0, err
	}
	return v, len(data), err
}
