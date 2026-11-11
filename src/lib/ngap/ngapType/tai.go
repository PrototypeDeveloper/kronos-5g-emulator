// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAI represents ASN.1 TAI.
type TAI struct {
	PLMNIdentity      PLMNIdentity
	TAC               TAC
	HasIEExtensions   bool
	IEExtensions      []TAIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAI", v) }

// DecodeTAI decodes a standalone APER value and reports the consumed octets.
func DecodeTAI(data []byte) (TAI, int, error) {
	var v TAI
	err := ngapcodec.DecodeParameter("TAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAI{}, 0, err
	}
	return v, len(data), err
}
