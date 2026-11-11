// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserLocationInformationWAGF represents ASN.1 UserLocationInformationW-AGF.
type UserLocationInformationWAGFPresent uint32

const (
	UserLocationInformationWAGFPresentNothing UserLocationInformationWAGFPresent = iota
	UserLocationInformationWAGFPresentGlobalLineID
	UserLocationInformationWAGFPresentHFCNodeID
	UserLocationInformationWAGFPresentChoiceExtensions
)

type UserLocationInformationWAGF struct {
	Present          UserLocationInformationWAGFPresent
	GlobalLineID     GlobalLineID
	HFCNodeID        HFCNodeID
	ChoiceExtensions UserLocationInformationWAGFExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationWAGF) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationW-AGF", v)
}

// DecodeUserLocationInformationWAGF decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationWAGF(data []byte) (UserLocationInformationWAGF, int, error) {
	var v UserLocationInformationWAGF
	err := ngapcodec.DecodeParameter("UserLocationInformationW-AGF", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationWAGF{}, 0, err
	}
	return v, len(data), err
}
