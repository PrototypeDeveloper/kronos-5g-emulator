// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalRANNodeID represents ASN.1 GlobalRANNodeID.
type GlobalRANNodeIDPresent uint32

const (
	GlobalRANNodeIDPresentNothing GlobalRANNodeIDPresent = iota
	GlobalRANNodeIDPresentGlobalGNBID
	GlobalRANNodeIDPresentGlobalNgENBID
	GlobalRANNodeIDPresentGlobalN3IWFID
	GlobalRANNodeIDPresentChoiceExtensions
)

type GlobalRANNodeID struct {
	Present          GlobalRANNodeIDPresent
	GlobalGNBID      GlobalGNBID
	GlobalNgENBID    GlobalNgENBID
	GlobalN3IWFID    GlobalN3IWFID
	ChoiceExtensions GlobalRANNodeIDExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalRANNodeID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalRANNodeID", v)
}

// DecodeGlobalRANNodeID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalRANNodeID(data []byte) (GlobalRANNodeID, int, error) {
	var v GlobalRANNodeID
	err := ngapcodec.DecodeParameter("GlobalRANNodeID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalRANNodeID{}, 0, err
	}
	return v, len(data), err
}
