// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGTMSI represents ASN.1 FiveG-TMSI.
type FiveGTMSI struct {
	Value []byte `aper:"sizeLB:4,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGTMSI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("FiveG-TMSI", v) }

// DecodeFiveGTMSI decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGTMSI(data []byte) (FiveGTMSI, int, error) {
	var v FiveGTMSI
	err := ngapcodec.DecodeParameter("FiveG-TMSI", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGTMSI{}, 0, err
	}
	return v, len(data), err
}
