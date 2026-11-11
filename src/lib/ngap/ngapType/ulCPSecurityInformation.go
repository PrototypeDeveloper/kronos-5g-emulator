// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ULCPSecurityInformation represents ASN.1 UL-CP-SecurityInformation.
type ULCPSecurityInformation struct {
	UlNASMAC          ULNASMAC
	UlNASCount        ULNASCount
	HasIEExtensions   bool
	IEExtensions      []ULCPSecurityInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULCPSecurityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UL-CP-SecurityInformation", v)
}

// DecodeULCPSecurityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeULCPSecurityInformation(data []byte) (ULCPSecurityInformation, int, error) {
	var v ULCPSecurityInformation
	err := ngapcodec.DecodeParameter("UL-CP-SecurityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULCPSecurityInformation{}, 0, err
	}
	return v, len(data), err
}
