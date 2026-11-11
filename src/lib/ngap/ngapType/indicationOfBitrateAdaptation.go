// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IndicationOfBitrateAdaptation represents ASN.1 Indication-of-Bitrate-Adaptation.
const (
	IndicationOfBitrateAdaptationPresentUplink Enumerated = 0
)

type IndicationOfBitrateAdaptation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IndicationOfBitrateAdaptation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Indication-of-Bitrate-Adaptation", v)
}

// DecodeIndicationOfBitrateAdaptation decodes a standalone APER value and reports the consumed octets.
func DecodeIndicationOfBitrateAdaptation(data []byte) (IndicationOfBitrateAdaptation, int, error) {
	var v IndicationOfBitrateAdaptation
	err := ngapcodec.DecodeParameter("Indication-of-Bitrate-Adaptation", data, &v)
	if ngapcodec.IsFatal(err) {
		return IndicationOfBitrateAdaptation{}, 0, err
	}
	return v, len(data), err
}
