// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRAencryptionAlgorithms represents ASN.1 EUTRAencryptionAlgorithms.
type EUTRAencryptionAlgorithms struct {
	Value BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAencryptionAlgorithms) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAencryptionAlgorithms", v)
}

// DecodeEUTRAencryptionAlgorithms decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAencryptionAlgorithms(data []byte) (EUTRAencryptionAlgorithms, int, error) {
	var v EUTRAencryptionAlgorithms
	err := ngapcodec.DecodeParameter("EUTRAencryptionAlgorithms", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAencryptionAlgorithms{}, 0, err
	}
	return v, len(data), err
}
