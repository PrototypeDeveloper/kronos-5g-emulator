// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDataForwardingResponseMRBItem represents ASN.1 MBS-DataForwardingResponseMRBItem.
type MBSDataForwardingResponseMRBItem struct {
	MRBID                        MRBID
	DLForwardingUPTNLInformation UPTransportLayerInformation
	HasMRBProgressInformation    bool
	MRBProgressInformation       MRBProgressInformation
	HasIEExtensions              bool
	IEExtensions                 []MBSDataForwardingResponseMRBItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDataForwardingResponseMRBItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DataForwardingResponseMRBItem", v)
}

// DecodeMBSDataForwardingResponseMRBItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDataForwardingResponseMRBItem(data []byte) (MBSDataForwardingResponseMRBItem, int, error) {
	var v MBSDataForwardingResponseMRBItem
	err := ngapcodec.DecodeParameter("MBS-DataForwardingResponseMRBItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDataForwardingResponseMRBItem{}, 0, err
	}
	return v, len(data), err
}
