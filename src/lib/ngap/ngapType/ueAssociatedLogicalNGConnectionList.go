// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAssociatedLogicalNGConnectionList represents ASN.1 UE-associatedLogicalNG-connectionList.
type UEAssociatedLogicalNGConnectionList struct {
	List []UEAssociatedLogicalNGConnectionItem `aper:"sizeLB:1,sizeUB:65536"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAssociatedLogicalNGConnectionList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-associatedLogicalNG-connectionList", v)
}

// DecodeUEAssociatedLogicalNGConnectionList decodes a standalone APER value and reports the consumed octets.
func DecodeUEAssociatedLogicalNGConnectionList(data []byte) (UEAssociatedLogicalNGConnectionList, int, error) {
	var v UEAssociatedLogicalNGConnectionList
	err := ngapcodec.DecodeParameter("UE-associatedLogicalNG-connectionList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAssociatedLogicalNGConnectionList{}, 0, err
	}
	return v, len(data), err
}
