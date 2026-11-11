// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRACGI represents ASN.1 EUTRA-CGI.
type EUTRACGI struct {
	PLMNIdentity      PLMNIdentity
	EUTRACellIdentity EUTRACellIdentity
	HasIEExtensions   bool
	IEExtensions      []EUTRACGIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRACGI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EUTRA-CGI", v) }

// DecodeEUTRACGI decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRACGI(data []byte) (EUTRACGI, int, error) {
	var v EUTRACGI
	err := ngapcodec.DecodeParameter("EUTRA-CGI", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRACGI{}, 0, err
	}
	return v, len(data), err
}
