// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SourceToTargetAMFInformationReroute represents ASN.1 SourceToTarget-AMFInformationReroute.
type SourceToTargetAMFInformationReroute struct {
	HasConfiguredNSSAI     bool
	ConfiguredNSSAI        ConfiguredNSSAI
	HasRejectedNSSAIinPLMN bool
	RejectedNSSAIinPLMN    RejectedNSSAIinPLMN
	HasRejectedNSSAIinTA   bool
	RejectedNSSAIinTA      RejectedNSSAIinTA
	HasIEExtensions        bool
	IEExtensions           []SourceToTargetAMFInformationRerouteExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceToTargetAMFInformationReroute) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceToTarget-AMFInformationReroute", v)
}

// DecodeSourceToTargetAMFInformationReroute decodes a standalone APER value and reports the consumed octets.
func DecodeSourceToTargetAMFInformationReroute(data []byte) (SourceToTargetAMFInformationReroute, int, error) {
	var v SourceToTargetAMFInformationReroute
	err := ngapcodec.DecodeParameter("SourceToTarget-AMFInformationReroute", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceToTargetAMFInformationReroute{}, 0, err
	}
	return v, len(data), err
}
