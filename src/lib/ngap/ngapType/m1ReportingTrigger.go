// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M1ReportingTrigger represents ASN.1 M1ReportingTrigger.
const (
	M1ReportingTriggerPresentPeriodic                 Enumerated = 0
	M1ReportingTriggerPresentA2eventtriggered         Enumerated = 1
	M1ReportingTriggerPresentA2eventtriggeredPeriodic Enumerated = 2
)

type M1ReportingTrigger struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ReportingTrigger) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1ReportingTrigger", v)
}

// DecodeM1ReportingTrigger decodes a standalone APER value and reports the consumed octets.
func DecodeM1ReportingTrigger(data []byte) (M1ReportingTrigger, int, error) {
	var v M1ReportingTrigger
	err := ngapcodec.DecodeParameter("M1ReportingTrigger", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ReportingTrigger{}, 0, err
	}
	return v, len(data), err
}
