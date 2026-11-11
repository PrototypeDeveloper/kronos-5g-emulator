// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNMobilityInformation represents ASN.1 SNPN-MobilityInformation.
type SNPNMobilityInformation struct {
	ServingNID        NID
	HasIEExtensions   bool
	IEExtensions      []SNPNMobilityInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNMobilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-MobilityInformation", v)
}

// DecodeSNPNMobilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNMobilityInformation(data []byte) (SNPNMobilityInformation, int, error) {
	var v SNPNMobilityInformation
	err := ngapcodec.DecodeParameter("SNPN-MobilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNMobilityInformation{}, 0, err
	}
	return v, len(data), err
}
