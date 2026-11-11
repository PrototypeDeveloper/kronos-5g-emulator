// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RIMInformation represents ASN.1 RIMInformation.
type RIMInformation struct {
	TargetgNBSetID    GNBSetID
	RIMRSDetection    RIMInformationRIMRSDetection `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions   bool
	IEExtensions      []RIMInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RIMInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RIMInformation", v)
}

// DecodeRIMInformation decodes a standalone APER value and reports the consumed octets.
func DecodeRIMInformation(data []byte) (RIMInformation, int, error) {
	var v RIMInformation
	err := ngapcodec.DecodeParameter("RIMInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return RIMInformation{}, 0, err
	}
	return v, len(data), err
}
