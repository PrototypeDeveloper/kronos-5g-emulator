// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer3MHUEtoNetworkRelay represents ASN.1 FiveGProSeLayer3MHUEtoNetworkRelay.
const (
	FiveGProSeLayer3MHUEtoNetworkRelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer3MHUEtoNetworkRelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer3MHUEtoNetworkRelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer3MHUEtoNetworkRelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer3MHUEtoNetworkRelay", v)
}

// DecodeFiveGProSeLayer3MHUEtoNetworkRelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer3MHUEtoNetworkRelay(data []byte) (FiveGProSeLayer3MHUEtoNetworkRelay, int, error) {
	var v FiveGProSeLayer3MHUEtoNetworkRelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer3MHUEtoNetworkRelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer3MHUEtoNetworkRelay{}, 0, err
	}
	return v, len(data), err
}
