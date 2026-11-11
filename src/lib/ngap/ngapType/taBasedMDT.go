// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TABasedMDT represents ASN.1 TABasedMDT.
type TABasedMDT struct {
	TAListforMDT      TAListforMDT
	HasIEExtensions   bool
	IEExtensions      []TABasedMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TABasedMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TABasedMDT", v) }

// DecodeTABasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeTABasedMDT(data []byte) (TABasedMDT, int, error) {
	var v TABasedMDT
	err := ngapcodec.DecodeParameter("TABasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return TABasedMDT{}, 0, err
	}
	return v, len(data), err
}
