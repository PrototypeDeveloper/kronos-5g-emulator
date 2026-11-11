// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANPagingPriority represents ASN.1 RANPagingPriority.
type RANPagingPriority struct {
	Value int64 `aper:"valueLB:1,valueUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANPagingPriority) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANPagingPriority", v)
}

// DecodeRANPagingPriority decodes a standalone APER value and reports the consumed octets.
func DecodeRANPagingPriority(data []byte) (RANPagingPriority, int, error) {
	var v RANPagingPriority
	err := ngapcodec.DecodeParameter("RANPagingPriority", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANPagingPriority{}, 0, err
	}
	return v, len(data), err
}
