// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeDirectDiscovery represents ASN.1 FiveGProSeDirectDiscovery.
const (
	FiveGProSeDirectDiscoveryPresentAuthorized    Enumerated = 0
	FiveGProSeDirectDiscoveryPresentNotAuthorized Enumerated = 1
)

type FiveGProSeDirectDiscovery struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeDirectDiscovery) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeDirectDiscovery", v)
}

// DecodeFiveGProSeDirectDiscovery decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeDirectDiscovery(data []byte) (FiveGProSeDirectDiscovery, int, error) {
	var v FiveGProSeDirectDiscovery
	err := ngapcodec.DecodeParameter("FiveGProSeDirectDiscovery", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeDirectDiscovery{}, 0, err
	}
	return v, len(data), err
}
