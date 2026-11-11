// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSServiceAreaInformation represents ASN.1 MBS-ServiceAreaInformation.
type MBSServiceAreaInformation struct {
	HasMBSServiceAreaCellList bool
	MBSServiceAreaCellList    MBSServiceAreaCellList
	HasMBSServiceAreaTAIList  bool
	MBSServiceAreaTAIList     MBSServiceAreaTAIList
	HasIEExtensions           bool
	IEExtensions              []MBSServiceAreaInformationExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaInformation", v)
}

// DecodeMBSServiceAreaInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaInformation(data []byte) (MBSServiceAreaInformation, int, error) {
	var v MBSServiceAreaInformation
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaInformation{}, 0, err
	}
	return v, len(data), err
}
