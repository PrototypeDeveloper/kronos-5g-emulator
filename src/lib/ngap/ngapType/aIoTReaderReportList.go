// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTReaderReportList represents ASN.1 AIoT-ReaderReportList.
type AIoTReaderReportList struct {
	List []AIoTReaderReportItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTReaderReportList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-ReaderReportList", v)
}

// DecodeAIoTReaderReportList decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTReaderReportList(data []byte) (AIoTReaderReportList, int, error) {
	var v AIoTReaderReportList
	err := ngapcodec.DecodeParameter("AIoT-ReaderReportList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTReaderReportList{}, 0, err
	}
	return v, len(data), err
}
