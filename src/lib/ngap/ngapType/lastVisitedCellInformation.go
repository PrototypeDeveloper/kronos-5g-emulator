// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedCellInformation represents ASN.1 LastVisitedCellInformation.
type LastVisitedCellInformationPresent uint32

const (
	LastVisitedCellInformationPresentNothing LastVisitedCellInformationPresent = iota
	LastVisitedCellInformationPresentNGRANCell
	LastVisitedCellInformationPresentEUTRANCell
	LastVisitedCellInformationPresentUTRANCell
	LastVisitedCellInformationPresentGERANCell
	LastVisitedCellInformationPresentChoiceExtensions
)

type LastVisitedCellInformation struct {
	Present          LastVisitedCellInformationPresent
	NGRANCell        LastVisitedNGRANCellInformation
	EUTRANCell       LastVisitedEUTRANCellInformation
	UTRANCell        LastVisitedUTRANCellInformation
	GERANCell        LastVisitedGERANCellInformation
	ChoiceExtensions LastVisitedCellInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedCellInformation", v)
}

// DecodeLastVisitedCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedCellInformation(data []byte) (LastVisitedCellInformation, int, error) {
	var v LastVisitedCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedCellInformation{}, 0, err
	}
	return v, len(data), err
}
