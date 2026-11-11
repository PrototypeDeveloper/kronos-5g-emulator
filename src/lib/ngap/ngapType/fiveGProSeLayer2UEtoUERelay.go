// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2UEtoUERelay represents ASN.1 FiveGProSeLayer2UEtoUERelay.
const (
	FiveGProSeLayer2UEtoUERelayPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2UEtoUERelayPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2UEtoUERelay struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2UEtoUERelay) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2UEtoUERelay", v)
}

// DecodeFiveGProSeLayer2UEtoUERelay decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2UEtoUERelay(data []byte) (FiveGProSeLayer2UEtoUERelay, int, error) {
	var v FiveGProSeLayer2UEtoUERelay
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2UEtoUERelay", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2UEtoUERelay{}, 0, err
	}
	return v, len(data), err
}
