// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedAMFName represents ASN.1 Extended-AMFName.
type ExtendedAMFName struct {
	HasAMFNameVisibleString bool
	AMFNameVisibleString    AMFNameVisibleString
	HasAMFNameUTF8String    bool
	AMFNameUTF8String       AMFNameUTF8String
	HasIEExtensions         bool
	IEExtensions            []ExtendedAMFNameExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedAMFName) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Extended-AMFName", v)
}

// DecodeExtendedAMFName decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedAMFName(data []byte) (ExtendedAMFName, int, error) {
	var v ExtendedAMFName
	err := ngapcodec.DecodeParameter("Extended-AMFName", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedAMFName{}, 0, err
	}
	return v, len(data), err
}
