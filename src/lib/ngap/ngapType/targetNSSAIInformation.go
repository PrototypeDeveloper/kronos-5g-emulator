// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TargetNSSAIInformation represents ASN.1 TargetNSSAIInformation.
type TargetNSSAIInformation struct {
	TargetNSSAI       TargetNSSAI
	IndexToRFSP       IndexToRFSP
	HasIEExtensions   bool
	IEExtensions      []TargetNSSAIInformationItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetNSSAIInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetNSSAIInformation", v)
}

// DecodeTargetNSSAIInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTargetNSSAIInformation(data []byte) (TargetNSSAIInformation, int, error) {
	var v TargetNSSAIInformation
	err := ngapcodec.DecodeParameter("TargetNSSAIInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetNSSAIInformation{}, 0, err
	}
	return v, len(data), err
}
