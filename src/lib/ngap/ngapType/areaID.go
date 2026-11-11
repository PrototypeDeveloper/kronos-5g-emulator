// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaID represents ASN.1 Area-ID.
type AreaIDPresent uint32

const (
	AreaIDPresentNothing AreaIDPresent = iota
	AreaIDPresentTAI
	AreaIDPresentGlobalRANNodeID
	AreaIDPresentNRCGI
	AreaIDPresentChoiceExtensions
)

type AreaID struct {
	Present          AreaIDPresent
	TAI              TAI
	GlobalRANNodeID  GlobalRANNodeID
	NRCGI            NRCGI
	ChoiceExtensions AreaIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Area-ID", v) }

// DecodeAreaID decodes a standalone APER value and reports the consumed octets.
func DecodeAreaID(data []byte) (AreaID, int, error) {
	var v AreaID
	err := ngapcodec.DecodeParameter("Area-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaID{}, 0, err
	}
	return v, len(data), err
}
