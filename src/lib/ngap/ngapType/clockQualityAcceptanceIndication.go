// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityAcceptanceIndication represents ASN.1 ClockQualityAcceptanceIndication.
type ClockQualityAcceptanceIndication struct {
	ClockQualityAcceptanceCriteria ClockQualityAcceptanceCriteria
	HasIEExtensions                bool
	IEExtensions                   []ClockQualityAcceptanceIndicationExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityAcceptanceIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityAcceptanceIndication", v)
}

// DecodeClockQualityAcceptanceIndication decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityAcceptanceIndication(data []byte) (ClockQualityAcceptanceIndication, int, error) {
	var v ClockQualityAcceptanceIndication
	err := ngapcodec.DecodeParameter("ClockQualityAcceptanceIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceIndication{}, 0, err
	}
	return v, len(data), err
}
