// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSMappingandDataForwardingRequestItem represents ASN.1 MBS-MappingandDataForwardingRequestItem.
type MBSMappingandDataForwardingRequestItem struct {
	MRBID                     MRBID
	MBSQoSFlowList            MBSQoSFlowList
	HasMRBProgressInformation bool
	MRBProgressInformation    MRBProgressInformation
	HasIEExtensions           bool
	IEExtensions              []MBSMappingandDataForwardingRequestItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSMappingandDataForwardingRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-MappingandDataForwardingRequestItem", v)
}

// DecodeMBSMappingandDataForwardingRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSMappingandDataForwardingRequestItem(data []byte) (MBSMappingandDataForwardingRequestItem, int, error) {
	var v MBSMappingandDataForwardingRequestItem
	err := ngapcodec.DecodeParameter("MBS-MappingandDataForwardingRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSMappingandDataForwardingRequestItem{}, 0, err
	}
	return v, len(data), err
}
