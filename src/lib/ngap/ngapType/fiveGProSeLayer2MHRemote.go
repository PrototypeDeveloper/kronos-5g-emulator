// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2MHRemote represents ASN.1 FiveGProSeLayer2MHRemote.
const (
	FiveGProSeLayer2MHRemotePresentAuthorized    Enumerated = 0
	FiveGProSeLayer2MHRemotePresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2MHRemote struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2MHRemote) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2MHRemote", v)
}

// DecodeFiveGProSeLayer2MHRemote decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2MHRemote(data []byte) (FiveGProSeLayer2MHRemote, int, error) {
	var v FiveGProSeLayer2MHRemote
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2MHRemote", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2MHRemote{}, 0, err
	}
	return v, len(data), err
}
