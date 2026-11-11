// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LoggingInterval represents ASN.1 LoggingInterval.
const (
	LoggingIntervalPresentMs320    Enumerated = 0
	LoggingIntervalPresentMs640    Enumerated = 1
	LoggingIntervalPresentMs1280   Enumerated = 2
	LoggingIntervalPresentMs2560   Enumerated = 3
	LoggingIntervalPresentMs5120   Enumerated = 4
	LoggingIntervalPresentMs10240  Enumerated = 5
	LoggingIntervalPresentMs20480  Enumerated = 6
	LoggingIntervalPresentMs30720  Enumerated = 7
	LoggingIntervalPresentMs40960  Enumerated = 8
	LoggingIntervalPresentMs61440  Enumerated = 9
	LoggingIntervalPresentInfinity Enumerated = 10
)

type LoggingInterval struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:10"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggingInterval) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LoggingInterval", v)
}

// DecodeLoggingInterval decodes a standalone APER value and reports the consumed octets.
func DecodeLoggingInterval(data []byte) (LoggingInterval, int, error) {
	var v LoggingInterval
	err := ngapcodec.DecodeParameter("LoggingInterval", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggingInterval{}, 0, err
	}
	return v, len(data), err
}
