// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommServiceType represents ASN.1 CommServiceType.
const (
	CommServiceTypePresentMbsBroadcast Enumerated = 0
	CommServiceTypePresentMbsMulticast Enumerated = 1
	CommServiceTypePresentUnicast      Enumerated = 2
)

type CommServiceType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommServiceType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommServiceType", v)
}

// DecodeCommServiceType decodes a standalone APER value and reports the consumed octets.
func DecodeCommServiceType(data []byte) (CommServiceType, int, error) {
	var v CommServiceType
	err := ngapcodec.DecodeParameter("CommServiceType", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommServiceType{}, 0, err
	}
	return v, len(data), err
}
