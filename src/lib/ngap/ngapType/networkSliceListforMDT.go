// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NetworkSliceListforMDT represents ASN.1 NetworkSliceListforMDT.
type NetworkSliceListforMDT struct {
	List []NetworkSliceItemforMDT `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkSliceListforMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkSliceListforMDT", v)
}

// DecodeNetworkSliceListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkSliceListforMDT(data []byte) (NetworkSliceListforMDT, int, error) {
	var v NetworkSliceListforMDT
	err := ngapcodec.DecodeParameter("NetworkSliceListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkSliceListforMDT{}, 0, err
	}
	return v, len(data), err
}
