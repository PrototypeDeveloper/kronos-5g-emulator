// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RSN represents ASN.1 RSN.
const (
	RSNPresentV1 Enumerated = 0
	RSNPresentV2 Enumerated = 1
)

type RSN struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSN) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RSN", v) }

// DecodeRSN decodes a standalone APER value and reports the consumed octets.
func DecodeRSN(data []byte) (RSN, int, error) {
	var v RSN
	err := ngapcodec.DecodeParameter("RSN", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSN{}, 0, err
	}
	return v, len(data), err
}
