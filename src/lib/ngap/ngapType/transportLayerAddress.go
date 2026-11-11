// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TransportLayerAddress represents ASN.1 TransportLayerAddress.
type TransportLayerAddress struct {
	Value BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TransportLayerAddress) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TransportLayerAddress", v)
}

// DecodeTransportLayerAddress decodes a standalone APER value and reports the consumed octets.
func DecodeTransportLayerAddress(data []byte) (TransportLayerAddress, int, error) {
	var v TransportLayerAddress
	err := ngapcodec.DecodeParameter("TransportLayerAddress", data, &v)
	if ngapcodec.IsFatal(err) {
		return TransportLayerAddress{}, 0, err
	}
	return v, len(data), err
}
