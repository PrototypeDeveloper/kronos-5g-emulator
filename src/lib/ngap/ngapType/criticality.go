// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Criticality represents ASN.1 Criticality.
const (
	CriticalityPresentReject Enumerated = 0
	CriticalityPresentIgnore Enumerated = 1
	CriticalityPresentNotify Enumerated = 2
)

type Criticality struct {
	Value Enumerated `aper:"valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Criticality) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Criticality", v) }

// DecodeCriticality decodes a standalone APER value and reports the consumed octets.
func DecodeCriticality(data []byte) (Criticality, int, error) {
	var v Criticality
	err := ngapcodec.DecodeParameter("Criticality", data, &v)
	if ngapcodec.IsFatal(err) {
		return Criticality{}, 0, err
	}
	return v, len(data), err
}
