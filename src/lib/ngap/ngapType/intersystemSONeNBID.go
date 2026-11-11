// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONeNBID represents ASN.1 IntersystemSONeNBID.
type IntersystemSONeNBID struct {
	GlobaleNBID       GlobalENBID
	SelectedEPSTAI    EPSTAI
	HasIEExtensions   bool
	IEExtensions      []IntersystemSONeNBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONeNBID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONeNBID", v)
}

// DecodeIntersystemSONeNBID decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONeNBID(data []byte) (IntersystemSONeNBID, int, error) {
	var v IntersystemSONeNBID
	err := ngapcodec.DecodeParameter("IntersystemSONeNBID", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONeNBID{}, 0, err
	}
	return v, len(data), err
}
