// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EquivalentSNPNsList represents ASN.1 EquivalentSNPNsList.
type EquivalentSNPNsList struct {
	List []EquivalentSNPNsItem `aper:"sizeLB:1,sizeUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EquivalentSNPNsList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EquivalentSNPNsList", v)
}

// DecodeEquivalentSNPNsList decodes a standalone APER value and reports the consumed octets.
func DecodeEquivalentSNPNsList(data []byte) (EquivalentSNPNsList, int, error) {
	var v EquivalentSNPNsList
	err := ngapcodec.DecodeParameter("EquivalentSNPNsList", data, &v)
	if ngapcodec.IsFatal(err) {
		return EquivalentSNPNsList{}, 0, err
	}
	return v, len(data), err
}
