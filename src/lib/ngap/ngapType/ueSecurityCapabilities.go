// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UESecurityCapabilities represents ASN.1 UESecurityCapabilities.
type UESecurityCapabilities struct {
	NRencryptionAlgorithms             NRencryptionAlgorithms
	NRintegrityProtectionAlgorithms    NRintegrityProtectionAlgorithms
	EUTRAencryptionAlgorithms          EUTRAencryptionAlgorithms
	EUTRAintegrityProtectionAlgorithms EUTRAintegrityProtectionAlgorithms
	HasIEExtensions                    bool
	IEExtensions                       []UESecurityCapabilitiesExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UESecurityCapabilities) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UESecurityCapabilities", v)
}

// DecodeUESecurityCapabilities decodes a standalone APER value and reports the consumed octets.
func DecodeUESecurityCapabilities(data []byte) (UESecurityCapabilities, int, error) {
	var v UESecurityCapabilities
	err := ngapcodec.DecodeParameter("UESecurityCapabilities", data, &v)
	if ngapcodec.IsFatal(err) {
		return UESecurityCapabilities{}, 0, err
	}
	return v, len(data), err
}
