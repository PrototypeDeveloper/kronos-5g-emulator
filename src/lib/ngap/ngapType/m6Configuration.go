// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M6Configuration represents ASN.1 M6Configuration.
type M6Configuration struct {
	M6reportInterval  M6reportInterval
	M6LinksToLog      LinksToLog
	HasIEExtensions   bool
	IEExtensions      []M6ConfigurationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M6Configuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M6Configuration", v)
}

// DecodeM6Configuration decodes a standalone APER value and reports the consumed octets.
func DecodeM6Configuration(data []byte) (M6Configuration, int, error) {
	var v M6Configuration
	err := ngapcodec.DecodeParameter("M6Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return M6Configuration{}, 0, err
	}
	return v, len(data), err
}
