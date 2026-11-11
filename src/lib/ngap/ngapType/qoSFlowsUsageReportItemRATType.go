// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowsUsageReportItemRATType represents ASN.1 QoSFlowsUsageReportItemRATType.
const (
	QoSFlowsUsageReportItemRATTypePresentNr              Enumerated = 0
	QoSFlowsUsageReportItemRATTypePresentEutra           Enumerated = 1
	QoSFlowsUsageReportItemRATTypePresentNrUnlicensed    Enumerated = 2
	QoSFlowsUsageReportItemRATTypePresentEUtraUnlicensed Enumerated = 3
)

type QoSFlowsUsageReportItemRATType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v QoSFlowsUsageReportItemRATType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:QoSFlowsUsageReportItemRATType", v)
}

// DecodeQoSFlowsUsageReportItemRATType decodes a standalone APER value and reports consumed octets.
func DecodeQoSFlowsUsageReportItemRATType(data []byte) (QoSFlowsUsageReportItemRATType, int, error) {
	var v QoSFlowsUsageReportItemRATType
	err := ngapcodec.DecodeParameter("inline:QoSFlowsUsageReportItemRATType", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowsUsageReportItemRATType{}, 0, err
	}
	return v, len(data), err
}
