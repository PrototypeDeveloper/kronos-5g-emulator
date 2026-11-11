// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetHomeENBID represents ASN.1 TargetHomeENB-ID.
type TargetHomeENBID struct {
	PLMNidentity      PLMNIdentity
	HomeENBID         BitString `aper:"sizeLB:28,sizeUB:28"`
	SelectedEPSTAI    EPSTAI
	HasIEExtensions   bool
	IEExtensions      []TargetHomeENBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetHomeENBID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetHomeENB-ID", v)
}

// DecodeTargetHomeENBID decodes a standalone APER value and reports the consumed octets.
func DecodeTargetHomeENBID(data []byte) (TargetHomeENBID, int, error) {
	var v TargetHomeENBID
	err := ngapcodec.DecodeParameter("TargetHomeENB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetHomeENBID{}, 0, err
	}
	return v, len(data), err
}
