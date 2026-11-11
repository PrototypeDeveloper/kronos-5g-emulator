// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WarningType represents ASN.1 WarningType.
type WarningType struct {
	Value []byte `aper:"sizeLB:2,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("WarningType", v) }

// DecodeWarningType decodes a standalone APER value and reports the consumed octets.
func DecodeWarningType(data []byte) (WarningType, int, error) {
	var v WarningType
	err := ngapcodec.DecodeParameter("WarningType", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningType{}, 0, err
	}
	return v, len(data), err
}
