// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSeDirectCommunication represents ASN.1 FiveGProSeDirectCommunication.
const (
	FiveGProSeDirectCommunicationPresentAuthorized    Enumerated = 0
	FiveGProSeDirectCommunicationPresentNotAuthorized Enumerated = 1
)

type FiveGProSeDirectCommunication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSeDirectCommunication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSeDirectCommunication", v)
}

// DecodeFiveGProSeDirectCommunication decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSeDirectCommunication(data []byte) (FiveGProSeDirectCommunication, int, error) {
	var v FiveGProSeDirectCommunication
	err := ngapcodec.DecodeParameter("FiveGProSeDirectCommunication", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSeDirectCommunication{}, 0, err
	}
	return v, len(data), err
}
