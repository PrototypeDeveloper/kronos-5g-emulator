// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M1ThresholdEventA2 represents ASN.1 M1ThresholdEventA2.
type M1ThresholdEventA2 struct {
	M1ThresholdType   M1ThresholdType
	HasIEExtensions   bool
	IEExtensions      []M1ThresholdEventA2ExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ThresholdEventA2) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1ThresholdEventA2", v)
}

// DecodeM1ThresholdEventA2 decodes a standalone APER value and reports the consumed octets.
func DecodeM1ThresholdEventA2(data []byte) (M1ThresholdEventA2, int, error) {
	var v M1ThresholdEventA2
	err := ngapcodec.DecodeParameter("M1ThresholdEventA2", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ThresholdEventA2{}, 0, err
	}
	return v, len(data), err
}
