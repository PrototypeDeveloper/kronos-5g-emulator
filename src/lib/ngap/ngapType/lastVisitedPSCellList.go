// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LastVisitedPSCellList represents ASN.1 LastVisitedPSCellList.
type LastVisitedPSCellList struct {
	List []LastVisitedPSCellInformation `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedPSCellList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedPSCellList", v)
}

// DecodeLastVisitedPSCellList decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedPSCellList(data []byte) (LastVisitedPSCellList, int, error) {
	var v LastVisitedPSCellList
	err := ngapcodec.DecodeParameter("LastVisitedPSCellList", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedPSCellList{}, 0, err
	}
	return v, len(data), err
}
