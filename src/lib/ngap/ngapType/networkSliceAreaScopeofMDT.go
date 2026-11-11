// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NetworkSliceAreaScopeofMDT represents ASN.1 NetworkSliceAreaScopeofMDT.
type NetworkSliceAreaScopeofMDT struct {
	NetworkSliceListforMDT NetworkSliceListforMDT
	HasIEExtensions        bool
	IEExtensions           []NetworkSliceAreaScopeofMDTExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkSliceAreaScopeofMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkSliceAreaScopeofMDT", v)
}

// DecodeNetworkSliceAreaScopeofMDT decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkSliceAreaScopeofMDT(data []byte) (NetworkSliceAreaScopeofMDT, int, error) {
	var v NetworkSliceAreaScopeofMDT
	err := ngapcodec.DecodeParameter("NetworkSliceAreaScopeofMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkSliceAreaScopeofMDT{}, 0, err
	}
	return v, len(data), err
}
