// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGSTMSI represents ASN.1 FiveG-S-TMSI.
type FiveGSTMSI struct {
	AMFSetID          AMFSetID
	AMFPointer        AMFPointer
	FiveGTMSI         FiveGTMSI
	HasIEExtensions   bool
	IEExtensions      []FiveGSTMSIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGSTMSI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("FiveG-S-TMSI", v) }

// DecodeFiveGSTMSI decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGSTMSI(data []byte) (FiveGSTMSI, int, error) {
	var v FiveGSTMSI
	err := ngapcodec.DecodeParameter("FiveG-S-TMSI", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGSTMSI{}, 0, err
	}
	return v, len(data), err
}
