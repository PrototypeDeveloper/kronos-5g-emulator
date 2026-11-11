// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPresence represents ASN.1 UEPresence.
const (
	UEPresencePresentIn      Enumerated = 0
	UEPresencePresentOut     Enumerated = 1
	UEPresencePresentUnknown Enumerated = 2
)

type UEPresence struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPresence) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("UEPresence", v) }

// DecodeUEPresence decodes a standalone APER value and reports the consumed octets.
func DecodeUEPresence(data []byte) (UEPresence, int, error) {
	var v UEPresence
	err := ngapcodec.DecodeParameter("UEPresence", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPresence{}, 0, err
	}
	return v, len(data), err
}
