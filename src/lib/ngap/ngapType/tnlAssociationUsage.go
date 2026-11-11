// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNLAssociationUsage represents ASN.1 TNLAssociationUsage.
const (
	TNLAssociationUsagePresentUe    Enumerated = 0
	TNLAssociationUsagePresentNonUe Enumerated = 1
	TNLAssociationUsagePresentBoth  Enumerated = 2
)

type TNLAssociationUsage struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNLAssociationUsage) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNLAssociationUsage", v)
}

// DecodeTNLAssociationUsage decodes a standalone APER value and reports the consumed octets.
func DecodeTNLAssociationUsage(data []byte) (TNLAssociationUsage, int, error) {
	var v TNLAssociationUsage
	err := ngapcodec.DecodeParameter("TNLAssociationUsage", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNLAssociationUsage{}, 0, err
	}
	return v, len(data), err
}
