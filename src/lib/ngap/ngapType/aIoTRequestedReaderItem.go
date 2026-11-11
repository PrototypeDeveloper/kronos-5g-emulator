// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTRequestedReaderItem represents ASN.1 AIoT-RequestedReaderItem.
type AIoTRequestedReaderItem struct {
	GlobalgNBID       GlobalGNBID
	ReaderIndex       AIoTReaderIndex
	HasIEExtensions   bool
	IEExtensions      []AIoTRequestedReaderItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedReaderItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedReaderItem", v)
}

// DecodeAIoTRequestedReaderItem decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedReaderItem(data []byte) (AIoTRequestedReaderItem, int, error) {
	var v AIoTRequestedReaderItem
	err := ngapcodec.DecodeParameter("AIoT-RequestedReaderItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedReaderItem{}, 0, err
	}
	return v, len(data), err
}
