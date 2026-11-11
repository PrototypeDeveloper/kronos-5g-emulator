// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLCPSecurityInformation represents ASN.1 DL-CP-SecurityInformation.
type DLCPSecurityInformation struct {
	DlNASMAC          DLNASMAC
	HasIEExtensions   bool
	IEExtensions      []DLCPSecurityInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLCPSecurityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DL-CP-SecurityInformation", v)
}

// DecodeDLCPSecurityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeDLCPSecurityInformation(data []byte) (DLCPSecurityInformation, int, error) {
	var v DLCPSecurityInformation
	err := ngapcodec.DecodeParameter("DL-CP-SecurityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLCPSecurityInformation{}, 0, err
	}
	return v, len(data), err
}
