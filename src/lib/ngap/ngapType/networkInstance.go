// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NetworkInstance represents ASN.1 NetworkInstance.
type NetworkInstance struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkInstance) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkInstance", v)
}

// DecodeNetworkInstance decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkInstance(data []byte) (NetworkInstance, int, error) {
	var v NetworkInstance
	err := ngapcodec.DecodeParameter("NetworkInstance", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkInstance{}, 0, err
	}
	return v, len(data), err
}
