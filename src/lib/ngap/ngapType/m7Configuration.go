// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M7Configuration represents ASN.1 M7Configuration.
type M7Configuration struct {
	M7period          M7period
	M7LinksToLog      LinksToLog
	HasIEExtensions   bool
	IEExtensions      []M7ConfigurationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M7Configuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M7Configuration", v)
}

// DecodeM7Configuration decodes a standalone APER value and reports the consumed octets.
func DecodeM7Configuration(data []byte) (M7Configuration, int, error) {
	var v M7Configuration
	err := ngapcodec.DecodeParameter("M7Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return M7Configuration{}, 0, err
	}
	return v, len(data), err
}
