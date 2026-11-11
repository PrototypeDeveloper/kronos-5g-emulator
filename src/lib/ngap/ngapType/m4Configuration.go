// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M4Configuration represents ASN.1 M4Configuration.
type M4Configuration struct {
	M4period          M4period
	M4LinksToLog      LinksToLog
	HasIEExtensions   bool
	IEExtensions      []M4ConfigurationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M4Configuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M4Configuration", v)
}

// DecodeM4Configuration decodes a standalone APER value and reports the consumed octets.
func DecodeM4Configuration(data []byte) (M4Configuration, int, error) {
	var v M4Configuration
	err := ngapcodec.DecodeParameter("M4Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return M4Configuration{}, 0, err
	}
	return v, len(data), err
}
