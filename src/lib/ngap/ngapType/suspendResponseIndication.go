// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuspendResponseIndication represents ASN.1 Suspend-Response-Indication.
const (
	SuspendResponseIndicationPresentSuspendIndicated Enumerated = 0
)

type SuspendResponseIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuspendResponseIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Suspend-Response-Indication", v)
}

// DecodeSuspendResponseIndication decodes a standalone APER value and reports the consumed octets.
func DecodeSuspendResponseIndication(data []byte) (SuspendResponseIndication, int, error) {
	var v SuspendResponseIndication
	err := ngapcodec.DecodeParameter("Suspend-Response-Indication", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuspendResponseIndication{}, 0, err
	}
	return v, len(data), err
}
