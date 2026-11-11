// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UPTransportLayerInformationList represents ASN.1 UPTransportLayerInformationList.
type UPTransportLayerInformationList struct {
	List []UPTransportLayerInformationItem `aper:"sizeLB:1,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationList", v)
}

// DecodeUPTransportLayerInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationList(data []byte) (UPTransportLayerInformationList, int, error) {
	var v UPTransportLayerInformationList
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationList{}, 0, err
	}
	return v, len(data), err
}
