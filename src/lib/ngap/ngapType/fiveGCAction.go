// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGCAction represents ASN.1 FiveGCAction.
type FiveGCActionPresent uint32

const (
	FiveGCActionPresentNothing FiveGCActionPresent = iota
	FiveGCActionPresentHLComActivate
	FiveGCActionPresentHLComDeactivate
	FiveGCActionPresentChoiceExtensions
)

type FiveGCAction struct {
	Present          FiveGCActionPresent
	HLComActivate    HLComActivate
	HLComDeactivate  HLComDeactivate
	ChoiceExtensions FiveGCActionExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGCAction) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("FiveGCAction", v) }

// DecodeFiveGCAction decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGCAction(data []byte) (FiveGCAction, int, error) {
	var v FiveGCAction
	err := ngapcodec.DecodeParameter("FiveGCAction", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGCAction{}, 0, err
	}
	return v, len(data), err
}
