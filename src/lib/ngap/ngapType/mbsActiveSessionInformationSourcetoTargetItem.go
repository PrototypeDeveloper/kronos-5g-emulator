// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSActiveSessionInformationSourcetoTargetItem represents ASN.1 MBS-ActiveSessionInformation-SourcetoTargetItem.
type MBSActiveSessionInformationSourcetoTargetItem struct {
	MBSSessionID                              MBSSessionID
	HasMBSAreaSessionID                       bool
	MBSAreaSessionID                          MBSAreaSessionID
	HasMBSServiceArea                         bool
	MBSServiceArea                            MBSServiceArea
	MBSQoSFlowsToBeSetupList                  MBSQoSFlowsToBeSetupList
	HasMBSMappingandDataForwardingRequestList bool
	MBSMappingandDataForwardingRequestList    MBSMappingandDataForwardingRequestList
	HasIEExtensions                           bool
	IEExtensions                              []MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationSourcetoTargetItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-SourcetoTargetItem", v)
}

// DecodeMBSActiveSessionInformationSourcetoTargetItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationSourcetoTargetItem(data []byte) (MBSActiveSessionInformationSourcetoTargetItem, int, error) {
	var v MBSActiveSessionInformationSourcetoTargetItem
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-SourcetoTargetItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationSourcetoTargetItem{}, 0, err
	}
	return v, len(data), err
}
