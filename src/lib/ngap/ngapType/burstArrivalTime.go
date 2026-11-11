// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BurstArrivalTime represents ASN.1 BurstArrivalTime.
type BurstArrivalTime struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BurstArrivalTime) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BurstArrivalTime", v)
}

// DecodeBurstArrivalTime decodes a standalone APER value and reports the consumed octets.
func DecodeBurstArrivalTime(data []byte) (BurstArrivalTime, int, error) {
	var v BurstArrivalTime
	err := ngapcodec.DecodeParameter("BurstArrivalTime", data, &v)
	if ngapcodec.IsFatal(err) {
		return BurstArrivalTime{}, 0, err
	}
	return v, len(data), err
}
