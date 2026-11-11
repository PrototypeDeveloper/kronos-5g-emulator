// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NetworkSliceItemforMDT represents ASN.1 NetworkSliceItemforMDT.
type NetworkSliceItemforMDT struct {
	PlmnID            PLMNIdentity
	SliceMDTList      SliceMDTList
	HasIEExtensions   bool
	IEExtensions      []NetworkSliceItemforMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkSliceItemforMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkSliceItemforMDT", v)
}

// DecodeNetworkSliceItemforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkSliceItemforMDT(data []byte) (NetworkSliceItemforMDT, int, error) {
	var v NetworkSliceItemforMDT
	err := ngapcodec.DecodeParameter("NetworkSliceItemforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkSliceItemforMDT{}, 0, err
	}
	return v, len(data), err
}
