// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPagingIdentity represents ASN.1 UEPagingIdentity.
type UEPagingIdentityPresent uint32

const (
	UEPagingIdentityPresentNothing UEPagingIdentityPresent = iota
	UEPagingIdentityPresentFiveGSTMSI
	UEPagingIdentityPresentChoiceExtensions
)

type UEPagingIdentity struct {
	Present          UEPagingIdentityPresent
	FiveGSTMSI       FiveGSTMSI
	ChoiceExtensions UEPagingIdentityExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPagingIdentity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEPagingIdentity", v)
}

// DecodeUEPagingIdentity decodes a standalone APER value and reports the consumed octets.
func DecodeUEPagingIdentity(data []byte) (UEPagingIdentity, int, error) {
	var v UEPagingIdentity
	err := ngapcodec.DecodeParameter("UEPagingIdentity", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPagingIdentity{}, 0, err
	}
	return v, len(data), err
}
