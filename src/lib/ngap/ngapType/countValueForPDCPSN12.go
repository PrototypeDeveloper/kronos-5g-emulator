// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// COUNTValueForPDCPSN12 represents ASN.1 COUNTValueForPDCP-SN12.
type COUNTValueForPDCPSN12 struct {
	PDCPSN12          int64 `aper:"valueLB:0,valueUB:4095"`
	HFNPDCPSN12       int64 `aper:"valueLB:0,valueUB:1048575"`
	HasIEExtensions   bool
	IEExtensions      []COUNTValueForPDCPSN12ExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v COUNTValueForPDCPSN12) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("COUNTValueForPDCP-SN12", v)
}

// DecodeCOUNTValueForPDCPSN12 decodes a standalone APER value and reports the consumed octets.
func DecodeCOUNTValueForPDCPSN12(data []byte) (COUNTValueForPDCPSN12, int, error) {
	var v COUNTValueForPDCPSN12
	err := ngapcodec.DecodeParameter("COUNTValueForPDCP-SN12", data, &v)
	if ngapcodec.IsFatal(err) {
		return COUNTValueForPDCPSN12{}, 0, err
	}
	return v, len(data), err
}
