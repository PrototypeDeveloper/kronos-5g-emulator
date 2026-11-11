// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QoSFlowsUsageReportItemExtIEsEntryPresent uint32

const (
	QoSFlowsUsageReportItemExtIEsEntryPresentNothing QoSFlowsUsageReportItemExtIEsEntryPresent = iota
	QoSFlowsUsageReportItemExtIEsEntryPresentUnknown
)

type QoSFlowsUsageReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QoSFlowsUsageReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowsUsageReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowsUsageReport-Item-ExtIEsEntry", v)
}

// DecodeQoSFlowsUsageReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowsUsageReportItemExtIEsEntry(data []byte) (QoSFlowsUsageReportItemExtIEsEntry, int, error) {
	var v QoSFlowsUsageReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QoSFlowsUsageReport-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowsUsageReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
