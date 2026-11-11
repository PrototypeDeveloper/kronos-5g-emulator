// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NPNAccessInformation represents ASN.1 NPN-AccessInformation.
type NPNAccessInformationPresent uint32

const (
	NPNAccessInformationPresentNothing NPNAccessInformationPresent = iota
	NPNAccessInformationPresentPNINPNAccessInformation
	NPNAccessInformationPresentChoiceExtensions
)

type NPNAccessInformation struct {
	Present                 NPNAccessInformationPresent
	PNINPNAccessInformation CellCAGList
	ChoiceExtensions        NPNAccessInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNAccessInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-AccessInformation", v)
}

// DecodeNPNAccessInformation decodes a standalone APER value and reports the consumed octets.
func DecodeNPNAccessInformation(data []byte) (NPNAccessInformation, int, error) {
	var v NPNAccessInformation
	err := ngapcodec.DecodeParameter("NPN-AccessInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNAccessInformation{}, 0, err
	}
	return v, len(data), err
}
