// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ServedGUAMIList represents ASN.1 ServedGUAMIList.
type ServedGUAMIList struct {
	List []ServedGUAMIItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServedGUAMIList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServedGUAMIList", v)
}

// DecodeServedGUAMIList decodes a standalone APER value and reports the consumed octets.
func DecodeServedGUAMIList(data []byte) (ServedGUAMIList, int, error) {
	var v ServedGUAMIList
	err := ngapcodec.DecodeParameter("ServedGUAMIList", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServedGUAMIList{}, 0, err
	}
	return v, len(data), err
}
