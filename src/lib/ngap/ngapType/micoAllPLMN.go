// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MicoAllPLMN represents ASN.1 MicoAllPLMN.
const (
	MicoAllPLMNPresentTrue Enumerated = 0
)

type MicoAllPLMN struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MicoAllPLMN) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MicoAllPLMN", v) }

// DecodeMicoAllPLMN decodes a standalone APER value and reports the consumed octets.
func DecodeMicoAllPLMN(data []byte) (MicoAllPLMN, int, error) {
	var v MicoAllPLMN
	err := ngapcodec.DecodeParameter("MicoAllPLMN", data, &v)
	if ngapcodec.IsFatal(err) {
		return MicoAllPLMN{}, 0, err
	}
	return v, len(data), err
}
