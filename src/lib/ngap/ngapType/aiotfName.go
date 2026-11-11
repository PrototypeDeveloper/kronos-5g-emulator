// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTFName represents ASN.1 AIOTFName.
type AIOTFName struct {
	HasAIOTFNameVisibleString bool
	AIOTFNameVisibleString    AIOTFNameVisibleString
	HasAIOTFNameUTF8String    bool
	AIOTFNameUTF8String       AIOTFNameUTF8String
	HasIEExtensions           bool
	IEExtensions              []AIOTFNameExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTFName) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AIOTFName", v) }

// DecodeAIOTFName decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTFName(data []byte) (AIOTFName, int, error) {
	var v AIOTFName
	err := ngapcodec.DecodeParameter("AIOTFName", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTFName{}, 0, err
	}
	return v, len(data), err
}
