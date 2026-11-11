// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionID represents ASN.1 PDUSessionID.
type PDUSessionID struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PDUSessionID", v) }

// DecodePDUSessionID decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionID(data []byte) (PDUSessionID, int, error) {
	var v PDUSessionID
	err := ngapcodec.DecodeParameter("PDUSessionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionID{}, 0, err
	}
	return v, len(data), err
}
