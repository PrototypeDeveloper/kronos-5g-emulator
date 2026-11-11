// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeStamp represents ASN.1 TimeStamp.
type TimeStamp struct {
	Value []byte `aper:"sizeLB:4,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeStamp) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TimeStamp", v) }

// DecodeTimeStamp decodes a standalone APER value and reports the consumed octets.
func DecodeTimeStamp(data []byte) (TimeStamp, int, error) {
	var v TimeStamp
	err := ngapcodec.DecodeParameter("TimeStamp", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeStamp{}, 0, err
	}
	return v, len(data), err
}
