// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NetworkControlledRepeaterAuthorized represents ASN.1 NetworkControlledRepeaterAuthorized.
const (
	NetworkControlledRepeaterAuthorizedPresentAuthorized    Enumerated = 0
	NetworkControlledRepeaterAuthorizedPresentNotAuthorized Enumerated = 1
)

type NetworkControlledRepeaterAuthorized struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkControlledRepeaterAuthorized) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkControlledRepeaterAuthorized", v)
}

// DecodeNetworkControlledRepeaterAuthorized decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkControlledRepeaterAuthorized(data []byte) (NetworkControlledRepeaterAuthorized, int, error) {
	var v NetworkControlledRepeaterAuthorized
	err := ngapcodec.DecodeParameter("NetworkControlledRepeaterAuthorized", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkControlledRepeaterAuthorized{}, 0, err
	}
	return v, len(data), err
}
