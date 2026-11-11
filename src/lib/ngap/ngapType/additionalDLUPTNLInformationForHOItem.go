// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalDLUPTNLInformationForHOItem represents ASN.1 AdditionalDLUPTNLInformationForHOItem.
type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation           UPTransportLayerInformation
	AdditionalQosFlowSetupResponseList        QosFlowListWithDataForwarding
	HasAdditionalDLForwardingUPTNLInformation bool
	AdditionalDLForwardingUPTNLInformation    UPTransportLayerInformation
	HasIEExtensions                           bool
	IEExtensions                              []AdditionalDLUPTNLInformationForHOItemExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalDLUPTNLInformationForHOItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalDLUPTNLInformationForHOItem", v)
}

// DecodeAdditionalDLUPTNLInformationForHOItem decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalDLUPTNLInformationForHOItem(data []byte) (AdditionalDLUPTNLInformationForHOItem, int, error) {
	var v AdditionalDLUPTNLInformationForHOItem
	err := ngapcodec.DecodeParameter("AdditionalDLUPTNLInformationForHOItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalDLUPTNLInformationForHOItem{}, 0, err
	}
	return v, len(data), err
}
