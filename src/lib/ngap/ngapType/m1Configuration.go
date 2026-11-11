// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M1Configuration represents ASN.1 M1Configuration.
type M1Configuration struct {
	M1reportingTrigger     M1ReportingTrigger
	HasM1thresholdEventA2  bool
	M1thresholdEventA2     M1ThresholdEventA2
	HasM1periodicReporting bool
	M1periodicReporting    M1PeriodicReporting
	HasIEExtensions        bool
	IEExtensions           []M1ConfigurationExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1Configuration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1Configuration", v)
}

// DecodeM1Configuration decodes a standalone APER value and reports the consumed octets.
func DecodeM1Configuration(data []byte) (M1Configuration, int, error) {
	var v M1Configuration
	err := ngapcodec.DecodeParameter("M1Configuration", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1Configuration{}, 0, err
	}
	return v, len(data), err
}
