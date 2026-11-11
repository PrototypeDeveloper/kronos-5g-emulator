// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterest represents ASN.1 AreaOfInterest.
type AreaOfInterest struct {
	HasAreaOfInterestTAIList     bool
	AreaOfInterestTAIList        AreaOfInterestTAIList
	HasAreaOfInterestCellList    bool
	AreaOfInterestCellList       AreaOfInterestCellList
	HasAreaOfInterestRANNodeList bool
	AreaOfInterestRANNodeList    AreaOfInterestRANNodeList
	HasIEExtensions              bool
	IEExtensions                 []AreaOfInterestExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterest", v)
}

// DecodeAreaOfInterest decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterest(data []byte) (AreaOfInterest, int, error) {
	var v AreaOfInterest
	err := ngapcodec.DecodeParameter("AreaOfInterest", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterest{}, 0, err
	}
	return v, len(data), err
}
