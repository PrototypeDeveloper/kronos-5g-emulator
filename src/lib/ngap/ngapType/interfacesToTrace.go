// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InterfacesToTrace represents ASN.1 InterfacesToTrace.
type InterfacesToTrace struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterfacesToTrace) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterfacesToTrace", v)
}

// DecodeInterfacesToTrace decodes a standalone APER value and reports the consumed octets.
func DecodeInterfacesToTrace(data []byte) (InterfacesToTrace, int, error) {
	var v InterfacesToTrace
	err := ngapcodec.DecodeParameter("InterfacesToTrace", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterfacesToTrace{}, 0, err
	}
	return v, len(data), err
}
