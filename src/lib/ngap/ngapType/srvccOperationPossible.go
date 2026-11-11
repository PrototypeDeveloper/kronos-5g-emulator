// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SRVCCOperationPossible represents ASN.1 SRVCCOperationPossible.
const (
	SRVCCOperationPossiblePresentPossible    Enumerated = 0
	SRVCCOperationPossiblePresentNotPossible Enumerated = 1
)

type SRVCCOperationPossible struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SRVCCOperationPossible) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SRVCCOperationPossible", v)
}

// DecodeSRVCCOperationPossible decodes a standalone APER value and reports the consumed octets.
func DecodeSRVCCOperationPossible(data []byte) (SRVCCOperationPossible, int, error) {
	var v SRVCCOperationPossible
	err := ngapcodec.DecodeParameter("SRVCCOperationPossible", data, &v)
	if ngapcodec.IsFatal(err) {
		return SRVCCOperationPossible{}, 0, err
	}
	return v, len(data), err
}
