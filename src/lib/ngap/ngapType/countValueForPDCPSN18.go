// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// COUNTValueForPDCPSN18 represents ASN.1 COUNTValueForPDCP-SN18.
type COUNTValueForPDCPSN18 struct {
	PDCPSN18          int64 `aper:"valueLB:0,valueUB:262143"`
	HFNPDCPSN18       int64 `aper:"valueLB:0,valueUB:16383"`
	HasIEExtensions   bool
	IEExtensions      []COUNTValueForPDCPSN18ExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v COUNTValueForPDCPSN18) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("COUNTValueForPDCP-SN18", v)
}

// DecodeCOUNTValueForPDCPSN18 decodes a standalone APER value and reports the consumed octets.
func DecodeCOUNTValueForPDCPSN18(data []byte) (COUNTValueForPDCPSN18, int, error) {
	var v COUNTValueForPDCPSN18
	err := ngapcodec.DecodeParameter("COUNTValueForPDCP-SN18", data, &v)
	if ngapcodec.IsFatal(err) {
		return COUNTValueForPDCPSN18{}, 0, err
	}
	return v, len(data), err
}
