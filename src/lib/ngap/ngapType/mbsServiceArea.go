// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceArea represents ASN.1 MBS-ServiceArea.
type MBSServiceAreaPresent uint32

const (
	MBSServiceAreaPresentNothing MBSServiceAreaPresent = iota
	MBSServiceAreaPresentLocationindependent
	MBSServiceAreaPresentLocationdependent
	MBSServiceAreaPresentChoiceExtensions
)

type MBSServiceArea struct {
	Present             MBSServiceAreaPresent
	Locationindependent MBSServiceAreaInformation
	Locationdependent   MBSServiceAreaInformationList
	ChoiceExtensions    MBSServiceAreaExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceArea) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceArea", v)
}

// DecodeMBSServiceArea decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceArea(data []byte) (MBSServiceArea, int, error) {
	var v MBSServiceArea
	err := ngapcodec.DecodeParameter("MBS-ServiceArea", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceArea{}, 0, err
	}
	return v, len(data), err
}
