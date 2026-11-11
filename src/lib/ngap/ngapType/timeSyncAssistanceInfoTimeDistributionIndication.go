// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeSyncAssistanceInfoTimeDistributionIndication represents ASN.1 TimeSyncAssistanceInfoTimeDistributionIndication.
const (
	TimeSyncAssistanceInfoTimeDistributionIndicationPresentEnabled  Enumerated = 0
	TimeSyncAssistanceInfoTimeDistributionIndicationPresentDisabled Enumerated = 1
)

type TimeSyncAssistanceInfoTimeDistributionIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v TimeSyncAssistanceInfoTimeDistributionIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:TimeSyncAssistanceInfoTimeDistributionIndication", v)
}

// DecodeTimeSyncAssistanceInfoTimeDistributionIndication decodes a standalone APER value and reports consumed octets.
func DecodeTimeSyncAssistanceInfoTimeDistributionIndication(data []byte) (TimeSyncAssistanceInfoTimeDistributionIndication, int, error) {
	var v TimeSyncAssistanceInfoTimeDistributionIndication
	err := ngapcodec.DecodeParameter("inline:TimeSyncAssistanceInfoTimeDistributionIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeSyncAssistanceInfoTimeDistributionIndication{}, 0, err
	}
	return v, len(data), err
}
