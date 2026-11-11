// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedRATRestrictionInformation represents ASN.1 ExtendedRATRestrictionInformation.
type ExtendedRATRestrictionInformation struct {
	PrimaryRATRestriction   BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
	SecondaryRATRestriction BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
	HasIEExtensions         bool
	IEExtensions            []ExtendedRATRestrictionInformationExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedRATRestrictionInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedRATRestrictionInformation", v)
}

// DecodeExtendedRATRestrictionInformation decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedRATRestrictionInformation(data []byte) (ExtendedRATRestrictionInformation, int, error) {
	var v ExtendedRATRestrictionInformation
	err := ngapcodec.DecodeParameter("ExtendedRATRestrictionInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedRATRestrictionInformation{}, 0, err
	}
	return v, len(data), err
}
