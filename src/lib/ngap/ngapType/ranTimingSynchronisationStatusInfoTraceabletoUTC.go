// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTimingSynchronisationStatusInfoTraceabletoUTC represents ASN.1 RANTimingSynchronisationStatusInfoTraceabletoUTC.
const (
	RANTimingSynchronisationStatusInfoTraceabletoUTCPresentTrue  Enumerated = 0
	RANTimingSynchronisationStatusInfoTraceabletoUTCPresentFalse Enumerated = 1
)

type RANTimingSynchronisationStatusInfoTraceabletoUTC struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v RANTimingSynchronisationStatusInfoTraceabletoUTC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:RANTimingSynchronisationStatusInfoTraceabletoUTC", v)
}

// DecodeRANTimingSynchronisationStatusInfoTraceabletoUTC decodes a standalone APER value and reports consumed octets.
func DecodeRANTimingSynchronisationStatusInfoTraceabletoUTC(data []byte) (RANTimingSynchronisationStatusInfoTraceabletoUTC, int, error) {
	var v RANTimingSynchronisationStatusInfoTraceabletoUTC
	err := ngapcodec.DecodeParameter("inline:RANTimingSynchronisationStatusInfoTraceabletoUTC", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTimingSynchronisationStatusInfoTraceabletoUTC{}, 0, err
	}
	return v, len(data), err
}
