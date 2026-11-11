// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecondaryRATUsageInformation represents ASN.1 SecondaryRATUsageInformation.
type SecondaryRATUsageInformation struct {
	HasPDUSessionUsageReport   bool
	PDUSessionUsageReport      PDUSessionUsageReport
	HasQosFlowsUsageReportList bool
	QosFlowsUsageReportList    QoSFlowsUsageReportList
	HasIEExtension             bool
	IEExtension                []SecondaryRATUsageInformationExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecondaryRATUsageInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecondaryRATUsageInformation", v)
}

// DecodeSecondaryRATUsageInformation decodes a standalone APER value and reports the consumed octets.
func DecodeSecondaryRATUsageInformation(data []byte) (SecondaryRATUsageInformation, int, error) {
	var v SecondaryRATUsageInformation
	err := ngapcodec.DecodeParameter("SecondaryRATUsageInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecondaryRATUsageInformation{}, 0, err
	}
	return v, len(data), err
}
