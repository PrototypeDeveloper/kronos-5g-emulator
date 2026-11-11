// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NgENBID represents ASN.1 NgENB-ID.
type NgENBIDPresent uint32

const (
	NgENBIDPresentNothing NgENBIDPresent = iota
	NgENBIDPresentMacroNgENBID
	NgENBIDPresentShortMacroNgENBID
	NgENBIDPresentLongMacroNgENBID
	NgENBIDPresentChoiceExtensions
)

type NgENBID struct {
	Present           NgENBIDPresent
	MacroNgENBID      BitString `aper:"sizeLB:20,sizeUB:20"`
	ShortMacroNgENBID BitString `aper:"sizeLB:18,sizeUB:18"`
	LongMacroNgENBID  BitString `aper:"sizeLB:21,sizeUB:21"`
	ChoiceExtensions  NgENBIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NgENBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NgENB-ID", v) }

// DecodeNgENBID decodes a standalone APER value and reports the consumed octets.
func DecodeNgENBID(data []byte) (NgENBID, int, error) {
	var v NgENBID
	err := ngapcodec.DecodeParameter("NgENB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return NgENBID{}, 0, err
	}
	return v, len(data), err
}
