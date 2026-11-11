// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RSPPQoSFlowItem represents ASN.1 RSPPQoSFlowItem.
type RSPPQoSFlowItem struct {
	PQI                 FiveQI
	HasRSPPFlowBitRates bool
	RSPPFlowBitRates    RSPPFlowBitRates
	HasRange            bool
	Range               Range
	HasIEExtensions     bool
	IEExtensions        []RSPPQoSFlowItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSPPQoSFlowItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RSPPQoSFlowItem", v)
}

// DecodeRSPPQoSFlowItem decodes a standalone APER value and reports the consumed octets.
func DecodeRSPPQoSFlowItem(data []byte) (RSPPQoSFlowItem, int, error) {
	var v RSPPQoSFlowItem
	err := ngapcodec.DecodeParameter("RSPPQoSFlowItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSPPQoSFlowItem{}, 0, err
	}
	return v, len(data), err
}
