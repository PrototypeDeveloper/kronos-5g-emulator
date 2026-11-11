// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2UEtoNetworkRelay represents ASN.1 FiveGProSeLayer2UEtoNetworkRelay.
const (
	FiveGProSeLayer2UEtoNetworkRelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2UEtoNetworkRelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2UEtoNetworkRelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2UEtoNetworkRelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2UEtoNetworkRelay", v)
}

// DecodeFiveGProSeLayer2UEtoNetworkRelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2UEtoNetworkRelay(data []byte) (FiveGProSeLayer2UEtoNetworkRelay, int, error) {
	var v FiveGProSeLayer2UEtoNetworkRelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2UEtoNetworkRelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2UEtoNetworkRelay{}, 0, err
	}
	return v, len(data), err
}
