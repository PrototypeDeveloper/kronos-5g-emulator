// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetID represents ASN.1 TargetID.
type TargetIDPresent uint32

const (
	TargetIDPresentNothing TargetIDPresent = iota
	TargetIDPresentTargetRANNodeID
	TargetIDPresentTargeteNBID
	TargetIDPresentChoiceExtensions
)

type TargetID struct {
	Present          TargetIDPresent
	TargetRANNodeID  TargetRANNodeID
	TargeteNBID      TargeteNBID
	ChoiceExtensions TargetIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TargetID", v) }

// DecodeTargetID decodes a standalone APER value and reports the consumed octets.
func DecodeTargetID(data []byte) (TargetID, int, error) {
	var v TargetID
	err := ngapcodec.DecodeParameter("TargetID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetID{}, 0, err
	}
	return v, len(data), err
}
