// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEDifferentiationInfoPeriodicCommunicationIndicator represents ASN.1 UEDifferentiationInfoPeriodicCommunicationIndicator.
const (
	UEDifferentiationInfoPeriodicCommunicationIndicatorPresentPeriodically Enumerated = 0
	UEDifferentiationInfoPeriodicCommunicationIndicatorPresentOndemand     Enumerated = 1
)

type UEDifferentiationInfoPeriodicCommunicationIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v UEDifferentiationInfoPeriodicCommunicationIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:UEDifferentiationInfoPeriodicCommunicationIndicator", v)
}

// DecodeUEDifferentiationInfoPeriodicCommunicationIndicator decodes a standalone APER value and reports consumed octets.
func DecodeUEDifferentiationInfoPeriodicCommunicationIndicator(data []byte) (UEDifferentiationInfoPeriodicCommunicationIndicator, int, error) {
	var v UEDifferentiationInfoPeriodicCommunicationIndicator
	err := ngapcodec.DecodeParameter("inline:UEDifferentiationInfoPeriodicCommunicationIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfoPeriodicCommunicationIndicator{}, 0, err
	}
	return v, len(data), err
}
