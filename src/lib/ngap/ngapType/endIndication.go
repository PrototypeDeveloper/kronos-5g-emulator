// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EndIndication represents ASN.1 EndIndication.
const (
	EndIndicationPresentNoFurtherData     Enumerated = 0
	EndIndicationPresentFurtherDataExists Enumerated = 1
)

type EndIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EndIndication) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EndIndication", v) }

// DecodeEndIndication decodes a standalone APER value and reports the consumed octets.
func DecodeEndIndication(data []byte) (EndIndication, int, error) {
	var v EndIndication
	err := ngapcodec.DecodeParameter("EndIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return EndIndication{}, 0, err
	}
	return v, len(data), err
}
