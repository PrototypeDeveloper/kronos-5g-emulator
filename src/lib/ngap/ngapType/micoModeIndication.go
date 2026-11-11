// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MICOModeIndication represents ASN.1 MICOModeIndication.
const (
	MICOModeIndicationPresentTrue Enumerated = 0
)

type MICOModeIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MICOModeIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MICOModeIndication", v)
}

// DecodeMICOModeIndication decodes a standalone APER value and reports the consumed octets.
func DecodeMICOModeIndication(data []byte) (MICOModeIndication, int, error) {
	var v MICOModeIndication
	err := ngapcodec.DecodeParameter("MICOModeIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return MICOModeIndication{}, 0, err
	}
	return v, len(data), err
}
