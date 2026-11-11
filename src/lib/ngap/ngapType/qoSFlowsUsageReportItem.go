// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowsUsageReportItem represents ASN.1 QoSFlowsUsageReport-Item.
type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       QosFlowIdentifier
	RATType                 QoSFlowsUsageReportItemRATType `aper:"valueExt,valueLB:0,valueUB:1"`
	QoSFlowsTimedReportList VolumeTimedReportList
	HasIEExtensions         bool
	IEExtensions            []QoSFlowsUsageReportItemExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowsUsageReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowsUsageReport-Item", v)
}

// DecodeQoSFlowsUsageReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowsUsageReportItem(data []byte) (QoSFlowsUsageReportItem, int, error) {
	var v QoSFlowsUsageReportItem
	err := ngapcodec.DecodeParameter("QoSFlowsUsageReport-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowsUsageReportItem{}, 0, err
	}
	return v, len(data), err
}
