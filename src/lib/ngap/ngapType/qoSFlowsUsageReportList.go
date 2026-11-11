// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowsUsageReportList represents ASN.1 QoSFlowsUsageReportList.
type QoSFlowsUsageReportList struct {
	List []QoSFlowsUsageReportItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowsUsageReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowsUsageReportList", v)
}

// DecodeQoSFlowsUsageReportList decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowsUsageReportList(data []byte) (QoSFlowsUsageReportList, int, error) {
	var v QoSFlowsUsageReportList
	err := ngapcodec.DecodeParameter("QoSFlowsUsageReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowsUsageReportList{}, 0, err
	}
	return v, len(data), err
}
