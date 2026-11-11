// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AvailableRANVisibleQoEMetricsExtIEsEntryPresent uint32

const (
	AvailableRANVisibleQoEMetricsExtIEsEntryPresentNothing AvailableRANVisibleQoEMetricsExtIEsEntryPresent = iota
	AvailableRANVisibleQoEMetricsExtIEsEntryPresentUnknown
)

type AvailableRANVisibleQoEMetricsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AvailableRANVisibleQoEMetricsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AvailableRANVisibleQoEMetricsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AvailableRANVisibleQoEMetrics-ExtIEsEntry", v)
}

// DecodeAvailableRANVisibleQoEMetricsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAvailableRANVisibleQoEMetricsExtIEsEntry(data []byte) (AvailableRANVisibleQoEMetricsExtIEsEntry, int, error) {
	var v AvailableRANVisibleQoEMetricsExtIEsEntry
	err := ngapcodec.DecodeParameter("AvailableRANVisibleQoEMetrics-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableRANVisibleQoEMetricsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
