// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UPTransportLayerInformationPairList represents ASN.1 UPTransportLayerInformationPairList.
type UPTransportLayerInformationPairList struct {
	List []UPTransportLayerInformationPairItem `aper:"sizeLB:1,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationPairList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationPairList", v)
}

// DecodeUPTransportLayerInformationPairList decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationPairList(data []byte) (UPTransportLayerInformationPairList, int, error) {
	var v UPTransportLayerInformationPairList
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationPairList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationPairList{}, 0, err
	}
	return v, len(data), err
}
