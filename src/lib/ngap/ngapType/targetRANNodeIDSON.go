// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetRANNodeIDSON represents ASN.1 TargetRANNodeID-SON.
type TargetRANNodeIDSON struct {
	GlobalRANNodeID   GlobalRANNodeID
	SelectedTAI       TAI
	HasIEExtensions   bool
	IEExtensions      []TargetRANNodeIDSONExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeIDSON) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID-SON", v)
}

// DecodeTargetRANNodeIDSON decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeIDSON(data []byte) (TargetRANNodeIDSON, int, error) {
	var v TargetRANNodeIDSON
	err := ngapcodec.DecodeParameter("TargetRANNodeID-SON", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeIDSON{}, 0, err
	}
	return v, len(data), err
}
