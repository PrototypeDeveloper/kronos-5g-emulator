// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEHistoryInformation represents ASN.1 UEHistoryInformation.
type UEHistoryInformation struct {
	List []LastVisitedCellItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEHistoryInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEHistoryInformation", v)
}

// DecodeUEHistoryInformation decodes a standalone APER value and reports the consumed octets.
func DecodeUEHistoryInformation(data []byte) (UEHistoryInformation, int, error) {
	var v UEHistoryInformation
	err := ngapcodec.DecodeParameter("UEHistoryInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEHistoryInformation{}, 0, err
	}
	return v, len(data), err
}
