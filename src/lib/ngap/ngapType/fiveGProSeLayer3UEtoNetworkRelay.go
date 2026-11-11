// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer3UEtoNetworkRelay represents ASN.1 FiveGProSeLayer3UEtoNetworkRelay.
const (
	FiveGProSeLayer3UEtoNetworkRelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer3UEtoNetworkRelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer3UEtoNetworkRelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer3UEtoNetworkRelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer3UEtoNetworkRelay", v)
}

// DecodeFiveGProSeLayer3UEtoNetworkRelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer3UEtoNetworkRelay(data []byte) (FiveGProSeLayer3UEtoNetworkRelay, int, error) {
	var v FiveGProSeLayer3UEtoNetworkRelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer3UEtoNetworkRelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer3UEtoNetworkRelay{}, 0, err
	}
	return v, len(data), err
}
