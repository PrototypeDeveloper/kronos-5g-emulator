// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PedestrianUE represents ASN.1 PedestrianUE.
const (
	PedestrianUEPresentAuthorized    Enumerated = 0
	PedestrianUEPresentNotAuthorized Enumerated = 1
)

type PedestrianUE struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PedestrianUE) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PedestrianUE", v) }

// DecodePedestrianUE decodes a standalone APER value and reports the consumed octets.
func DecodePedestrianUE(data []byte) (PedestrianUE, int, error) {
	var v PedestrianUE
	err := ngapcodec.DecodeParameter("PedestrianUE", data, &v)
	if ngapcodec.IsFatal(err) {
		return PedestrianUE{}, 0, err
	}
	return v, len(data), err
}
