// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DirectForwardingPathAvailability represents ASN.1 DirectForwardingPathAvailability.
const (
	DirectForwardingPathAvailabilityPresentDirectPathAvailable Enumerated = 0
)

type DirectForwardingPathAvailability struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DirectForwardingPathAvailability) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DirectForwardingPathAvailability", v)
}

// DecodeDirectForwardingPathAvailability decodes a standalone APER value and reports the consumed octets.
func DecodeDirectForwardingPathAvailability(data []byte) (DirectForwardingPathAvailability, int, error) {
	var v DirectForwardingPathAvailability
	err := ngapcodec.DecodeParameter("DirectForwardingPathAvailability", data, &v)
	if ngapcodec.IsFatal(err) {
		return DirectForwardingPathAvailability{}, 0, err
	}
	return v, len(data), err
}
