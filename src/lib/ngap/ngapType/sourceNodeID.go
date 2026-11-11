// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceNodeID represents ASN.1 SourceNodeID.
type SourceNodeIDPresent uint32

const (
	SourceNodeIDPresentNothing SourceNodeIDPresent = iota
	SourceNodeIDPresentSourceengNBID
	SourceNodeIDPresentChoiceExtensions
)

type SourceNodeID struct {
	Present          SourceNodeIDPresent
	SourceengNBID    GlobalGNBID
	ChoiceExtensions SourceNodeIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceNodeID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SourceNodeID", v) }

// DecodeSourceNodeID decodes a standalone APER value and reports the consumed octets.
func DecodeSourceNodeID(data []byte) (SourceNodeID, int, error) {
	var v SourceNodeID
	err := ngapcodec.DecodeParameter("SourceNodeID", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceNodeID{}, 0, err
	}
	return v, len(data), err
}
