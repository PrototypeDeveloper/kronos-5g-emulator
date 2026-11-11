// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionPairID represents ASN.1 PDUSessionPairID.
type PDUSessionPairID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionPairID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionPairID", v)
}

// DecodePDUSessionPairID decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionPairID(data []byte) (PDUSessionPairID, int, error) {
	var v PDUSessionPairID
	err := ngapcodec.DecodeParameter("PDUSessionPairID", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionPairID{}, 0, err
	}
	return v, len(data), err
}
