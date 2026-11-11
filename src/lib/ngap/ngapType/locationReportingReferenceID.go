// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LocationReportingReferenceID represents ASN.1 LocationReportingReferenceID.
type LocationReportingReferenceID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LocationReportingReferenceID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LocationReportingReferenceID", v)
}

// DecodeLocationReportingReferenceID decodes a standalone APER value and reports the consumed octets.
func DecodeLocationReportingReferenceID(data []byte) (LocationReportingReferenceID, int, error) {
	var v LocationReportingReferenceID
	err := ngapcodec.DecodeParameter("LocationReportingReferenceID", data, &v)
	if ngapcodec.IsFatal(err) {
		return LocationReportingReferenceID{}, 0, err
	}
	return v, len(data), err
}
