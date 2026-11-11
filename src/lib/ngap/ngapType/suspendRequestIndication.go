// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SuspendRequestIndication represents ASN.1 Suspend-Request-Indication.
const (
	SuspendRequestIndicationPresentSuspendRequested Enumerated = 0
)

type SuspendRequestIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuspendRequestIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Suspend-Request-Indication", v)
}

// DecodeSuspendRequestIndication decodes a standalone APER value and reports the consumed octets.
func DecodeSuspendRequestIndication(data []byte) (SuspendRequestIndication, int, error) {
	var v SuspendRequestIndication
	err := ngapcodec.DecodeParameter("Suspend-Request-Indication", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuspendRequestIndication{}, 0, err
	}
	return v, len(data), err
}
