// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSActiveSessionInformationTargettoSourceItem represents ASN.1 MBS-ActiveSessionInformation-TargettoSourceItem.
type MBSActiveSessionInformationTargettoSourceItem struct {
	MBSSessionID                        MBSSessionID
	HasMBSDataForwardingResponseMRBList bool
	MBSDataForwardingResponseMRBList    MBSDataForwardingResponseMRBList
	HasIEExtensions                     bool
	IEExtensions                        []MBSActiveSessionInformationTargettoSourceItemExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationTargettoSourceItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-TargettoSourceItem", v)
}

// DecodeMBSActiveSessionInformationTargettoSourceItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationTargettoSourceItem(data []byte) (MBSActiveSessionInformationTargettoSourceItem, int, error) {
	var v MBSActiveSessionInformationTargettoSourceItem
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-TargettoSourceItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationTargettoSourceItem{}, 0, err
	}
	return v, len(data), err
}
