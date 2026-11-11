// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PEIPSassistanceInformation represents ASN.1 PEIPSassistanceInformation.
type PEIPSassistanceInformation struct {
	CNsubgroupID      CNsubgroupID
	HasIEExtensions   bool
	IEExtensions      []PEIPSassistanceInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PEIPSassistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PEIPSassistanceInformation", v)
}

// DecodePEIPSassistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodePEIPSassistanceInformation(data []byte) (PEIPSassistanceInformation, int, error) {
	var v PEIPSassistanceInformation
	err := ngapcodec.DecodeParameter("PEIPSassistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PEIPSassistanceInformation{}, 0, err
	}
	return v, len(data), err
}
