// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AveragingWindow represents ASN.1 AveragingWindow.
type AveragingWindow struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4095"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AveragingWindow) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AveragingWindow", v)
}

// DecodeAveragingWindow decodes a standalone APER value and reports the consumed octets.
func DecodeAveragingWindow(data []byte) (AveragingWindow, int, error) {
	var v AveragingWindow
	err := ngapcodec.DecodeParameter("AveragingWindow", data, &v)
	if ngapcodec.IsFatal(err) {
		return AveragingWindow{}, 0, err
	}
	return v, len(data), err
}
