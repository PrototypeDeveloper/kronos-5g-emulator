// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoEReference represents ASN.1 QoEReference.
type QoEReference struct {
	Value []byte `aper:"sizeLB:6,sizeUB:6"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoEReference) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("QoEReference", v) }

// DecodeQoEReference decodes a standalone APER value and reports the consumed octets.
func DecodeQoEReference(data []byte) (QoEReference, int, error) {
	var v QoEReference
	err := ngapcodec.DecodeParameter("QoEReference", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoEReference{}, 0, err
	}
	return v, len(data), err
}
