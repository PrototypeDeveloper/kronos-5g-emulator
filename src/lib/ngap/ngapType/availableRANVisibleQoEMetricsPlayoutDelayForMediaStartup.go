// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup represents ASN.1 AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup.
const (
	AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartupPresentTrue Enumerated = 0
)

type AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup", v)
}

// DecodeAvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup decodes a standalone APER value and reports consumed octets.
func DecodeAvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup(data []byte) (AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup, int, error) {
	var v AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup
	err := ngapcodec.DecodeParameter("inline:AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup{}, 0, err
	}
	return v, len(data), err
}
