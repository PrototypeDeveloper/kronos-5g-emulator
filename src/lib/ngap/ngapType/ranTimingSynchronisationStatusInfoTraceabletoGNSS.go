// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTimingSynchronisationStatusInfoTraceabletoGNSS represents ASN.1 RANTimingSynchronisationStatusInfoTraceabletoGNSS.
const (
	RANTimingSynchronisationStatusInfoTraceabletoGNSSPresentTrue  Enumerated = 0
	RANTimingSynchronisationStatusInfoTraceabletoGNSSPresentFalse Enumerated = 1
)

type RANTimingSynchronisationStatusInfoTraceabletoGNSS struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v RANTimingSynchronisationStatusInfoTraceabletoGNSS) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:RANTimingSynchronisationStatusInfoTraceabletoGNSS", v)
}

// DecodeRANTimingSynchronisationStatusInfoTraceabletoGNSS decodes a standalone APER value and reports consumed octets.
func DecodeRANTimingSynchronisationStatusInfoTraceabletoGNSS(data []byte) (RANTimingSynchronisationStatusInfoTraceabletoGNSS, int, error) {
	var v RANTimingSynchronisationStatusInfoTraceabletoGNSS
	err := ngapcodec.DecodeParameter("inline:RANTimingSynchronisationStatusInfoTraceabletoGNSS", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTimingSynchronisationStatusInfoTraceabletoGNSS{}, 0, err
	}
	return v, len(data), err
}
