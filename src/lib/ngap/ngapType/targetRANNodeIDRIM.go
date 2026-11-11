// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetRANNodeIDRIM represents ASN.1 TargetRANNodeID-RIM.
type TargetRANNodeIDRIM struct {
	GlobalRANNodeID   GlobalRANNodeID
	SelectedTAI       TAI
	HasIEExtensions   bool
	IEExtensions      []TargetRANNodeIDRIMExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeIDRIM) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID-RIM", v)
}

// DecodeTargetRANNodeIDRIM decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeIDRIM(data []byte) (TargetRANNodeIDRIM, int, error) {
	var v TargetRANNodeIDRIM
	err := ngapcodec.DecodeParameter("TargetRANNodeID-RIM", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeIDRIM{}, 0, err
	}
	return v, len(data), err
}
