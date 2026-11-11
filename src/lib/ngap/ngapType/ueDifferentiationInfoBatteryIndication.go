// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEDifferentiationInfoBatteryIndication represents ASN.1 UEDifferentiationInfoBatteryIndication.
const (
	UEDifferentiationInfoBatteryIndicationPresentBatteryPowered                             Enumerated = 0
	UEDifferentiationInfoBatteryIndicationPresentBatteryPoweredNotRechargeableOrReplaceable Enumerated = 1
	UEDifferentiationInfoBatteryIndicationPresentNotBatteryPowered                          Enumerated = 2
)

type UEDifferentiationInfoBatteryIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value.
func (v UEDifferentiationInfoBatteryIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:UEDifferentiationInfoBatteryIndication", v)
}

// DecodeUEDifferentiationInfoBatteryIndication decodes a standalone APER value and reports consumed octets.
func DecodeUEDifferentiationInfoBatteryIndication(data []byte) (UEDifferentiationInfoBatteryIndication, int, error) {
	var v UEDifferentiationInfoBatteryIndication
	err := ngapcodec.DecodeParameter("inline:UEDifferentiationInfoBatteryIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfoBatteryIndication{}, 0, err
	}
	return v, len(data), err
}
