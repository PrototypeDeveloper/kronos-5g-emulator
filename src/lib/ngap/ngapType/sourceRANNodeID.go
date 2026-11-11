// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceRANNodeID represents ASN.1 SourceRANNodeID.
type SourceRANNodeID struct {
	GlobalRANNodeID   GlobalRANNodeID
	SelectedTAI       TAI
	HasIEExtensions   bool
	IEExtensions      []SourceRANNodeIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceRANNodeID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceRANNodeID", v)
}

// DecodeSourceRANNodeID decodes a standalone APER value and reports the consumed octets.
func DecodeSourceRANNodeID(data []byte) (SourceRANNodeID, int, error) {
	var v SourceRANNodeID
	err := ngapcodec.DecodeParameter("SourceRANNodeID", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceRANNodeID{}, 0, err
	}
	return v, len(data), err
}
