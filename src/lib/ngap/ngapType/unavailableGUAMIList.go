// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UnavailableGUAMIList represents ASN.1 UnavailableGUAMIList.
type UnavailableGUAMIList struct {
	List []UnavailableGUAMIItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UnavailableGUAMIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UnavailableGUAMIList", v)
}

// DecodeUnavailableGUAMIList decodes a standalone APER value and reports the consumed octets.
func DecodeUnavailableGUAMIList(data []byte) (UnavailableGUAMIList, int, error) {
	var v UnavailableGUAMIList
	err := ngapcodec.DecodeParameter("UnavailableGUAMIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UnavailableGUAMIList{}, 0, err
	}
	return v, len(data), err
}
