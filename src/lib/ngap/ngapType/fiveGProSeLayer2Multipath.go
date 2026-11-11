// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2Multipath represents ASN.1 FiveGProSeLayer2Multipath.
const (
	FiveGProSeLayer2MultipathPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2MultipathPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2Multipath struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2Multipath) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2Multipath", v)
}

// DecodeFiveGProSeLayer2Multipath decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2Multipath(data []byte) (FiveGProSeLayer2Multipath, int, error) {
	var v FiveGProSeLayer2Multipath
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2Multipath", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2Multipath{}, 0, err
	}
	return v, len(data), err
}
