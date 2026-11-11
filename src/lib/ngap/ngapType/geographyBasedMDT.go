// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GeographyBasedMDT represents ASN.1 GeographyBasedMDT.
type GeographyBasedMDT struct {
	NTNGeographicalArea NTNGeographicalAreaList
	HasNTNPLMNList      bool
	NTNPLMNList         MDTPLMNList
	HasIEExtensions     bool
	IEExtensions        []GeographyBasedMDTExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GeographyBasedMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GeographyBasedMDT", v)
}

// DecodeGeographyBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeGeographyBasedMDT(data []byte) (GeographyBasedMDT, int, error) {
	var v GeographyBasedMDT
	err := ngapcodec.DecodeParameter("GeographyBasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return GeographyBasedMDT{}, 0, err
	}
	return v, len(data), err
}
