// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TABasedQMC represents ASN.1 TABasedQMC.
type TABasedQMC struct {
	TAListforQMC      TAListforQMC
	HasIEExtensions   bool
	IEExtensions      []TABasedQMCExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TABasedQMC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TABasedQMC", v) }

// DecodeTABasedQMC decodes a standalone APER value and reports the consumed octets.
func DecodeTABasedQMC(data []byte) (TABasedQMC, int, error) {
	var v TABasedQMC
	err := ngapcodec.DecodeParameter("TABasedQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return TABasedQMC{}, 0, err
	}
	return v, len(data), err
}
