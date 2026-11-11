// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NPNMobilityInformation represents ASN.1 NPN-MobilityInformation.
type NPNMobilityInformationPresent uint32

const (
	NPNMobilityInformationPresentNothing NPNMobilityInformationPresent = iota
	NPNMobilityInformationPresentSNPNMobilityInformation
	NPNMobilityInformationPresentPNINPNMobilityInformation
	NPNMobilityInformationPresentChoiceExtensions
)

type NPNMobilityInformation struct {
	Present                   NPNMobilityInformationPresent
	SNPNMobilityInformation   SNPNMobilityInformation
	PNINPNMobilityInformation PNINPNMobilityInformation
	ChoiceExtensions          NPNMobilityInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNMobilityInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-MobilityInformation", v)
}

// DecodeNPNMobilityInformation decodes a standalone APER value and reports the consumed octets.
func DecodeNPNMobilityInformation(data []byte) (NPNMobilityInformation, int, error) {
	var v NPNMobilityInformation
	err := ngapcodec.DecodeParameter("NPN-MobilityInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNMobilityInformation{}, 0, err
	}
	return v, len(data), err
}
