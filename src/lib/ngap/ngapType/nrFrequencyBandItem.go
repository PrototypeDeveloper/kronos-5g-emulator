// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRFrequencyBandItem represents ASN.1 NRFrequencyBandItem.
type NRFrequencyBandItem struct {
	NrFrequencyBand   NRFrequencyBand
	HasIEExtension    bool
	IEExtension       []NRFrequencyBandItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyBandItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyBandItem", v)
}

// DecodeNRFrequencyBandItem decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyBandItem(data []byte) (NRFrequencyBandItem, int, error) {
	var v NRFrequencyBandItem
	err := ngapcodec.DecodeParameter("NRFrequencyBandItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyBandItem{}, 0, err
	}
	return v, len(data), err
}
