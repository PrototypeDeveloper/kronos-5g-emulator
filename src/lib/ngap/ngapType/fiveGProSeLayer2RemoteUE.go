// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeLayer2RemoteUE represents ASN.1 FiveGProSeLayer2RemoteUE.
const (
	FiveGProSeLayer2RemoteUEPresentAuthorized    Enumerated = 0
	FiveGProSeLayer2RemoteUEPresentNotAuthorized Enumerated = 1
)

type FiveGProSeLayer2RemoteUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeLayer2RemoteUE) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeLayer2RemoteUE", v)
}

// DecodeFiveGProSeLayer2RemoteUE decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeLayer2RemoteUE(data []byte) (FiveGProSeLayer2RemoteUE, int, error) {
	var v FiveGProSeLayer2RemoteUE
	err := ngapcodec.DecodeParameter("FiveGProSeLayer2RemoteUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeLayer2RemoteUE{}, 0, err
	}
	return v, len(data), err
}
