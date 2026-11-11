// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2MHUEtoNetworkRelay represents ASN.1 FiveGProSeLayer2MHUEtoNetworkRelay.
const (
	FiveGProSeLayer2MHUEtoNetworkRelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2MHUEtoNetworkRelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2MHUEtoNetworkRelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2MHUEtoNetworkRelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2MHUEtoNetworkRelay", v)
}

// DecodeFiveGProSeLayer2MHUEtoNetworkRelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2MHUEtoNetworkRelay(data []byte) (FiveGProSeLayer2MHUEtoNetworkRelay, int, error) {
	var v FiveGProSeLayer2MHUEtoNetworkRelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2MHUEtoNetworkRelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2MHUEtoNetworkRelay{}, 0, err
	}
	return v, len(data), err
}
