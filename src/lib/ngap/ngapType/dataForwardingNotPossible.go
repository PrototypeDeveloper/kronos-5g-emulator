// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingNotPossible represents ASN.1 DataForwardingNotPossible.
const (
	DataForwardingNotPossiblePresentDataForwardingNotPossible Enumerated = 0
)

type DataForwardingNotPossible struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingNotPossible) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingNotPossible", v)
}

// DecodeDataForwardingNotPossible decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingNotPossible(data []byte) (DataForwardingNotPossible, int, error) {
	var v DataForwardingNotPossible
	err := ngapcodec.DecodeParameter("DataForwardingNotPossible", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingNotPossible{}, 0, err
	}
	return v, len(data), err
}
