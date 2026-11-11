// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRAintegrityProtectionAlgorithms represents ASN.1 EUTRAintegrityProtectionAlgorithms.
type EUTRAintegrityProtectionAlgorithms struct {
	Value BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAintegrityProtectionAlgorithms) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAintegrityProtectionAlgorithms", v)
}

// DecodeEUTRAintegrityProtectionAlgorithms decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAintegrityProtectionAlgorithms(data []byte) (EUTRAintegrityProtectionAlgorithms, int, error) {
	var v EUTRAintegrityProtectionAlgorithms
	err := ngapcodec.DecodeParameter("EUTRAintegrityProtectionAlgorithms", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAintegrityProtectionAlgorithms{}, 0, err
	}
	return v, len(data), err
}
