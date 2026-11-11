// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GWContextReleaseIndication represents ASN.1 GW-Context-Release-Indication.
const (
	GWContextReleaseIndicationPresentTrue Enumerated = 0
)

type GWContextReleaseIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GWContextReleaseIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GW-Context-Release-Indication", v)
}

// DecodeGWContextReleaseIndication decodes a standalone APER value and reports the consumed octets.
func DecodeGWContextReleaseIndication(data []byte) (GWContextReleaseIndication, int, error) {
	var v GWContextReleaseIndication
	err := ngapcodec.DecodeParameter("GW-Context-Release-Indication", data, &v)
	if ngapcodec.IsFatal(err) {
		return GWContextReleaseIndication{}, 0, err
	}
	return v, len(data), err
}
