// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ECNMarkingorCongestionInformationReportingStatus represents ASN.1 ECNMarkingorCongestionInformationReportingStatus.
type ECNMarkingorCongestionInformationReportingStatus struct {
	List []ECNMarkingorCongestionInformationReportingStatusItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingorCongestionInformationReportingStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingorCongestionInformationReportingStatus", v)
}

// DecodeECNMarkingorCongestionInformationReportingStatus decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingorCongestionInformationReportingStatus(data []byte) (ECNMarkingorCongestionInformationReportingStatus, int, error) {
	var v ECNMarkingorCongestionInformationReportingStatus
	err := ngapcodec.DecodeParameter("ECNMarkingorCongestionInformationReportingStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingorCongestionInformationReportingStatus{}, 0, err
	}
	return v, len(data), err
}
