// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LoggingDuration represents ASN.1 LoggingDuration.
const (
	LoggingDurationPresentM10  Enumerated = 0
	LoggingDurationPresentM20  Enumerated = 1
	LoggingDurationPresentM40  Enumerated = 2
	LoggingDurationPresentM60  Enumerated = 3
	LoggingDurationPresentM90  Enumerated = 4
	LoggingDurationPresentM120 Enumerated = 5
)

type LoggingDuration struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggingDuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LoggingDuration", v)
}

// DecodeLoggingDuration decodes a standalone APER value and reports the consumed octets.
func DecodeLoggingDuration(data []byte) (LoggingDuration, int, error) {
	var v LoggingDuration
	err := ngapcodec.DecodeParameter("LoggingDuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggingDuration{}, 0, err
	}
	return v, len(data), err
}
