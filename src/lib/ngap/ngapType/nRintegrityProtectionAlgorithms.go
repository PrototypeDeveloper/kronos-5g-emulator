// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRintegrityProtectionAlgorithms represents ASN.1 NRintegrityProtectionAlgorithms.
type NRintegrityProtectionAlgorithms struct {
	Value BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRintegrityProtectionAlgorithms) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRintegrityProtectionAlgorithms", v)
}

// DecodeNRintegrityProtectionAlgorithms decodes a standalone APER value and reports the consumed octets.
func DecodeNRintegrityProtectionAlgorithms(data []byte) (NRintegrityProtectionAlgorithms, int, error) {
	var v NRintegrityProtectionAlgorithms
	err := ngapcodec.DecodeParameter("NRintegrityProtectionAlgorithms", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRintegrityProtectionAlgorithms{}, 0, err
	}
	return v, len(data), err
}
