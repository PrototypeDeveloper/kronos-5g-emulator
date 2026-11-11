// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LocationReportingAdditionalInfo represents ASN.1 LocationReportingAdditionalInfo.
const (
	LocationReportingAdditionalInfoPresentIncludePSCell Enumerated = 0
)

type LocationReportingAdditionalInfo struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LocationReportingAdditionalInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LocationReportingAdditionalInfo", v)
}

// DecodeLocationReportingAdditionalInfo decodes a standalone APER value and reports the consumed octets.
func DecodeLocationReportingAdditionalInfo(data []byte) (LocationReportingAdditionalInfo, int, error) {
	var v LocationReportingAdditionalInfo
	err := ngapcodec.DecodeParameter("LocationReportingAdditionalInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return LocationReportingAdditionalInfo{}, 0, err
	}
	return v, len(data), err
}
