// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRencryptionAlgorithms represents ASN.1 NRencryptionAlgorithms.
type NRencryptionAlgorithms struct {
	Value BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRencryptionAlgorithms) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRencryptionAlgorithms", v)
}

// DecodeNRencryptionAlgorithms decodes a standalone APER value and reports the consumed octets.
func DecodeNRencryptionAlgorithms(data []byte) (NRencryptionAlgorithms, int, error) {
	var v NRencryptionAlgorithms
	err := ngapcodec.DecodeParameter("NRencryptionAlgorithms", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRencryptionAlgorithms{}, 0, err
	}
	return v, len(data), err
}
