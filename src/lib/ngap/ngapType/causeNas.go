// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CauseNas represents ASN.1 CauseNas.
const (
	CauseNasPresentNormalRelease          Enumerated = 0
	CauseNasPresentAuthenticationFailure  Enumerated = 1
	CauseNasPresentDeregister             Enumerated = 2
	CauseNasPresentUnspecified            Enumerated = 3
	CauseNasPresentUENotInPLMNServingArea Enumerated = 4
	CauseNasPresentMobileIABNotAuthorized Enumerated = 5
	CauseNasPresentIABNotAuthorized       Enumerated = 6
	CauseNasPresentNoAiotNASResponse      Enumerated = 7
)

type CauseNas struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseNas) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CauseNas", v) }

// DecodeCauseNas decodes a standalone APER value and reports the consumed octets.
func DecodeCauseNas(data []byte) (CauseNas, int, error) {
	var v CauseNas
	err := ngapcodec.DecodeParameter("CauseNas", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseNas{}, 0, err
	}
	return v, len(data), err
}
