// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CauseTransport represents ASN.1 CauseTransport.
const (
	CauseTransportPresentTransportResourceUnavailable Enumerated = 0
	CauseTransportPresentUnspecified                  Enumerated = 1
)

type CauseTransport struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseTransport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CauseTransport", v)
}

// DecodeCauseTransport decodes a standalone APER value and reports the consumed octets.
func DecodeCauseTransport(data []byte) (CauseTransport, int, error) {
	var v CauseTransport
	err := ngapcodec.DecodeParameter("CauseTransport", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseTransport{}, 0, err
	}
	return v, len(data), err
}
