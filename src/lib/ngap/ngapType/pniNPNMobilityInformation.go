// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PNINPNMobilityInformation represents ASN.1 PNI-NPN-MobilityInformation.
type PNINPNMobilityInformation struct {
	AllowedPNINPIList AllowedPNINPNList
	HasIEExtensions   bool
	IEExtensions      []PNINPNMobilityInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNMobilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPN-MobilityInformation", v)
}

// DecodePNINPNMobilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNMobilityInformation(data []byte) (PNINPNMobilityInformation, int, error) {
	var v PNINPNMobilityInformation
	err := ngapcodec.DecodeParameter("PNI-NPN-MobilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNMobilityInformation{}, 0, err
	}
	return v, len(data), err
}
