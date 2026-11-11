// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedNGRANCellInformation represents ASN.1 LastVisitedNGRANCellInformation.
type LastVisitedNGRANCellInformation struct {
	GlobalCellID                             NGRANCGI
	CellType                                 CellType
	TimeUEStayedInCell                       TimeUEStayedInCell
	HasTimeUEStayedInCellEnhancedGranularity bool
	TimeUEStayedInCellEnhancedGranularity    TimeUEStayedInCellEnhancedGranularity
	HasHOCauseValue                          bool
	HOCauseValue                             Cause
	HasIEExtensions                          bool
	IEExtensions                             []LastVisitedNGRANCellInformationExtIEsEntry
	UnknownExtensions                        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedNGRANCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedNGRANCellInformation", v)
}

// DecodeLastVisitedNGRANCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedNGRANCellInformation(data []byte) (LastVisitedNGRANCellInformation, int, error) {
	var v LastVisitedNGRANCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedNGRANCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedNGRANCellInformation{}, 0, err
	}
	return v, len(data), err
}
