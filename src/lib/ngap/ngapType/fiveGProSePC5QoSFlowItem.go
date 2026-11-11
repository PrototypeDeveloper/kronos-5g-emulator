// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSePC5QoSFlowItem represents ASN.1 FiveGProSePC5QoSFlowItem.
type FiveGProSePC5QoSFlowItem struct {
	FiveGproSepQI                FiveQI
	HasFiveGproSepc5FlowBitRates bool
	FiveGproSepc5FlowBitRates    FiveGProSePC5FlowBitRates
	HasFiveGproSerange           bool
	FiveGproSerange              Range
	HasIEExtensions              bool
	IEExtensions                 []FiveGProSePC5QoSFlowItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5QoSFlowItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSePC5QoSFlowItem", v)
}

// DecodeFiveGProSePC5QoSFlowItem decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5QoSFlowItem(data []byte) (FiveGProSePC5QoSFlowItem, int, error) {
	var v FiveGProSePC5QoSFlowItem
	err := ngapcodec.DecodeParameter("FiveGProSePC5QoSFlowItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5QoSFlowItem{}, 0, err
	}
	return v, len(data), err
}
