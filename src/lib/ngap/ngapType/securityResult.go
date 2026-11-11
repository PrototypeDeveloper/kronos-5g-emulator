// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecurityResult represents ASN.1 SecurityResult.
type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult
	ConfidentialityProtectionResult ConfidentialityProtectionResult
	HasIEExtensions                 bool
	IEExtensions                    []SecurityResultExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityResult) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityResult", v)
}

// DecodeSecurityResult decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityResult(data []byte) (SecurityResult, int, error) {
	var v SecurityResult
	err := ngapcodec.DecodeParameter("SecurityResult", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityResult{}, 0, err
	}
	return v, len(data), err
}
