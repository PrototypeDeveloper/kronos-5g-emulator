// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPresenceInAreaOfInterestList represents ASN.1 UEPresenceInAreaOfInterestList.
type UEPresenceInAreaOfInterestList struct {
	List []UEPresenceInAreaOfInterestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPresenceInAreaOfInterestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEPresenceInAreaOfInterestList", v)
}

// DecodeUEPresenceInAreaOfInterestList decodes a standalone APER value and reports the consumed octets.
func DecodeUEPresenceInAreaOfInterestList(data []byte) (UEPresenceInAreaOfInterestList, int, error) {
	var v UEPresenceInAreaOfInterestList
	err := ngapcodec.DecodeParameter("UEPresenceInAreaOfInterestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPresenceInAreaOfInterestList{}, 0, err
	}
	return v, len(data), err
}
