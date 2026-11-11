// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ENBID represents ASN.1 ENB-ID.
type ENBIDPresent uint32

const (
	ENBIDPresentNothing ENBIDPresent = iota
	ENBIDPresentMacroENBID
	ENBIDPresentHomeENBID
	ENBIDPresentShortMacroENBID
	ENBIDPresentLongMacroENBID
	ENBIDPresentChoiceExtensions
)

type ENBID struct {
	Present          ENBIDPresent
	MacroENBID       BitString `aper:"sizeLB:20,sizeUB:20"`
	HomeENBID        BitString `aper:"sizeLB:28,sizeUB:28"`
	ShortMacroENBID  BitString `aper:"sizeLB:18,sizeUB:18"`
	LongMacroENBID   BitString `aper:"sizeLB:21,sizeUB:21"`
	ChoiceExtensions ENBIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ENBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ENB-ID", v) }

// DecodeENBID decodes a standalone APER value and reports the consumed octets.
func DecodeENBID(data []byte) (ENBID, int, error) {
	var v ENBID
	err := ngapcodec.DecodeParameter("ENB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return ENBID{}, 0, err
	}
	return v, len(data), err
}
