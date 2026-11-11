// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetRANNodeID represents ASN.1 TargetRANNodeID.
type TargetRANNodeID struct {
	GlobalRANNodeID   GlobalRANNodeID
	SelectedTAI       TAI
	HasIEExtensions   bool
	IEExtensions      []TargetRANNodeIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID", v)
}

// DecodeTargetRANNodeID decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeID(data []byte) (TargetRANNodeID, int, error) {
	var v TargetRANNodeID
	err := ngapcodec.DecodeParameter("TargetRANNodeID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeID{}, 0, err
	}
	return v, len(data), err
}
