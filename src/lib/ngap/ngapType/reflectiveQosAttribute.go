// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ReflectiveQosAttribute represents ASN.1 ReflectiveQosAttribute.
const (
	ReflectiveQosAttributePresentSubjectTo Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ReflectiveQosAttribute) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ReflectiveQosAttribute", v)
}

// DecodeReflectiveQosAttribute decodes a standalone APER value and reports the consumed octets.
func DecodeReflectiveQosAttribute(data []byte) (ReflectiveQosAttribute, int, error) {
	var v ReflectiveQosAttribute
	err := ngapcodec.DecodeParameter("ReflectiveQosAttribute", data, &v)
	if ngapcodec.IsFatal(err) {
		return ReflectiveQosAttribute{}, 0, err
	}
	return v, len(data), err
}
