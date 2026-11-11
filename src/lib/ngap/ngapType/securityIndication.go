// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SecurityIndication represents ASN.1 SecurityIndication.
type SecurityIndication struct {
	IntegrityProtectionIndication          IntegrityProtectionIndication
	ConfidentialityProtectionIndication    ConfidentialityProtectionIndication
	HasMaximumIntegrityProtectedDataRateUL bool
	MaximumIntegrityProtectedDataRateUL    MaximumIntegrityProtectedDataRate
	HasIEExtensions                        bool
	IEExtensions                           []SecurityIndicationExtIEsEntry
	UnknownExtensions                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecurityIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecurityIndication", v)
}

// DecodeSecurityIndication decodes a standalone APER value and reports the consumed octets.
func DecodeSecurityIndication(data []byte) (SecurityIndication, int, error) {
	var v SecurityIndication
	err := ngapcodec.DecodeParameter("SecurityIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecurityIndication{}, 0, err
	}
	return v, len(data), err
}
