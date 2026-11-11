// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRCGI represents ASN.1 NR-CGI.
type NRCGI struct {
	PLMNIdentity      PLMNIdentity
	NRCellIdentity    NRCellIdentity
	HasIEExtensions   bool
	IEExtensions      []NRCGIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRCGI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NR-CGI", v) }

// DecodeNRCGI decodes a standalone APER value and reports the consumed octets.
func DecodeNRCGI(data []byte) (NRCGI, int, error) {
	var v NRCGI
	err := ngapcodec.DecodeParameter("NR-CGI", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRCGI{}, 0, err
	}
	return v, len(data), err
}
