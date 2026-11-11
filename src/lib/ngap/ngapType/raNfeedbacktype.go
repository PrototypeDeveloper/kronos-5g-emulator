// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANfeedbacktype represents ASN.1 RANfeedbacktype.
type RANfeedbacktypePresent uint32

const (
	RANfeedbacktypePresentNothing RANfeedbacktypePresent = iota
	RANfeedbacktypePresentProactive
	RANfeedbacktypePresentReactive
	RANfeedbacktypePresentChoiceExtensions
)

type RANfeedbacktype struct {
	Present          RANfeedbacktypePresent
	Proactive        RANfeedbacktypeProactive
	Reactive         RANfeedbacktypeReactive
	ChoiceExtensions RANfeedbacktypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktype) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype", v)
}

// DecodeRANfeedbacktype decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktype(data []byte) (RANfeedbacktype, int, error) {
	var v RANfeedbacktype
	err := ngapcodec.DecodeParameter("RANfeedbacktype", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktype{}, 0, err
	}
	return v, len(data), err
}
