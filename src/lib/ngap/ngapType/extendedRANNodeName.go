// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedRANNodeName represents ASN.1 Extended-RANNodeName.
type ExtendedRANNodeName struct {
	HasRANNodeNameVisibleString bool
	RANNodeNameVisibleString    RANNodeNameVisibleString
	HasRANNodeNameUTF8String    bool
	RANNodeNameUTF8String       RANNodeNameUTF8String
	HasIEExtensions             bool
	IEExtensions                []ExtendedRANNodeNameExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedRANNodeName) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Extended-RANNodeName", v)
}

// DecodeExtendedRANNodeName decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedRANNodeName(data []byte) (ExtendedRANNodeName, int, error) {
	var v ExtendedRANNodeName
	err := ngapcodec.DecodeParameter("Extended-RANNodeName", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedRANNodeName{}, 0, err
	}
	return v, len(data), err
}
