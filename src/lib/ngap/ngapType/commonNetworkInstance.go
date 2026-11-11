// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommonNetworkInstance represents ASN.1 CommonNetworkInstance.
type CommonNetworkInstance struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommonNetworkInstance) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommonNetworkInstance", v)
}

// DecodeCommonNetworkInstance decodes a standalone APER value and reports the consumed octets.
func DecodeCommonNetworkInstance(data []byte) (CommonNetworkInstance, int, error) {
	var v CommonNetworkInstance
	err := ngapcodec.DecodeParameter("CommonNetworkInstance", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommonNetworkInstance{}, 0, err
	}
	return v, len(data), err
}
