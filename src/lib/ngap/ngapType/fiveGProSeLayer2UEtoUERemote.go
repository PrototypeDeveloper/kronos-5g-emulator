// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2UEtoUERemote represents ASN.1 FiveGProSeLayer2UEtoUERemote.
const (
	FiveGProSeLayer2UEtoUERemotePresentAuthorized    Enumerated = 0
	FiveGProSeLayer2UEtoUERemotePresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2UEtoUERemote struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2UEtoUERemote) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2UEtoUERemote", v)
}

// DecodeFiveGProSeLayer2UEtoUERemote decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2UEtoUERemote(data []byte) (FiveGProSeLayer2UEtoUERemote, int, error) {
	var v FiveGProSeLayer2UEtoUERemote
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2UEtoUERemote", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2UEtoUERemote{}, 0, err
	}
	return v, len(data), err
}
