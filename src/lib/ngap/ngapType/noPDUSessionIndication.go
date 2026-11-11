// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NoPDUSessionIndication represents ASN.1 NoPDUSessionIndication.
const (
	NoPDUSessionIndicationPresentTrue Enumerated = 0
)

type NoPDUSessionIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NoPDUSessionIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NoPDUSessionIndication", v)
}

// DecodeNoPDUSessionIndication decodes a standalone APER value and reports the consumed octets.
func DecodeNoPDUSessionIndication(data []byte) (NoPDUSessionIndication, int, error) {
	var v NoPDUSessionIndication
	err := ngapcodec.DecodeParameter("NoPDUSessionIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return NoPDUSessionIndication{}, 0, err
	}
	return v, len(data), err
}
