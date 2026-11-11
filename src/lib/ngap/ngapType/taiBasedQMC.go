// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBasedQMC represents ASN.1 TAIBasedQMC.
type TAIBasedQMC struct {
	TAIListforQMC     TAIListforQMC
	HasIEExtensions   bool
	IEExtensions      []TAIBasedQMCExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBasedQMC) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAIBasedQMC", v) }

// DecodeTAIBasedQMC decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBasedQMC(data []byte) (TAIBasedQMC, int, error) {
	var v TAIBasedQMC
	err := ngapcodec.DecodeParameter("TAIBasedQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBasedQMC{}, 0, err
	}
	return v, len(data), err
}
