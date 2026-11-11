// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ECNMarkingorCongestionInformationReportingStatusItem represents ASN.1 ECNMarkingorCongestionInformationReportingStatus-Item.
type ECNMarkingorCongestionInformationReportingStatusItem struct {
	QosFlowIdentifier QosFlowIdentifier
	ActivationStatus  ActivationStatus
	HasIEExtension    bool
	IEExtension       []ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingorCongestionInformationReportingStatusItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingorCongestionInformationReportingStatus-Item", v)
}

// DecodeECNMarkingorCongestionInformationReportingStatusItem decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingorCongestionInformationReportingStatusItem(data []byte) (ECNMarkingorCongestionInformationReportingStatusItem, int, error) {
	var v ECNMarkingorCongestionInformationReportingStatusItem
	err := ngapcodec.DecodeParameter("ECNMarkingorCongestionInformationReportingStatus-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingorCongestionInformationReportingStatusItem{}, 0, err
	}
	return v, len(data), err
}
