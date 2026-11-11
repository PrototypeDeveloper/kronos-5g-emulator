// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEMIndication represents ASN.1 LTEM-Indication.
const (
	LTEMIndicationPresentLteM Enumerated = 0
)

type LTEMIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEMIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEM-Indication", v)
}

// DecodeLTEMIndication decodes a standalone APER value and reports the consumed octets.
func DecodeLTEMIndication(data []byte) (LTEMIndication, int, error) {
	var v LTEMIndication
	err := ngapcodec.DecodeParameter("LTEM-Indication", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEMIndication{}, 0, err
	}
	return v, len(data), err
}
