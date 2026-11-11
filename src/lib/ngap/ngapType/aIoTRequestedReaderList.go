// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTRequestedReaderList represents ASN.1 AIoT-RequestedReaderList.
type AIoTRequestedReaderList struct {
	List []AIoTRequestedReaderItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedReaderList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedReaderList", v)
}

// DecodeAIoTRequestedReaderList decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedReaderList(data []byte) (AIoTRequestedReaderList, int, error) {
	var v AIoTRequestedReaderList
	err := ngapcodec.DecodeParameter("AIoT-RequestedReaderList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedReaderList{}, 0, err
	}
	return v, len(data), err
}
