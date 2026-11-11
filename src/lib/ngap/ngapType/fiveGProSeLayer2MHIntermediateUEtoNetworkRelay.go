// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2MHIntermediateUEtoNetworkRelay represents ASN.1 FiveGProSeLayer2MHIntermediateUEtoNetworkRelay.
const (
	FiveGProSeLayer2MHIntermediateUEtoNetworkRelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2MHIntermediateUEtoNetworkRelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2MHIntermediateUEtoNetworkRelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2MHIntermediateUEtoNetworkRelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2MHIntermediateUEtoNetworkRelay", v)
}

// DecodeFiveGProSeLayer2MHIntermediateUEtoNetworkRelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2MHIntermediateUEtoNetworkRelay(data []byte) (FiveGProSeLayer2MHIntermediateUEtoNetworkRelay, int, error) {
	var v FiveGProSeLayer2MHIntermediateUEtoNetworkRelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2MHIntermediateUEtoNetworkRelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2MHIntermediateUEtoNetworkRelay{}, 0, err
	}
	return v, len(data), err
}
