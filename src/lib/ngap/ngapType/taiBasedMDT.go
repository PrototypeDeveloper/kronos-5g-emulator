// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBasedMDT represents ASN.1 TAIBasedMDT.
type TAIBasedMDT struct {
	TAIListforMDT     TAIListforMDT
	HasIEExtensions   bool
	IEExtensions      []TAIBasedMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBasedMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TAIBasedMDT", v) }

// DecodeTAIBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBasedMDT(data []byte) (TAIBasedMDT, int, error) {
	var v TAIBasedMDT
	err := ngapcodec.DecodeParameter("TAIBasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBasedMDT{}, 0, err
	}
	return v, len(data), err
}
