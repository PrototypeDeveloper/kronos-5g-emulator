// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GUAMIList represents ASN.1 GUAMIList.
type GUAMIList struct {
	List []GUAMIListItem `aper:"sizeLB:1,sizeUB:1024"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GUAMIList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GUAMIList", v) }

// DecodeGUAMIList decodes a standalone APER value and reports the consumed octets.
func DecodeGUAMIList(data []byte) (GUAMIList, int, error) {
	var v GUAMIList
	err := ngapcodec.DecodeParameter("GUAMIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return GUAMIList{}, 0, err
	}
	return v, len(data), err
}
