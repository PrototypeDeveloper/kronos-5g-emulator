// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTReaderReportItem represents ASN.1 AIoT-ReaderReportItem.
type AIoTReaderReportItem struct {
	ReaderIndex       AIoTReaderIndex
	DeviceReportList  AIoTDeviceReportList
	HasIEExtensions   bool
	IEExtensions      []AIoTReaderReportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTReaderReportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-ReaderReportItem", v)
}

// DecodeAIoTReaderReportItem decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTReaderReportItem(data []byte) (AIoTReaderReportItem, int, error) {
	var v AIoTReaderReportItem
	err := ngapcodec.DecodeParameter("AIoT-ReaderReportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTReaderReportItem{}, 0, err
	}
	return v, len(data), err
}
