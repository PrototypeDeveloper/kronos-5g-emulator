// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AvailableRANVisibleQoEMetrics represents ASN.1 AvailableRANVisibleQoEMetrics.
type AvailableRANVisibleQoEMetrics struct {
	HasApplicationLayerBufferLevelList bool
	ApplicationLayerBufferLevelList    AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList `aper:"valueExt,valueLB:0,valueUB:0"`
	HasPlayoutDelayForMediaStartup     bool
	PlayoutDelayForMediaStartup        AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartup `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions                    bool
	IEExtensions                       []AvailableRANVisibleQoEMetricsExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AvailableRANVisibleQoEMetrics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AvailableRANVisibleQoEMetrics", v)
}

// DecodeAvailableRANVisibleQoEMetrics decodes a standalone APER value and reports the consumed octets.
func DecodeAvailableRANVisibleQoEMetrics(data []byte) (AvailableRANVisibleQoEMetrics, int, error) {
	var v AvailableRANVisibleQoEMetrics
	err := ngapcodec.DecodeParameter("AvailableRANVisibleQoEMetrics", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableRANVisibleQoEMetrics{}, 0, err
	}
	return v, len(data), err
}
