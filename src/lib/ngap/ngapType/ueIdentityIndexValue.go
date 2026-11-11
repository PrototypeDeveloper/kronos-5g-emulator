// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEIdentityIndexValue represents ASN.1 UEIdentityIndexValue.
type UEIdentityIndexValuePresent uint32

const (
	UEIdentityIndexValuePresentNothing UEIdentityIndexValuePresent = iota
	UEIdentityIndexValuePresentIndexLength10
	UEIdentityIndexValuePresentChoiceExtensions
)

type UEIdentityIndexValue struct {
	Present          UEIdentityIndexValuePresent
	IndexLength10    BitString `aper:"sizeLB:10,sizeUB:10"`
	ChoiceExtensions UEIdentityIndexValueExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEIdentityIndexValue) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEIdentityIndexValue", v)
}

// DecodeUEIdentityIndexValue decodes a standalone APER value and reports the consumed octets.
func DecodeUEIdentityIndexValue(data []byte) (UEIdentityIndexValue, int, error) {
	var v UEIdentityIndexValue
	err := ngapcodec.DecodeParameter("UEIdentityIndexValue", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEIdentityIndexValue{}, 0, err
	}
	return v, len(data), err
}
