// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList represents ASN.1 AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList.
const (
	AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelListPresentTrue Enumerated = 0
)

type AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList", v)
}

// DecodeAvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList decodes a standalone APER value and reports consumed octets.
func DecodeAvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList(data []byte) (AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList, int, error) {
	var v AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList
	err := ngapcodec.DecodeParameter("inline:AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelList{}, 0, err
	}
	return v, len(data), err
}
