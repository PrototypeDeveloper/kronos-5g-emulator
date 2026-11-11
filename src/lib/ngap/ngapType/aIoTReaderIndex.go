// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTReaderIndex represents ASN.1 AIoT-ReaderIndex.
type AIoTReaderIndex struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTReaderIndex) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-ReaderIndex", v)
}

// DecodeAIoTReaderIndex decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTReaderIndex(data []byte) (AIoTReaderIndex, int, error) {
	var v AIoTReaderIndex
	err := ngapcodec.DecodeParameter("AIoT-ReaderIndex", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTReaderIndex{}, 0, err
	}
	return v, len(data), err
}
