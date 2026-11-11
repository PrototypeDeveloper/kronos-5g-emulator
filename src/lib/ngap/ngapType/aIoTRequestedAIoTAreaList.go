// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTRequestedAIoTAreaList represents ASN.1 AIoT-RequestedAIoTAreaList.
type AIoTRequestedAIoTAreaList struct {
	List []AIoTAreaID `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedAIoTAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedAIoTAreaList", v)
}

// DecodeAIoTRequestedAIoTAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedAIoTAreaList(data []byte) (AIoTRequestedAIoTAreaList, int, error) {
	var v AIoTRequestedAIoTAreaList
	err := ngapcodec.DecodeParameter("AIoT-RequestedAIoTAreaList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedAIoTAreaList{}, 0, err
	}
	return v, len(data), err
}
