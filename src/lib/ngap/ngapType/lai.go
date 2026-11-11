// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LAI represents ASN.1 LAI.
type LAI struct {
	PLMNidentity      PLMNIdentity
	LAC               LAC
	HasIEExtensions   bool
	IEExtensions      []LAIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("LAI", v) }

// DecodeLAI decodes a standalone APER value and reports the consumed octets.
func DecodeLAI(data []byte) (LAI, int, error) {
	var v LAI
	err := ngapcodec.DecodeParameter("LAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return LAI{}, 0, err
	}
	return v, len(data), err
}
