// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedPSCellInformation represents ASN.1 LastVisitedPSCellInformation.
type LastVisitedPSCellInformation struct {
	HasPSCellID       bool
	PSCellID          NGRANCGI
	TimeStay          int64 `aper:"valueLB:0,valueUB:40950"`
	HasIEExtensions   bool
	IEExtensions      []LastVisitedPSCellInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedPSCellInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedPSCellInformation", v)
}

// DecodeLastVisitedPSCellInformation decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedPSCellInformation(data []byte) (LastVisitedPSCellInformation, int, error) {
	var v LastVisitedPSCellInformation
	err := ngapcodec.DecodeParameter("LastVisitedPSCellInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedPSCellInformation{}, 0, err
	}
	return v, len(data), err
}
