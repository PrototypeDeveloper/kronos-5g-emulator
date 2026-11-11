// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTModeNr represents ASN.1 MDTModeNr.
type MDTModeNrPresent uint32

const (
	MDTModeNrPresentNothing MDTModeNrPresent = iota
	MDTModeNrPresentImmediateMDTNr
	MDTModeNrPresentLoggedMDTNr
	MDTModeNrPresentChoiceExtensions
)

type MDTModeNr struct {
	Present          MDTModeNrPresent
	ImmediateMDTNr   ImmediateMDTNr
	LoggedMDTNr      LoggedMDTNr
	ChoiceExtensions MDTModeNrExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTModeNr) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MDTModeNr", v) }

// DecodeMDTModeNr decodes a standalone APER value and reports the consumed octets.
func DecodeMDTModeNr(data []byte) (MDTModeNr, int, error) {
	var v MDTModeNr
	err := ngapcodec.DecodeParameter("MDTModeNr", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTModeNr{}, 0, err
	}
	return v, len(data), err
}
