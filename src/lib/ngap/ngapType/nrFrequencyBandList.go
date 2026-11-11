// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRFrequencyBandList represents ASN.1 NRFrequencyBand-List.
type NRFrequencyBandList struct {
	List []NRFrequencyBandItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyBandList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyBand-List", v)
}

// DecodeNRFrequencyBandList decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyBandList(data []byte) (NRFrequencyBandList, int, error) {
	var v NRFrequencyBandList
	err := ngapcodec.DecodeParameter("NRFrequencyBand-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyBandList{}, 0, err
	}
	return v, len(data), err
}
