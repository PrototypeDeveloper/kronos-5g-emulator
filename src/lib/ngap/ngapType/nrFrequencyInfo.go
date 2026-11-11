// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRFrequencyInfo represents ASN.1 NRFrequencyInfo.
type NRFrequencyInfo struct {
	NrARFCN           NRARFCN
	FrequencyBandList NRFrequencyBandList
	HasIEExtension    bool
	IEExtension       []NRFrequencyInfoExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRFrequencyInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRFrequencyInfo", v)
}

// DecodeNRFrequencyInfo decodes a standalone APER value and reports the consumed octets.
func DecodeNRFrequencyInfo(data []byte) (NRFrequencyInfo, int, error) {
	var v NRFrequencyInfo
	err := ngapcodec.DecodeParameter("NRFrequencyInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRFrequencyInfo{}, 0, err
	}
	return v, len(data), err
}
