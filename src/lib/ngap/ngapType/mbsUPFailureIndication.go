// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSUPFailureIndication represents ASN.1 MBS-UP-FailureIndication.
const (
	MBSUPFailureIndicationPresentNgUPathFailureDetected Enumerated = 0
)

type MBSUPFailureIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSUPFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-UP-FailureIndication", v)
}

// DecodeMBSUPFailureIndication decodes a standalone APER value and reports the consumed octets.
func DecodeMBSUPFailureIndication(data []byte) (MBSUPFailureIndication, int, error) {
	var v MBSUPFailureIndication
	err := ngapcodec.DecodeParameter("MBS-UP-FailureIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSUPFailureIndication{}, 0, err
	}
	return v, len(data), err
}
