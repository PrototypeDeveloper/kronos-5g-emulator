// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M5Configuration represents ASN.1 M5Configuration.
type M5Configuration struct {
	M5period          M5period
	M5LinksToLog      LinksToLog
	HasIEExtensions   bool
	IEExtensions      []M5ConfigurationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M5Configuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M5Configuration", v)
}

// DecodeM5Configuration decodes a standalone APER value and reports the consumed octets.
func DecodeM5Configuration(data []byte) (M5Configuration, int, error) {
	var v M5Configuration
	err := ngapcodec.DecodeParameter("M5Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return M5Configuration{}, 0, err
	}
	return v, len(data), err
}
