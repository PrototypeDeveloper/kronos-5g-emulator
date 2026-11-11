// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedCAGListPerPLMN represents ASN.1 Allowed-CAG-List-per-PLMN.
type AllowedCAGListPerPLMN struct {
	List []CAGID `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedCAGListPerPLMN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Allowed-CAG-List-per-PLMN", v)
}

// DecodeAllowedCAGListPerPLMN decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedCAGListPerPLMN(data []byte) (AllowedCAGListPerPLMN, int, error) {
	var v AllowedCAGListPerPLMN
	err := ngapcodec.DecodeParameter("Allowed-CAG-List-per-PLMN", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedCAGListPerPLMN{}, 0, err
	}
	return v, len(data), err
}
